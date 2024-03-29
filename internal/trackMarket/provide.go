package trackMarket

import (
	"context"
	"github.com/bhbosman/goCommonMarketData/fullMarketDataHelper"
	"github.com/bhbosman/goCommonMarketData/fullMarketDataManagerService"
	"github.com/bhbosman/goConn"
	fxAppManager "github.com/bhbosman/goFxAppManager/service"
	"github.com/bhbosman/goTrader/internal/lunoService"
	"github.com/bhbosman/goTrader/internal/strategyStateManagerService"
	"github.com/bhbosman/goTrader/internal/strategyStateManagerView"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/messages"
	"github.com/bhbosman/gocommon/services/interfaces"
	"github.com/cskr/pubsub"
	"github.com/openlyinc/pointy"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type OnITrackMarketServiceCreate func(modelSettings IPricingVolumeCalculation) (ITrackMarketService, error)
type OnITrackMarketDataCreate func(modelSettings IPricingVolumeCalculation) (ITrackMarketData, error)

func Provide() fx.Option {
	return fx.Options(
		fx.Provide(
			func(
				params struct {
					fx.In
					ApplicationContext     context.Context `name:"Application"`
					TrackMarketViewService strategyStateManagerView.ITrackMarketViewService
					FullMarketDataHelper   fullMarketDataHelper.IFullMarketDataHelper
					FmdService             fullMarketDataManagerService.IFmdManagerService
					PubSub                 *pubsub.PubSub `name:"Application"`
					LunoServiceService     lunoService.ILunoServiceService
					StrategyManagerService strategyStateManagerService.IStrategyManagerService
				},
			) (OnITrackMarketDataCreate, error) {
				return func(modelSettings IPricingVolumeCalculation) (ITrackMarketData, error) {
					return newData(
						params.ApplicationContext,
						params.StrategyManagerService,
						modelSettings,
						params.FmdService,
						params.FullMarketDataHelper,
						params.PubSub,
						params.LunoServiceService,
					)
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
					},
				) (OnITrackMarketServiceCreate, error) {
					return func(modelSettings IPricingVolumeCalculation) (ITrackMarketService, error) {
						serviceInstance, err := newService(
							params.ApplicationContext,
							params.OnData,
							params.Logger,
							params.PubSub,
							params.GoFunctionCounter,
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
					Logger                      *zap.Logger
					ApplicationContext          context.Context `name:"Application"`
				},
			) error {
				Strategies := Strategies{
					PegToPrice: []*PegToPrice{
						{
							strategyName:     "Peg Luno.XBTGBP BID @ 10,000",
							instrument:       "Luno.XBTGBP",
							Side:             "BID",
							PegPrice:         10000,
							Volume:           nil,
							Consideration:    pointy.Float64(1000),
							TradingInterface: "",
						},
						{
							strategyName:     "Peg Luno.XBTZAR BID @ 150,000",
							instrument:       "Luno.XBTZAR",
							Side:             "BID",
							PegPrice:         150000,
							Volume:           nil,
							Consideration:    pointy.Float64(1000),
							TradingInterface: "",
						},
						{
							strategyName:     "Peg Luno.XBTZAR BID @ 100,000",
							instrument:       "Luno.XBTZAR",
							Side:             "BID",
							PegPrice:         100000,
							Volume:           nil,
							Consideration:    pointy.Float64(1000),
							TradingInterface: "",
						},
					},
					modelSettings: []modelSettings{
						{
							Name:       "Luno.XBTZAR 01",
							instrument: "Luno.XBTZAR",
						},
					},
				}
				for _, m := range Strategies.PegToPrice {
					localModel := m
					params.Lifecycle.Append(
						fx.Hook{
							OnStart: func(ctx context.Context) error {
								return params.FxManagerService.Add(localModel.strategyName,
									func() (messages.IApp, goConn.ICancellationContext, error) {
										trackMarketService, err := params.OnITrackMarketServiceCreate(localModel)
										if err != nil {
											return nil, nil, err
										}
										app := newAppWrapper(trackMarketService)
										namedLogger := params.Logger.Named(m.strategyName)
										ctx, cancelFunc := context.WithCancel(params.ApplicationContext)
										cancellationContext, err := goConn.NewCancellationContextNoCloser(
											m.strategyName,
											cancelFunc,
											ctx,
											namedLogger,
										)

										_ = goConn.RegisterConnectionShutdown(
											m.strategyName,
											func() {

											},
											cancellationContext)

										return app, cancellationContext, nil
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
