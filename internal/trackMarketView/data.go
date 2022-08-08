package trackMarketView

import (
	"github.com/bhbosman/goTrader/publish"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
	"sort"
)

type data struct {
	MessageRouter       *messageRouter.MessageRouter
	strategyMap         map[string]publish.IStrategy
	strategyListIsDirty bool
	onListChange        func(data []string) bool
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

func (self *data) handlePublishData(msg *publish.PublishData) {
	if _, ok := self.strategyMap[msg.GetStrategyName()]; !ok {
		self.strategyMap[msg.GetStrategyName()] = msg
		self.strategyListIsDirty = true
	}
}

func (self *data) handleEmptyQueue(msg *messages.EmptyQueue) {
	if self.strategyListIsDirty {
		ss := make([]string, 0, len(self.strategyMap))
		for key, _ := range self.strategyMap {
			ss = append(ss, key)
		}
		sort.Strings(ss)
		if self.onListChange(ss) {
			self.strategyListIsDirty = false
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
	//
	return result, nil
}
