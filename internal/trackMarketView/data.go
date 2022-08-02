package trackMarketView

import (
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
)

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

func (self *data) handleEmptyQueue(msg *messages.EmptyQueue) {
}

func newData() (ITrackMarketViewData, error) {
	result := &data{
		MessageRouter: messageRouter.NewMessageRouter(),
	}
	result.MessageRouter.Add(result.handleEmptyQueue)
	//
	return result, nil
}
