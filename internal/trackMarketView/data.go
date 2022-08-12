package trackMarketView

import (
	"github.com/bhbosman/goTrader/internal/publish"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
	"sort"
)

type data struct {
	MessageRouter        *messageRouter.MessageRouter
	strategyMap          map[string]publish.IStrategy
	strategyListIsDirty  bool
	onListChange         func(data []string) bool
	onStrategyDataChange func(name string, data publish.IStrategy) bool
	activeItem           string
	activeData           publish.IStrategy
}

func (self *data) Unsubscribe(item string) {
	self.activeItem = ""
}

func (self *data) Subscribe(item string) {
	self.activeItem = item

	self.activeData, _ = self.strategyMap[self.activeItem]
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

func (self *data) handleDeleteStrategy(msg *publish.DeleteStrategy) {
	if _, ok := self.strategyMap[msg.StrategyName]; ok {
		delete(self.strategyMap, msg.StrategyName)
		self.strategyListIsDirty = true
	}
}

func (self *data) handlePublishData(msg *publish.PublishData) {
	if _, ok := self.strategyMap[msg.GetStrategyName()]; !ok {
		self.strategyListIsDirty = true
	}
	self.strategyMap[msg.GetStrategyName()] = msg
	if self.activeItem == msg.GetStrategyName() {
		self.activeData = msg
	}
}

func (self *data) handleEmptyQueue(msg *messages.EmptyQueue) {
	if self.strategyListIsDirty {
		if self.onListChange != nil {
			ss := make([]string, 0, len(self.strategyMap))
			for key, _ := range self.strategyMap {
				ss = append(ss, key)
			}
			sort.Strings(ss)
			if self.onListChange(ss) {
				msg.ErrorHappen = true
				self.strategyListIsDirty = false
			} else {
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
		strategyMap:   make(map[string]publish.IStrategy),
	}
	result.MessageRouter.Add(result.handleEmptyQueue)
	result.MessageRouter.Add(result.handlePublishData)
	result.MessageRouter.Add(result.handleDeleteStrategy)
	//
	return result, nil
}
