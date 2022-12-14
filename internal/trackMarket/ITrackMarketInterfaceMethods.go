// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/bhbosman/goTrader/internal/trackMarket (interfaces: ITrackMarket)

// Package trackMarket is a generated GoMock package.
package trackMarket

import (
	"context"
	fmt "fmt"

	errors "github.com/bhbosman/gocommon/errors"
)

// Interface A Comment
// Interface github.com/bhbosman/goTrader/internal/trackMarket
// Interface ITrackMarket
// Interface ITrackMarket, Method: MultiSend
type ITrackMarketMultiSendIn struct {
	arg0 []interface{}
}

type ITrackMarketMultiSendOut struct {
}
type ITrackMarketMultiSendError struct {
	InterfaceName string
	MethodName    string
	Reason        string
}

func (self *ITrackMarketMultiSendError) Error() string {
	return fmt.Sprintf("error in data coming back from %v::%v. Reason: %v", self.InterfaceName, self.MethodName, self.Reason)
}

type ITrackMarketMultiSend struct {
	inData         ITrackMarketMultiSendIn
	outDataChannel chan ITrackMarketMultiSendOut
}

func NewITrackMarketMultiSend(waitToComplete bool, arg0 ...interface{}) *ITrackMarketMultiSend {
	var outDataChannel chan ITrackMarketMultiSendOut
	if waitToComplete {
		outDataChannel = make(chan ITrackMarketMultiSendOut)
	} else {
		outDataChannel = nil
	}
	return &ITrackMarketMultiSend{
		inData: ITrackMarketMultiSendIn{
			arg0: arg0,
		},
		outDataChannel: outDataChannel,
	}
}

func (self *ITrackMarketMultiSend) Wait(onError func(interfaceName string, methodName string, err error) error) (ITrackMarketMultiSendOut, error) {
	data, ok := <-self.outDataChannel
	if !ok {
		generatedError := &ITrackMarketMultiSendError{
			InterfaceName: "ITrackMarket",
			MethodName:    "MultiSend",
			Reason:        "Channel for ITrackMarket::MultiSend returned false",
		}
		if onError != nil {
			err := onError("ITrackMarket", "MultiSend", generatedError)
			return ITrackMarketMultiSendOut{}, err
		} else {
			return ITrackMarketMultiSendOut{}, generatedError
		}
	}
	return data, nil
}

func (self *ITrackMarketMultiSend) Close() error {
	close(self.outDataChannel)
	return nil
}
func CallITrackMarketMultiSend(context context.Context, channel chan<- interface{}, waitToComplete bool, arg0 ...interface{}) (ITrackMarketMultiSendOut, error) {
	if context != nil && context.Err() != nil {
		return ITrackMarketMultiSendOut{}, context.Err()
	}
	data := NewITrackMarketMultiSend(waitToComplete, arg0...)
	if waitToComplete {
		defer func(data *ITrackMarketMultiSend) {
			err := data.Close()
			if err != nil {
			}
		}(data)
	}
	if context != nil && context.Err() != nil {
		return ITrackMarketMultiSendOut{}, context.Err()
	}
	channel <- data
	var err error
	var v ITrackMarketMultiSendOut
	if waitToComplete {
		v, err = data.Wait(func(interfaceName string, methodName string, err error) error {
			return err
		})
	} else {
		err = errors.NoWaitOperationError
	}
	if err != nil {
		return ITrackMarketMultiSendOut{}, err
	}
	return v, nil
}

// Interface ITrackMarket, Method: Send
type ITrackMarketSendIn struct {
	arg0 interface{}
}

type ITrackMarketSendOut struct {
	Args0 error
}
type ITrackMarketSendError struct {
	InterfaceName string
	MethodName    string
	Reason        string
}

func (self *ITrackMarketSendError) Error() string {
	return fmt.Sprintf("error in data coming back from %v::%v. Reason: %v", self.InterfaceName, self.MethodName, self.Reason)
}

type ITrackMarketSend struct {
	inData         ITrackMarketSendIn
	outDataChannel chan ITrackMarketSendOut
}

func NewITrackMarketSend(waitToComplete bool, arg0 interface{}) *ITrackMarketSend {
	var outDataChannel chan ITrackMarketSendOut
	if waitToComplete {
		outDataChannel = make(chan ITrackMarketSendOut)
	} else {
		outDataChannel = nil
	}
	return &ITrackMarketSend{
		inData: ITrackMarketSendIn{
			arg0: arg0,
		},
		outDataChannel: outDataChannel,
	}
}

func (self *ITrackMarketSend) Wait(onError func(interfaceName string, methodName string, err error) error) (ITrackMarketSendOut, error) {
	data, ok := <-self.outDataChannel
	if !ok {
		generatedError := &ITrackMarketSendError{
			InterfaceName: "ITrackMarket",
			MethodName:    "Send",
			Reason:        "Channel for ITrackMarket::Send returned false",
		}
		if onError != nil {
			err := onError("ITrackMarket", "Send", generatedError)
			return ITrackMarketSendOut{}, err
		} else {
			return ITrackMarketSendOut{}, generatedError
		}
	}
	return data, nil
}

func (self *ITrackMarketSend) Close() error {
	close(self.outDataChannel)
	return nil
}
func CallITrackMarketSend(context context.Context, channel chan<- interface{}, waitToComplete bool, arg0 interface{}) (ITrackMarketSendOut, error) {
	if context != nil && context.Err() != nil {
		return ITrackMarketSendOut{}, context.Err()
	}
	data := NewITrackMarketSend(waitToComplete, arg0)
	if waitToComplete {
		defer func(data *ITrackMarketSend) {
			err := data.Close()
			if err != nil {
			}
		}(data)
	}
	if context != nil && context.Err() != nil {
		return ITrackMarketSendOut{}, context.Err()
	}
	channel <- data
	var err error
	var v ITrackMarketSendOut
	if waitToComplete {
		v, err = data.Wait(func(interfaceName string, methodName string, err error) error {
			return err
		})
	} else {
		err = errors.NoWaitOperationError
	}
	if err != nil {
		return ITrackMarketSendOut{}, err
	}
	return v, nil
}

func ChannelEventsForITrackMarket(next ITrackMarket, event interface{}) (bool, error) {
	switch v := event.(type) {
	case *ITrackMarketMultiSend:
		data := ITrackMarketMultiSendOut{}
		next.MultiSend(v.inData.arg0...)
		if v.outDataChannel != nil {
			v.outDataChannel <- data
		}
		return true, nil
	case *ITrackMarketSend:
		data := ITrackMarketSendOut{}
		data.Args0 = next.Send(v.inData.arg0)
		if v.outDataChannel != nil {
			v.outDataChannel <- data
		}
		return true, nil
	default:
		return false, nil
	}
}
