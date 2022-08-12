package trackMarket

import (
	"context"
	stream2 "github.com/bhbosman/goCommonMarketData/fullMarketData/stream"
	"github.com/bhbosman/goCommonMarketData/fullMarketDataHelper"
	"github.com/bhbosman/goCommonMarketData/fullMarketDataManagerService"
	"github.com/bhbosman/goMessages/marketData/stream"
	"github.com/bhbosman/goTrader/internal/lunoService"
	"github.com/bhbosman/goTrader/internal/publish"
	"github.com/bhbosman/goTrader/internal/trackMarketView"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
	"github.com/cskr/pubsub"
)

type StateFlags uint64

func (self *StateFlags) Set(flag StateFlags) StateFlags    { return *self | flag }
func (self *StateFlags) Clear(flag StateFlags) StateFlags  { return *self &^ flag }
func (self *StateFlags) Toggle(flag StateFlags) StateFlags { return *self ^ flag }
func (self *StateFlags) Has(flag StateFlags) bool          { return *self&flag == flag }

const (
	marketDataStatusReceived StateFlags = 1 << iota
	subscriptionChannelReceived
	marketDataRegisteredState
	listOrderOutstanding
)

type data struct {
	MessageRouter          *messageRouter.MessageRouter
	activeDataMap          map[string]*stream.PublishTop5
	modelSettings          IPricingVolumeCalculation
	TrackMarketViewService trackMarketView.ITrackMarketViewService
	stateFunc              func() (bool, string, error)
	state                  string
	flags                  StateFlags
	instrumentStatus       []*stream2.InstrumentStatus
	FmdService             fullMarketDataManagerService.IFmdManagerService
	FullMarketDataHelper   fullMarketDataHelper.IFullMarketDataHelper
	subscriptionReceiver   *pubsub.NextFuncSubscription
	pubSub                 *pubsub.PubSub
	LunoServiceService     lunoService.ILunoServiceService
	cancelCtx              context.Context
}

func (self *data) SetSubscriptionReceiver(channel *pubsub.NextFuncSubscription) {
	self.subscriptionReceiver = channel
	self.flags = self.flags.Set(subscriptionChannelReceived)

	key := self.FullMarketDataHelper.InstrumentListChannelName()
	self.pubSub.AddSub(self.subscriptionReceiver, key)

	msg := &stream2.FullMarketData_InstrumentList_RequestWrapper{
		Data: &stream2.FullMarketData_InstrumentList_Request{},
	}
	msg.SetNext(self.subscriptionReceiver)
	_ = self.FmdService.Send(msg)
}

func (self *data) MultiSend(messages ...interface{}) {
	self.MessageRouter.MultiRoute(messages...)
}

func (self *data) Send(message interface{}) error {
	self.MessageRouter.Route(message)
	return nil
}

func (self *data) ShutDown() error {
	self.unregisterMarketData()
	_ = self.TrackMarketViewService.Send(
		&publish.DeleteStrategy{
			StrategyName: self.modelSettings.StrategyName(),
		},
	)
	return nil
}

//goland:noinspection GoSnakeCaseUsage
func (self *data) handleFullMarketData_InstrumentList_Response(msg *stream2.FullMarketData_InstrumentList_Response) {
	self.instrumentStatus = msg.Instruments
	self.flags = self.flags.Set(marketDataStatusReceived)
}

func (self *data) handleCallbackMessage(msg *callbackMessage) {
	if msg.cb != nil {
		msg.cb()
	}
}

func (self *data) handlePublishTop5(msg *stream.PublishTop5) {
	self.activeDataMap[msg.Instrument] = msg
}

func (self *data) handleEmptyQueue(msg *messages.EmptyQueue) {
	for true {
		var state bool
		var err error
		state, self.state, err = self.stateFunc()
		if err != nil {
			break
		}
		if !state {
			break
		}
	}
	publishData := &publish.PublishData{
		StrategyName: self.modelSettings.StrategyName(),
		State:        self.state,
	}

	if top5, ok := self.activeDataMap[self.modelSettings.Instruments()[0]]; ok {
		maxIndex := func(a int, b []*stream.Point) int {
			if len(b) < a {
				return len(b)
			}
			return a
		}
		for i := 0; i < maxIndex(5, top5.Ask); i++ {
			publishData.Lines[i].Ask.Price = top5.Ask[i].Price
			publishData.Lines[i].Ask.Volume = top5.Ask[i].Volume
		}
		for i := 0; i < maxIndex(5, top5.Bid); i++ {
			publishData.Lines[i].Bid.Price = top5.Bid[i].Price
			publishData.Lines[i].Bid.Volume = top5.Bid[i].Volume
		}
	}
	_ = self.TrackMarketViewService.Send(publishData)
}
func (self *data) initState() (bool, string, error) {
	const state = "Init"
	switch {
	case self.flags.Has(subscriptionChannelReceived):
		self.stateFunc = self.queryForOpenOrdersState
		self.flags = self.flags.Clear(subscriptionChannelReceived)
		return true, state, nil
	default:
		return false, state, nil
	}
}

