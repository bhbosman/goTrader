package strategyStateManagerService

import (
	"fmt"
	"github.com/bhbosman/goTrader/internal/publish"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
	"github.com/cskr/pubsub"
	"sort"
)

const StrategyListUpdate = "StrategyListUpdate"

func StrategyUpdate(name string) string {
	return fmt.Sprintf("STRATEGY_%v", name)
}

type data struct {
	MessageRouter       *messageRouter.MessageRouter
	strategyMap         map[string]publish.IStrategy
	dirtyStrategies     map[string]bool
	strategyListIsDirty bool
	PubSub              *pubsub.PubSub
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

func (self *data) handleEmptyQueue(msg *messages.EmptyQueue) {
	if self.strategyListIsDirty {
		l := self.createStrategyList()
		if self.PubSub.PubWithContext(l, StrategyListUpdate) {
			self.strategyListIsDirty = false
		} else {
			msg.ErrorHappen = true
			return
		}
	}

	if len(self.dirtyStrategies) > 0 {
		errorOccurred := false
		handled := make([]string, 0, len(self.dirtyStrategies))
		for s := range self.dirtyStrategies {
			v := self.strategyMap[s]
			if self.PubSub.PubWithContext(v, StrategyUpdate(s)) {
				handled = append(handled, s)
			} else {
				errorOccurred = true
				msg.ErrorHappen = true
				break
			}
		}
		if errorOccurred {
			for _, s := range handled {
				delete(self.dirtyStrategies, s)
			}
		} else {
			self.dirtyStrategies = make(map[string]bool)
		}
	}
}

func (self *data) handleRequestStrategyItem(msg *publish.RequestStrategyItem) {
	if msg.Callback != nil {
		if v, ok := self.strategyMap[msg.Item]; ok {
			msg.Callback.Add(v)
		}
	}
}

func (self *data) handleRequestStrategyList(msg *publish.RequestStrategyList) {
	if msg.Callback != nil {
		msg.Callback.Add(self.createStrategyList())
	}
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
	self.dirtyStrategies[msg.GetStrategyName()] = true
}

func (self *data) createStrategyList() *publish.StrategyList {
	ss := make([]string, 0, len(self.strategyMap))
	for s := range self.strategyMap {
		ss = append(ss, s)
	}
	sort.Strings(ss)
	return &publish.StrategyList{
		Strategies: ss,
	}
}

func newData(PubSub *pubsub.PubSub) (IStrategyManagerData, error) {
	result := &data{
		MessageRouter:   messageRouter.NewMessageRouter(),
		strategyMap:     make(map[string]publish.IStrategy),
		dirtyStrategies: make(map[string]bool),
		PubSub:          PubSub,
	}
	result.MessageRouter.Add(result.handleEmptyQueue)
	result.MessageRouter.Add(result.handlePublishData)
	result.MessageRouter.Add(result.handleDeleteStrategy)
	result.MessageRouter.Add(result.handleRequestStrategyList)
	result.MessageRouter.Add(result.handleRequestStrategyItem)
	//
	return result, nil
}
