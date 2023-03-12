package strategyStateManagerView

import (
	"github.com/bhbosman/goTrader/internal/publish"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
)

type data struct {
	MessageRouter        messageRouter.IMessageRouter
	strategyListIsDirty  bool
	onListChange         func(data []string) bool
	onStrategyDataChange func(name string, data publish.IStrategy) bool
	activeItem           string
	activeData           publish.IStrategy
	strategies           []string
}

func (self *data) Unsubscribe(item string) {
	self.activeItem = ""
}

func (self *data) Subscribe(item string) {
	self.activeItem = item
}

func (self *data) SetStrategyDataChange(onStrategyDataChange func(name string, data publish.IStrategy) bool) {
	self.onStrategyDataChange = onStrategyDataChange
}

func (self *data) SetListChange(onListChange func(data []string) bool) {
	self.onListChange = onListChange
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

func (self *data) handleStrategyList(msg *publish.StrategyList) {
	self.strategies = msg.Strategies
	self.strategyListIsDirty = true
}

func (self *data) handlePublishData(msg *publish.PublishData) {
	if self.activeItem == msg.GetStrategyName() {
		self.activeData = msg
	}
}

func (self *data) handleEmptyQueue(msg *messages.EmptyQueue) {
	if self.strategyListIsDirty {
		if self.onListChange != nil {
			if self.onListChange(self.strategies) {
				self.strategyListIsDirty = false
			} else {
				msg.ErrorHappen = true
				return
			}
		}
	}
	if self.activeData != nil {
		if self.onStrategyDataChange != nil {
			b := self.onStrategyDataChange(self.activeData.GetStrategyName(), self.activeData)
			if !b {
				msg.ErrorHappen = true
				return
			} else {
				self.activeData = nil
			}
		}
	}
}

func newData() (ITrackMarketViewData, error) {
	result := &data{
		MessageRouter: messageRouter.NewMessageRouter(),
	}
	_ = result.MessageRouter.Add(result.handleEmptyQueue)
	_ = result.MessageRouter.Add(result.handlePublishData)
	_ = result.MessageRouter.Add(result.handleStrategyList)
	//
	return result, nil
}
