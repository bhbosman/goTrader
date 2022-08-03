package trackMarket

import (
	"github.com/bhbosman/goMessages/marketData/stream"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
)

type data struct {
	MessageRouter *messageRouter.MessageRouter
	activeDataMap map[string]*stream.PublishTop5
	modelSettings modelSettings
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
}

func newData(modelSettings modelSettings) (ITrackMarketData, error) {
	result := &data{
		MessageRouter: messageRouter.NewMessageRouter(),
		modelSettings: modelSettings,
		activeDataMap: make(map[string]*stream.PublishTop5),
	}
	result.MessageRouter.Add(result.handleEmptyQueue)
	result.MessageRouter.Add(result.handlePublishTop5)
	//
	return result, nil
}