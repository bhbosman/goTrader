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

type IUi interface {
	SetListChange(onListChange func(data []string) bool)
}

type ITrackMarketViewService interface {
	ITrackMarketView
	IFxService.IFxServices
	IUi
}

type ITrackMarketViewData interface {
	ITrackMarketView
	IDataShutDown.IDataShutDown
	IUi
}
