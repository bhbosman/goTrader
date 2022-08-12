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

// Interface ITrackMarketView, Method: Subscribe
type ITrackMarketViewSubscribeIn struct {
	arg0 string
}

type ITrackMarketViewSubscribeOut struct {
}
type ITrackMarketViewSubscribeError struct {
	InterfaceName string
	MethodName    string
	Reason        string
}

func (self *ITrackMarketViewSubscribeError) Error() string {
	return fmt.Sprintf("error in data coming back from %v::%v. Reason: %v", self.InterfaceName, self.MethodName, self.Reason)
}

type ITrackMarketViewSubscribe struct {
	inData         ITrackMarketViewSubscribeIn
	outDataChannel chan ITrackMarketViewSubscribeOut
}

func NewITrackMarketViewSubscribe(waitToComplete bool, arg0 string) *ITrackMarketViewSubscribe {
	var outDataChannel chan ITrackMarketViewSubscribeOut
	if waitToComplete {
		outDataChannel = make(chan ITrackMarketViewSubscribeOut)
	} else {
		outDataChannel = nil
	}
	return &ITrackMarketViewSubscribe{
		inData: ITrackMarketViewSubscribeIn{
			arg0: arg0,
		},
		outDataChannel: outDataChannel,
	}
}

func (self *ITrackMarketViewSubscribe) Wait(onError func(interfaceName string, methodName string, err error) error) (ITrackMarketViewSubscribeOut, error) {
	data, ok := <-self.outDataChannel
	if !ok {
		generatedError := &ITrackMarketViewSubscribeError{
			InterfaceName: "ITrackMarketView",
			MethodName:    "Subscribe",
			Reason:        "Channel for ITrackMarketView::Subscribe returned false",
		}
		if onError != nil {
			err := onError("ITrackMarketView", "Subscribe", generatedError)
			return ITrackMarketViewSubscribeOut{}, err
		} else {
			return ITrackMarketViewSubscribeOut{}, generatedError
		}
	}
	return data, nil
}

func (self *ITrackMarketViewSubscribe) Close() error {
	close(self.outDataChannel)
	return nil
}
func CallITrackMarketViewSubscribe(context context.Context, channel chan<- interface{}, waitToComplete bool, arg0 string) (ITrackMarketViewSubscribeOut, error) {
	if context != nil && context.Err() != nil {
		return ITrackMarketViewSubscribeOut{}, context.Err()
	}
	data := NewITrackMarketViewSubscribe(waitToComplete, arg0)
	if waitToComplete {
		defer func(data *ITrackMarketViewSubscribe) {
			err := data.Close()
			if err != nil {
			}
		}(data)
	}
	if context != nil && context.Err() != nil {
		return ITrackMarketViewSubscribeOut{}, context.Err()
	}
	channel <- data
	var err error
	var v ITrackMarketViewSubscribeOut
	if waitToComplete {
		v, err = data.Wait(func(interfaceName string, methodName string, err error) error {
			return err
		})
	} else {
		err = errors.NoWaitOperationError
	}
	if err != nil {
		return ITrackMarketViewSubscribeOut{}, err
	}
	return v, nil
}

// Interface ITrackMarketView, Method: Unsubscribe
type ITrackMarketViewUnsubscribeIn struct {
	arg0 string
}

type ITrackMarketViewUnsubscribeOut struct {
}
type ITrackMarketViewUnsubscribeError struct {
	InterfaceName string
	MethodName    string
	Reason        string
}

func (self *ITrackMarketViewUnsubscribeError) Error() string {
	return fmt.Sprintf("error in data coming back from %v::%v. Reason: %v", self.InterfaceName, self.MethodName, self.Reason)
}

type ITrackMarketViewUnsubscribe struct {
	inData         ITrackMarketViewUnsubscribeIn
	outDataChannel chan ITrackMarketViewUnsubscribeOut
}

func NewITrackMarketViewUnsubscribe(waitToComplete bool, arg0 string) *ITrackMarketViewUnsubscribe {
	var outDataChannel chan ITrackMarketViewUnsubscribeOut
	if waitToComplete {
		outDataChannel = make(chan ITrackMarketViewUnsubscribeOut)
	} else {
		outDataChannel = nil
	}
	return &ITrackMarketViewUnsubscribe{
		inData: ITrackMarketViewUnsubscribeIn{
			arg0: arg0,
		},
		outDataChannel: outDataChannel,
	}
}

func (self *ITrackMarketViewUnsubscribe) Wait(onError func(interfaceName string, methodName string, err error) error) (ITrackMarketViewUnsubscribeOut, error) {
	data, ok := <-self.outDataChannel
	if !ok {
		generatedError := &ITrackMarketViewUnsubscribeError{
			InterfaceName: "ITrackMarketView",
			MethodName:    "Unsubscribe",
			Reason:        "Channel for ITrackMarketView::Unsubscribe returned false",
		}
		if onError != nil {
			err := onError("ITrackMarketView", "Unsubscribe", generatedError)
			return ITrackMarketViewUnsubscribeOut{}, err
		} else {
			return ITrackMarketViewUnsubscribeOut{}, generatedError
		}
	}
	return data, nil
}

func (self *ITrackMarketViewUnsubscribe) Close() error {
	close(self.outDataChannel)
	return nil
}
func CallITrackMarketViewUnsubscribe(context context.Context, channel chan<- interface{}, waitToComplete bool, arg0 string) (ITrackMarketViewUnsubscribeOut, error) {
	if context != nil && context.Err() != nil {
		return ITrackMarketViewUnsubscribeOut{}, context.Err()
	}
	data := NewITrackMarketViewUnsubscribe(waitToComplete, arg0)
	if waitToComplete {
		defer func(data *ITrackMarketViewUnsubscribe) {
			err := data.Close()
			if err != nil {
			}
		}(data)
	}
	if context != nil && context.Err() != nil {
		return ITrackMarketViewUnsubscribeOut{}, context.Err()
	}
	channel <- data
	var err error
	var v ITrackMarketViewUnsubscribeOut
	if waitToComplete {
		v, err = data.Wait(func(interfaceName string, methodName string, err error) error {
			return err
		})
	} else {
		err = errors.NoWaitOperationError
	}
	if err != nil {
		return ITrackMarketViewUnsubscribeOut{}, err
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
	case *ITrackMarketViewSubscribe:
		data := ITrackMarketViewSubscribeOut{}
		next.Subscribe(v.inData.arg0)
		if v.outDataChannel != nil {
			v.outDataChannel <- data
		}
		return true, nil
	case *ITrackMarketViewUnsubscribe:
		data := ITrackMarketViewUnsubscribeOut{}
		next.Unsubscribe(v.inData.arg0)
		if v.outDataChannel != nil {
			v.outDataChannel <- data
		}
		return true, nil
	default:
		return false, nil
	}
}
