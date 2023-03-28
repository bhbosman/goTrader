package trackMarket

import (
	"github.com/bhbosman/goMessages/marketData/stream"
	"github.com/bhbosman/gocommon/services/IDataShutDown"
	"github.com/bhbosman/gocommon/services/IFxService"
	"github.com/bhbosman/gocommon/services/ISendMessage"
	"github.com/cskr/pubsub"
)

type ITrackMarket interface {
	ISendMessage.ISendMessage
	ISendMessage.IMultiSendMessage
}

type ITrackMarketService interface {
	ITrackMarket
	IFxService.IFxServices
}

type ITrackMarketData interface {
	ITrackMarket
	IDataShutDown.IDataShutDown
	SetSubscriptionReceiver(channel *pubsub.NextFuncSubscription)
}

type PriceVolumeResponse struct {
	TradingInterface string
	OrderId          int
	Side             string
	Price            float64
	Volume           float64
}

type IPricingVolumeCalculation interface {
	Instruments() []string
	Calculate(activeDataMap map[string]*stream.PublishTop5) ([]*PriceVolumeResponse, error)
	StrategyName() string
}
