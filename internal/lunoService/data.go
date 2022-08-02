package lunoService

import (
	"github.com/bhbosman/goTrader/internal/lunoApi/client"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
	"net/http"
)

type data struct {
	MessageRouter *messageRouter.MessageRouter
	roundTripper  http.RoundTripper
	RequestDoer   client.HttpRequestDoer
}

func (self *data) Start() {
	//lunoClient, err := client.NewClientWithResponses(
	//	"https://api.luno.com",
	//	func(c *client.Client) error {
	//		c.Client = self.RequestDoer
	//		return nil
	//	},
	//)
	//if err != nil {
	//	return
	//}
	//balancesWithResponse, err := lunoClient.GetBalancesWithResponse(
	//	context.Background(),
	//	&components.GetBalancesParams{
	//		Assets: nil,
	//	})
	//if err != nil {
	//	return
	//}
	//_ = balancesWithResponse.JSON200.Balance
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

func newData(
	roundTripper http.RoundTripper,
	RequestDoer client.HttpRequestDoer) (ILunoServiceData, error) {
	result := &data{
		roundTripper:  roundTripper,
		MessageRouter: messageRouter.NewMessageRouter(),
		RequestDoer:   RequestDoer,
	}
	result.MessageRouter.Add(result.handleEmptyQueue)
	//
	return result, nil
}
