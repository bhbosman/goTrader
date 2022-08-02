package trackMarketView

import (
	"github.com/bhbosman/gocommon/Services/IDataShutDown"
	"github.com/bhbosman/gocommon/Services/IFxService"
	"github.com/bhbosman/gocommon/services/ISendMessage"
)

type ITrackMarketView interface {
	ISendMessage.ISendMessage
	ISendMessage.IMultiSendMessage
}

type ITrackMarketViewService interface {
	ITrackMarketView
	IFxService.IFxServices
}

type ITrackMarketViewData interface {
	ITrackMarketView
	IDataShutDown.IDataShutDown
}
