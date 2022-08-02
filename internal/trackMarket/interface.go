package trackMarket

import (
	"github.com/bhbosman/gocommon/Services/IDataShutDown"
	"github.com/bhbosman/gocommon/Services/IFxService"
	"github.com/bhbosman/gocommon/services/ISendMessage"
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
}
