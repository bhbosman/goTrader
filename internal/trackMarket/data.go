package trackMarket

import (
	"github.com/bhbosman/goMessages/marketData/stream"
	"github.com/bhbosman/goTrader/internal/trackMarketView"
	"github.com/bhbosman/goTrader/publish"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
)

type data struct {
	MessageRouter          *messageRouter.MessageRouter
	activeDataMap          map[string]*stream.PublishTop5
	modelSettings          modelSettings
	TrackMarketViewService trackMarketView.ITrackMarketViewService
}

func (self *data) MultiSend(messages ...interface{}) {
	self.MessageRouter.MultiRoute(messages...)
}

func (self *data) Send(message interface{}) error {
	self.MessageRouter.Route(message)
	return nil
}

func (self *data) ShutDown() error {
	return nil
}

func (self *data) handlePublishTop5(msg *stream.PublishTop5) {
	self.activeDataMap[msg.Instrument] = msg
}

func (self *data) handleEmptyQueue(msg *messages.EmptyQueue) {
	if top5, ok := self.activeDataMap[self.modelSettings.instrument]; ok {
		maxIndex := func(a int, b []*stream.Point) int {
			if len(b) < a {
				return len(b)
			}
			return a
		}
		publishData := &publish.PublishData{
			StrategyName: self.modelSettings.Name,
		}
		for i := 0; i < maxIndex(5, top5.Ask); i++ {
			publishData.Lines[i].Ask.Price = top5.Ask[i].Price
			publishData.Lines[i].Ask.Volume = top5.Ask[i].Volume
		}
		for i := 0; i < maxIndex(5, top5.Bid); i++ {
			publishData.Lines[i].Bid.Price = top5.Bid[i].Price
			publishData.Lines[i].Bid.Volume = top5.Bid[i].Volume
		}
		_ = self.TrackMarketViewService.Send(publishData)
	}
}

func newData(
	trackMarketViewService trackMarketView.ITrackMarketViewService,
	modelSettings modelSettings,
) (ITrackMarketData, error) {
	result := &data{
		MessageRouter:          messageRouter.NewMessageRouter(),
		activeDataMap:          make(map[string]*stream.PublishTop5),
		modelSettings:          modelSettings,
		TrackMarketViewService: trackMarketViewService,
	}
	result.MessageRouter.Add(result.handleEmptyQueue)
	result.MessageRouter.Add(result.handlePublishTop5)
	//
	return result, nil
}
