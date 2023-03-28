package strategyStateManagerView

import (
	"github.com/bhbosman/goTrader/internal/strategyStateManagerService"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/services/interfaces"
	"github.com/cskr/pubsub"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func Provide() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Name:  "",
				Group: "",
				Target: func(
					params struct {
						fx.In
					},
				) (func() (ITrackMarketViewData, error), error) {
					return func() (ITrackMarketViewData, error) {
						return newData()
					}, nil
				},
			},
		),
		fx.Provide(
			fx.Annotated{
				Name:  "",
				Group: "",
				Target: func(
					params struct {
						fx.In
						PubSub                 *pubsub.PubSub  `name:"Application"`
						ApplicationContext     context.Context `name:"Application"`
						OnData                 func() (ITrackMarketViewData, error)
						Lifecycle              fx.Lifecycle
						Logger                 *zap.Logger
						UniqueReferenceService interfaces.IUniqueReferenceService
						UniqueSessionNumber    interfaces.IUniqueSessionNumber
						GoFunctionCounter      GoFunctionCounter.IService
						StrategyManagerService strategyStateManagerService.IStrategyManagerService
					},
				) (ITrackMarketViewService, error) {
					serviceInstance, err := newService(
						params.ApplicationContext,
						params.OnData,
						params.Logger,
						params.PubSub,
						params.GoFunctionCounter,
						params.StrategyManagerService,
					)
					if err != nil {
						return nil, err
					}
					params.Lifecycle.Append(
						fx.Hook{
							OnStart: serviceInstance.OnStart,
							OnStop:  serviceInstance.OnStop,
						},
					)
					return serviceInstance, nil
				},
			},
		),
	)
}
