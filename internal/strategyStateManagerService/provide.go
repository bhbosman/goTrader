package strategyStateManagerService

import (
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/Services/interfaces"
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
						PubSub *pubsub.PubSub `name:"Application"`
					},
				) (func() (IStrategyManagerData, error), error) {
					return func() (IStrategyManagerData, error) {
						return newData(params.PubSub)
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
						OnData                 func() (IStrategyManagerData, error)
						Lifecycle              fx.Lifecycle
						Logger                 *zap.Logger
						UniqueReferenceService interfaces.IUniqueReferenceService
						UniqueSessionNumber    interfaces.IUniqueSessionNumber
						GoFunctionCounter      GoFunctionCounter.IService
					},
				) (IStrategyManagerService, error) {
					serviceInstance, err := newService(
						params.ApplicationContext,
						params.OnData,
						params.Logger,
						params.PubSub,
						params.GoFunctionCounter,
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
