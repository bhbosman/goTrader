package strategyStateManagerService

import (
	"github.com/bhbosman/gocommon/services/IDataShutDown"
	"github.com/bhbosman/gocommon/services/IFxService"
	"github.com/bhbosman/gocommon/services/ISendMessage"
)

type IStrategyStateManager interface {
	ISendMessage.ISendMessage
	ISendMessage.IMultiSendMessage
}

type IStrategyManagerService interface {
	IStrategyStateManager
	IFxService.IFxServices
}

type IStrategyManagerData interface {
	IStrategyStateManager
	IDataShutDown.IDataShutDown
}
