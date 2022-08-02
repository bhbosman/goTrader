package trackMarket

import (
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
)

type OnITrackMarketDataCreate func(modelSettings modelSettings) (ITrackMarketData, error)

type data struct {
	MessageRouter *messageRouter.MessageRouter
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

func (self *data) handlePublishTop5(msg *messages.EmptyQueue) {
}

func (self *data) handleEmptyQueue(msg *messages.EmptyQueue) {
}

func newData(modelSettings modelSettings) (ITrackMarketData, error) {
	result := &data{
		MessageRouter: messageRouter.NewMessageRouter(),
	}
	result.MessageRouter.Add(result.handleEmptyQueue)
	result.MessageRouter.Add(result.handlePublishTop5)

	//
	return result, nil
}
