package lunoService

import (
	"github.com/bhbosman/gocommon/services/IDataShutDown"
	"github.com/bhbosman/gocommon/services/IFxService"
	"github.com/bhbosman/gocommon/services/ISendMessage"
)

type ILunoService interface {
	ISendMessage.ISendMessage
	ISendMessage.IMultiSendMessage
	IExchange
}

type ILunoServiceService interface {
	ILunoService
	IFxService.IFxServices
}

type ILunoServiceData interface {
	ILunoService
	IDataShutDown.IDataShutDown
	Start()
}