func (self *data) queryForOpenOrdersState() (bool, string, error) {
	const state = "Query OpenOrders"
	self.LunoServiceService.ListOrders(
		self.cancelCtx,
		lunoService.ListOrderRequest{},
		func(response *lunoService.ListOrderResponse) {
			self.subscriptionReceiver.Add(
				&callbackMessage{
					cb: func() {
						self.flags = self.flags.Clear(listOrderOutstanding)
					},
				},
			)
		},
	)
	self.flags = self.flags.Set(listOrderOutstanding)
	self.stateFunc = self.waitForListOrderResponse
	return true, state, nil
}

func (self *data) waitForListOrderResponse() (bool, string, error) {
	const state = "WaitForListOrderResponse"
	switch {
	case self.flags.Has(listOrderOutstanding):
		return false, state, nil
	default:
		self.stateFunc = self.waitForMarketData
		return true, state, nil
	}
}

func (self *data) waitForMarketData() (bool, string, error) {
	const state = "waitForMarketData"
	switch {
	case self.flags.Has(subscriptionChannelReceived):
		self.stateFunc = self.initState
		return true, state, nil
	case self.flags.Has(marketDataStatusReceived):
		defer func() {
			self.flags = self.flags.Clear(marketDataStatusReceived)
		}()
		m := make(map[string]*stream2.InstrumentStatus)
		for _, status := range self.instrumentStatus {
			m[status.Instrument] = status
		}
		status, ok := m[self.modelSettings.Instruments()[0]]
		if ok && status.Status != "" {
			self.registerMarketData()
			self.stateFunc = self.calculateAndWaitingForInstruction
			return true, state, nil
		} else {
			self.unregisterMarketData()
		}
		return false, state, nil
	default:
		return false, state, nil
	}
}

func (self *data) calculateAndWaitingForInstruction() (bool, string, error) {
	const state = "calculateAndWaitingForInstruction"
	switch {
	case self.flags.Has(subscriptionChannelReceived):
		self.stateFunc = self.initState
		return true, state, nil
	case self.flags.Has(marketDataStatusReceived):
		self.stateFunc = self.waitForMarketData
		return true, state, nil
	case self.flags.Has(marketDataRegisteredState):
		return false, state, nil
	default:
		return false, state, nil
	}
}

func (self *data) unregisterMarketData() {
	if self.flags.Has(marketDataRegisteredState) {
		key := self.FullMarketDataHelper.InstrumentChannelNameForTop5(self.modelSettings.Instruments()[0])
		self.pubSub.Unsub(self.subscriptionReceiver, key)
		self.FmdService.UnsubscribeFullMarketData(self.modelSettings.Instruments()[0])
		self.flags = self.flags.Clear(marketDataRegisteredState)
	}
}

func (self *data) registerMarketData() {
	if !self.flags.Has(marketDataRegisteredState) {
		key := self.FullMarketDataHelper.InstrumentChannelNameForTop5(self.modelSettings.Instruments()[0])
		self.pubSub.AddSub(self.subscriptionReceiver, key)
		self.FmdService.SubscribeFullMarketData(self.modelSettings.Instruments()[0])
		self.flags = self.flags.Set(marketDataRegisteredState)
	}
}

func newData(
	cancelCtx context.Context,
	trackMarketViewService trackMarketView.ITrackMarketViewService,
	modelSettings IPricingVolumeCalculation,
	FmdService fullMarketDataManagerService.IFmdManagerService,
	FullMarketDataHelper fullMarketDataHelper.IFullMarketDataHelper,
	pubSub *pubsub.PubSub,
	LunoServiceService lunoService.ILunoServiceService,
) (ITrackMarketData, error) {
	result := &data{
		cancelCtx:              cancelCtx,
		MessageRouter:          messageRouter.NewMessageRouter(),
		activeDataMap:          make(map[string]*stream.PublishTop5),
		modelSettings:          modelSettings,
		TrackMarketViewService: trackMarketViewService,
		flags:                  0,
		FmdService:             FmdService,
		FullMarketDataHelper:   FullMarketDataHelper,
		pubSub:                 pubSub,
		LunoServiceService:     LunoServiceService,
	}
	result.MessageRouter.Add(result.handleEmptyQueue)
	result.MessageRouter.Add(result.handlePublishTop5)
	result.MessageRouter.Add(result.handleCallbackMessage)
	result.MessageRouter.Add(result.handleFullMarketData_InstrumentList_Response)
	result.stateFunc = result.initState
	//
	return result, nil
}
