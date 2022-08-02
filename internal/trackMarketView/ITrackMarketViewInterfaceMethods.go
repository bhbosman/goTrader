// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bhbosman/goTrader/internal/trackMarketView (interfaces: ITrackMarketView)

// Package trackMarketView is a generated GoMock package.
package trackMarketView

import (
	fmt "fmt"

	errors "github.com/bhbosman/gocommon/errors"
	"golang.org/x/net/context"
)

// Interface A Comment
// Interface github.com/bhbosman/goTrader/internal/trackMarketView
// Interface ITrackMarketView
// Interface ITrackMarketView, Method: MultiSend
type ITrackMarketViewMultiSendIn struct {
	arg0 []interface{}
}

type ITrackMarketViewMultiSendOut struct {
}
type ITrackMarketViewMultiSendError struct {
	InterfaceName string
	MethodName    string
	Reason        string
}

func (self *ITrackMarketViewMultiSendError) Error() string {
	return fmt.Sprintf("error in data coming back from %v::%v. Reason: %v", self.InterfaceName, self.MethodName, self.Reason)
}

type ITrackMarketViewMultiSend struct {
	inData         ITrackMarketViewMultiSendIn
	outDataChannel chan ITrackMarketViewMultiSendOut
}

func NewITrackMarketViewMultiSend(waitToComplete bool, arg0 ...interface{}) *ITrackMarketViewMultiSend {
	var outDataChannel chan ITrackMarketViewMultiSendOut
	if waitToComplete {
		outDataChannel = make(chan ITrackMarketViewMultiSendOut)
	} else {
		outDataChannel = nil
	}
	return &ITrackMarketViewMultiSend{
		inData: ITrackMarketViewMultiSendIn{
			arg0: arg0,
		},
		outDataChannel: outDataChannel,
	}
}

func (self *ITrackMarketViewMultiSend) Wait(onError func(interfaceName string, methodName string, err error) error) (ITrackMarketViewMultiSendOut, error) {
	data, ok := <-self.outDataChannel
	if !ok {
		generatedError := &ITrackMarketViewMultiSendError{
			InterfaceName: "ITrackMarketView",
			MethodName:    "MultiSend",
			Reason:        "Channel for ITrackMarketView::MultiSend returned false",
		}
		if onError != nil {
			err := onError("ITrackMarketView", "MultiSend", generatedError)
			return ITrackMarketViewMultiSendOut{}, err
		} else {
			return ITrackMarketViewMultiSendOut{}, generatedError
		}
	}
	return data, nil
}

func (self *ITrackMarketViewMultiSend) Close() error {
	close(self.outDataChannel)
	return nil
}
func CallITrackMarketViewMultiSend(context context.Context, channel chan<- interface{}, waitToComplete bool, arg0 ...interface{}) (ITrackMarketViewMultiSendOut, error) {
	if context != nil && context.Err() != nil {
		return ITrackMarketViewMultiSendOut{}, context.Err()
	}
	data := NewITrackMarketViewMultiSend(waitToComplete, arg0...)
	if waitToComplete {
		defer func(data *ITrackMarketViewMultiSend) {
			err := data.Close()
			if err != nil {
			}
		}(data)
	}
	if context != nil && context.Err() != nil {
		return ITrackMarketViewMultiSendOut{}, context.Err()
	}
	channel <- data
	var err error
	var v ITrackMarketViewMultiSendOut
	if waitToComplete {
		v, err = data.Wait(func(interfaceName string, methodName string, err error) error {
			return err
		})
	} else {
		err = errors.NoWaitOperationError
	}
	if err != nil {
		return ITrackMarketViewMultiSendOut{}, err
	}
	return v, nil
}

// Interface ITrackMarketView, Method: Send
type ITrackMarketViewSendIn struct {
	arg0 interface{}
}

type ITrackMarketViewSendOut struct {
	Args0 error
}
type ITrackMarketViewSendError struct {
	InterfaceName string
	MethodName    string
	Reason        string
}

func (self *ITrackMarketViewSendError) Error() string {
	return fmt.Sprintf("error in data coming back from %v::%v. Reason: %v", self.InterfaceName, self.MethodName, self.Reason)
}

type ITrackMarketViewSend struct {
	inData         ITrackMarketViewSendIn
	outDataChannel chan ITrackMarketViewSendOut
}

func NewITrackMarketViewSend(waitToComplete bool, arg0 interface{}) *ITrackMarketViewSend {
	var outDataChannel chan ITrackMarketViewSendOut
	if waitToComplete {
		outDataChannel = make(chan ITrackMarketViewSendOut)
	} else {
		outDataChannel = nil
	}
	return &ITrackMarketViewSend{
		inData: ITrackMarketViewSendIn{
			arg0: arg0,
		},
		outDataChannel: outDataChannel,
	}
}

func (self *ITrackMarketViewSend) Wait(onError func(interfaceName string, methodName string, err error) error) (ITrackMarketViewSendOut, error) {
	data, ok := <-self.outDataChannel
	if !ok {
		generatedError := &ITrackMarketViewSendError{
			InterfaceName: "ITrackMarketView",
			MethodName:    "Send",
			Reason:        "Channel for ITrackMarketView::Send returned false",
		}
		if onError != nil {
			err := onError("ITrackMarketView", "Send", generatedError)
			return ITrackMarketViewSendOut{}, err
		} else {
			return ITrackMarketViewSendOut{}, generatedError
		}
	}
	return data, nil
}

func (self *ITrackMarketViewSend) Close() error {
	close(self.outDataChannel)
	return nil
}
func CallITrackMarketViewSend(context context.Context, channel chan<- interface{}, waitToComplete bool, arg0 interface{}) (ITrackMarketViewSendOut, error) {
	if context != nil && context.Err() != nil {
		return ITrackMarketViewSendOut{}, context.Err()
	}
	data := NewITrackMarketViewSend(waitToComplete, arg0)
	if waitToComplete {
		defer func(data *ITrackMarketViewSend) {
			err := data.Close()
			if err != nil {
			}
		}(data)
	}
	if context != nil && context.Err() != nil {
		return ITrackMarketViewSendOut{}, context.Err()
	}
	channel <- data
	var err error
	var v ITrackMarketViewSendOut
	if waitToComplete {
		v, err = data.Wait(func(interfaceName string, methodName string, err error) error {
			return err
		})
	} else {
		err = errors.NoWaitOperationError
	}
	if err != nil {
		return ITrackMarketViewSendOut{}, err
	}
	return v, nil
}

func ChannelEventsForITrackMarketView(next ITrackMarketView, event interface{}) (bool, error) {
	switch v := event.(type) {
	case *ITrackMarketViewMultiSend:
		data := ITrackMarketViewMultiSendOut{}
		next.MultiSend(v.inData.arg0...)
		if v.outDataChannel != nil {
			v.outDataChannel <- data
		}
		return true, nil
	case *ITrackMarketViewSend:
		data := ITrackMarketViewSendOut{}
		data.Args0 = next.Send(v.inData.arg0)
		if v.outDataChannel != nil {
			v.outDataChannel <- data
		}
		return true, nil
	default:
		return false, nil
	}
}
