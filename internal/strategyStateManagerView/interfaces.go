package strategyStateManagerView

import (
	"github.com/bhbosman/goTrader/internal/publish"
	"github.com/bhbosman/gocommon/services/IDataShutDown"
	"github.com/bhbosman/gocommon/services/IFxService"
	"github.com/bhbosman/gocommon/services/ISendMessage"
	"github.com/rivo/tview"
	"io"
)

type ITrackMarketView interface {
	ISendMessage.ISendMessage
	ISendMessage.IMultiSendMessage
	Unsubscribe(item string)
	Subscribe(item string)
}

type IUi interface {
	SetListChange(onListChange func(data []string) bool)
	SetStrategyDataChange(onStrategyDataChange func(name string, data publish.IStrategy) bool)
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

type IAlgoViewer interface {
	io.Closer
	tview.Primitive
	SetData(data interface{})
}
