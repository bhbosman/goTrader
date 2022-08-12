package lunoService

import "context"

type CancelOrderRequest struct {
	MessageId string
}
type CancelOrderRequestResponse struct {
	MessageId string
}
type CancelOrderRequestResponseCallback func(cancelResponse *CancelOrderRequestResponse)

type ListOrderRequest struct {
	MessageId  string
	Instrument string
}

type OrderInformation struct {
	ClientReference string  `json:"client_reference,omitempty"`
	OrderId         string  `json:"order_id,omitempty"`
	Instrument      string  `json:"instrument,omitempty"`
	Price           float64 `json:"price,omitempty"`
	Volume          float64 `json:"volume,omitempty"`
}

type ListOrderResponse struct {
	MessageId    string
	Error        error
	Status       int
	ErrorMessage string
	Orders       []*OrderInformation
}

type ListOrderResponseCallback func(listResponse *ListOrderResponse)

type IExchange interface {
	CancelOrder(ctx context.Context, params CancelOrderRequest, cb CancelOrderRequestResponseCallback)
	ListOrders(ctx context.Context, params ListOrderRequest, cb ListOrderResponseCallback)
}
