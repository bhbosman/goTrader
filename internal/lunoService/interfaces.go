package lunoService

import (
	"github.com/bhbosman/gocommon/Services/IDataShutDown"
	"github.com/bhbosman/gocommon/Services/IFxService"
	"github.com/bhbosman/gocommon/services/ISendMessage"
)

type ILunoService interface {
	ISendMessage.ISendMessage
	ISendMessage.IMultiSendMessage
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
