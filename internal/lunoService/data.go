package lunoService

import (
	"context"
	"github.com/bhbosman/goTrader/internal/lunoApi/client"
	"github.com/bhbosman/goTrader/internal/lunoApi/components"
	"github.com/bhbosman/gocommon/messageRouter"
	"github.com/bhbosman/gocommon/messages"
	"strconv"
)

type data struct {
	MessageRouter       *messageRouter.MessageRouter
	ClientWithResponses client.ClientWithResponsesInterface
}

func (self *data) ListOrders(ctx context.Context, params ListOrderRequest, cb ListOrderResponseCallback) {
	response, err := self.ClientWithResponses.ListOrdersV2WithResponse(
		ctx,
		&components.ListOrdersV2Params{
			Pair:          &params.Instrument,
			Closed:        nil,
			CreatedBefore: nil,
			Limit:         nil,
		},
	)
	orders := make([]*OrderInformation, len(*response.JSON200.Orders))

	for i, order := range *response.JSON200.Orders {
		price, _ := strconv.ParseFloat(*order.LimitPrice, 64)
		volume, _ := strconv.ParseFloat(*order.LimitVolume, 64)
		orders[i] = &OrderInformation{
			ClientReference: *order.ClientOrderID,
			OrderId:         *order.Ref,
			Instrument:      *order.Pair,
			Price:           price,
			Volume:          volume,
		}

	}
	if cb != nil {
		cb(
			&ListOrderResponse{
				MessageId: params.MessageId,
				Status:    response.StatusCode(),
				ErrorMessage: func() string {
					switch {
					case response.JSONDefault == nil:
						return ""
					default:
						return *response.JSONDefault.Message
					}
				}(),
				Error:  err,
				Orders: orders,
			},
		)
	}
}

func (self *data) CancelOrder(ctx context.Context, params CancelOrderRequest, cb CancelOrderRequestResponseCallback) {
}

func (self *data) Start() {
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
	ClientWithResponses client.ClientWithResponsesInterface,
) (ILunoServiceData, error) {
	result := &data{
		MessageRouter:       messageRouter.NewMessageRouter(),
		ClientWithResponses: ClientWithResponses,
	}
	result.MessageRouter.Add(result.handleEmptyQueue)
	//
	return result, nil
}
