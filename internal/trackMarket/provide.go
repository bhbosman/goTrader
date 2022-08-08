package trackMarket

import (
	"context"
	"github.com/bhbosman/goCommonMarketData/fullMarketDataHelper"
	"github.com/bhbosman/goCommonMarketData/fullMarketDataManagerService"
	fxAppManager "github.com/bhbosman/goFxAppManager/service"
	"github.com/bhbosman/goTrader/internal/trackMarketView"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/Services/interfaces"
	"github.com/bhbosman/gocommon/messages"
	"github.com/cskr/pubsub"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type OnITrackMarketServiceCreate func(modelSettings modelSettings) (ITrackMarketService, error)
type OnITrackMarketDataCreate func(modelSettings modelSettings) (ITrackMarketData, error)

func Provide() fx.Option {
	return fx.Options(
		fx.Provide(
			func(
				params struct {
					fx.In
					TrackMarketViewService trackMarketView.ITrackMarketViewService
				},
			) (OnITrackMarketDataCreate, error) {
				return func(modelSettings modelSettings) (ITrackMarketData, error) {
					return newData(params.TrackMarketViewService, modelSettings)
				}, nil
			},
		),
		fx.Provide(
			fx.Annotated{
				Target: func(
					params struct {
						fx.In
						PubSub                 *pubsub.PubSub  `name:"Application"`
						ApplicationContext     context.Context `name:"Application"`
						OnData                 OnITrackMarketDataCreate
						Logger                 *zap.Logger
						UniqueReferenceService interfaces.IUniqueReferenceService
						UniqueSessionNumber    interfaces.IUniqueSessionNumber
						GoFunctionCounter      GoFunctionCounter.IService
						FullMarketDataHelper   fullMarketDataHelper.IFullMarketDataHelper
						FmdService             fullMarketDataManagerService.IFmdManagerService
					},
				) (OnITrackMarketServiceCreate, error) {
					return func(modelSettings modelSettings) (ITrackMarketService, error) {
						serviceInstance, err := newService(
							params.ApplicationContext,
							params.OnData,
							params.Logger,
							params.PubSub,
							params.GoFunctionCounter,
							params.FullMarketDataHelper,
							params.FmdService,
							modelSettings,
						)
						if err != nil {
							return nil, err
						}
						return serviceInstance, nil
					}, nil
				},
			},
		),
		fx.Invoke(
			func(
				params struct {
					fx.In
					Lifecycle                   fx.Lifecycle
					OnITrackMarketServiceCreate OnITrackMarketServiceCreate
					FxManagerService            fxAppManager.IFxManagerService
				},
			) error {
				models := []modelSettings{
					{
						Name:       "Luno.XBTZAR 01",
						instrument: "Luno.XBTZAR",
					},
					{
						Name:       "Luno.XBTZAR 02",
						instrument: "Luno.XBTZAR",
					},
					{
						Name:       "Luno.XBTZAR 03",
						instrument: "Luno.XBTZAR",
					},
				}
				for _, m := range models {
					localModel := m
					params.Lifecycle.Append(
						fx.Hook{
							OnStart: func(ctx context.Context) error {
								return params.FxManagerService.Add(localModel.Name,
									func() (messages.IApp, context.CancelFunc, error) {
										trackMarketService, err := params.OnITrackMarketServiceCreate(localModel)
										if err != nil {
											return nil, nil, err
										}
										app := newAppWrapper(trackMarketService)
										return app, func() {}, nil
									},
								)
							},
							OnStop: nil,
						},
					)
				}
				return nil
			},
		),
	)
}

type appWrapper struct {
	trackMarketService ITrackMarketService
	err                error
}

func newAppWrapper(trackMarketService ITrackMarketService) *appWrapper {
	return &appWrapper{trackMarketService: trackMarketService}
}

func (self *appWrapper) Start(ctx context.Context) error {
	self.err = self.trackMarketService.OnStart(ctx)
	return self.err
}

func (self *appWrapper) Stop(ctx context.Context) error {
	return self.trackMarketService.OnStop(ctx)
}

func (self *appWrapper) Err() error {
	return self.err
}
