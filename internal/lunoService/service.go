package lunoService

import (
	"context"
	"github.com/bhbosman/goCommsDefinitions"
	"github.com/bhbosman/gocommon/ChannelHandler"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/Services/IFxService"
	"github.com/bhbosman/gocommon/pubSub"
	"github.com/bhbosman/gocommon/services/ISendMessage"
	"github.com/cskr/pubsub"
	"go.uber.org/zap"
)

type service struct {
	parentContext     context.Context
	ctx               context.Context
	cancelFunc        context.CancelFunc
	cmdChannel        chan interface{}
	onData            func() (ILunoServiceData, error)
	Logger            *zap.Logger
	state             IFxService.State
	pubSub            *pubsub.PubSub
	goFunctionCounter GoFunctionCounter.IService
	subscribeChannel  *pubsub.ChannelSubscription
}

func (self *service) MultiSend(messages ...interface{}) {
	_, err := CallILunoServiceMultiSend(self.ctx, self.cmdChannel, false, messages...)
	if err != nil {
		return
	}
}

func (self *service) Send(message interface{}) error {
	send, err := CallILunoServiceSend(self.ctx, self.cmdChannel, false, message)
	if err != nil {
		return err
	}
	return send.Args0
}

func (self *service) OnStart(ctx context.Context) error {
	err := self.start(ctx)
	if err != nil {
		return err
	}
	self.state = IFxService.Started
	return nil
}

func (self *service) OnStop(ctx context.Context) error {
	err := self.shutdown(ctx)
	close(self.cmdChannel)
	self.state = IFxService.Stopped
	return err
}

func (self *service) shutdown(_ context.Context) error {
	self.cancelFunc()
	return pubSub.Unsubscribe("", self.pubSub, self.goFunctionCounter, self.subscribeChannel)
}

func (self *service) start(_ context.Context) error {
	instanceData, err := self.onData()
	if err != nil {
		return err
	}

	return self.goFunctionCounter.GoRun(
		"FMD Service",
		func() {
			self.goStart(instanceData)
		},
	)
}

func (self *service) goStart(instanceData ILunoServiceData) {
	self.subscribeChannel = pubsub.NewChannelSubscription(32)

	channelHandlerCallback := ChannelHandler.CreateChannelHandlerCallback(
		self.ctx,
		instanceData,
		[]ChannelHandler.ChannelHandler{
			{
				Cb: func(next interface{}, message interface{}) (bool, error) {
					if unk, ok := next.(ILunoService); ok {
						return ChannelEventsForILunoService(unk, message)
					}
					return false, nil

				},
			},
			{
				Cb: func(next interface{}, message interface{}) (bool, error) {
					if unk, ok := next.(ISendMessage.ISendMessage); ok {
						return true, unk.Send(message)
					}
					return false, nil
				},
			},
			// TODO: add handlers here
		},
		func() int {
			return len(self.cmdChannel)
		},
		goCommsDefinitions.CreateTryNextFunc(self.cmdChannel),
	)

	instanceData.Start()
loop:
	for {
		select {
		case <-self.ctx.Done():
			err := instanceData.ShutDown()
			if err != nil {
				self.Logger.Error(
					"error on done",
					zap.Error(err))
			}
			break loop
		case event, ok := <-self.cmdChannel:
			if !ok {
				return
			}
			breakLoop, err := channelHandlerCallback(event)
			if err != nil || breakLoop {
				break loop
			}
		case event, ok := <-self.subscribeChannel.Data:
			if !ok {
				return
			}
			breakLoop, err := channelHandlerCallback(event)
			if err != nil || breakLoop {
				break loop
			}
		}
	}
}

func (self *service) State() IFxService.State {
	return self.state
}

func (self service) ServiceName() string {
	return "LunoService"
}

func newService(
	parentContext context.Context,
	onData func() (ILunoServiceData, error),
	logger *zap.Logger,
	pubSub *pubsub.PubSub,
	goFunctionCounter GoFunctionCounter.IService,
) (ILunoServiceService, error) {
	localCtx, localCancelFunc := context.WithCancel(parentContext)
	return &service{
		parentContext:     parentContext,
		ctx:               localCtx,
		cancelFunc:        localCancelFunc,
		cmdChannel:        make(chan interface{}, 32),
		onData:            onData,
		Logger:            logger,
		pubSub:            pubSub,
		goFunctionCounter: goFunctionCounter,
	}, nil
}
