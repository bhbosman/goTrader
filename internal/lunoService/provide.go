package lunoService

import (
	"context"
	service2 "github.com/bhbosman/goFxAppManager/service"
	"github.com/bhbosman/goTrader/internal/lunoApi/client"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/Services/interfaces"
	"github.com/cskr/pubsub"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
)

func Provide() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Name: "Default",
				Target: func(
					params struct {
						fx.In
					},
				) (http.RoundTripper, error) {
					return http.DefaultTransport, nil
				},
			},
		),
		fx.Provide(
			fx.Annotated{
				Name: "Luno",
				Target: func(
					params struct {
						fx.In
						Keys         *LunoKeys
						RoundTripper http.RoundTripper `name:"Default"`
					},
				) (http.RoundTripper, error) {
					return newLunoTransport(params.Keys.Key, params.Keys.Secret, params.RoundTripper)
				},
			},
		),
		fx.Provide(
			fx.Annotated{
				Name: "Luno",
				Target: func(
					params struct {
						fx.In
						RoundTripper http.RoundTripper `name:"Luno"`
					},
				) (client.HttpRequestDoer, error) {
					return newLunoHttpRequestDoer(params.RoundTripper)
				},
			},
		),

		fx.Provide(
			fx.Annotated{
				Target: func(
					params struct {
						fx.In
						RoundTripper http.RoundTripper      `name:"Luno"`
						RequestDoer  client.HttpRequestDoer `name:"Luno"`
					},
				) (func() (ILunoServiceData, error), error) {
					return func() (ILunoServiceData, error) {
						return newData(params.RoundTripper, params.RequestDoer)
					}, nil
				},
			},
		),
		fx.Invoke(
			func(
				params struct {
					fx.In
					PubSub                 *pubsub.PubSub  `name:"Application"`
					ApplicationContext     context.Context `name:"Application"`
					OnData                 func() (ILunoServiceData, error)
					Lifecycle              fx.Lifecycle
					Logger                 *zap.Logger
					UniqueReferenceService interfaces.IUniqueReferenceService
					UniqueSessionNumber    interfaces.IUniqueSessionNumber
					GoFunctionCounter      GoFunctionCounter.IService
					RoundTripper           http.RoundTripper `name:"Luno"`
					FxManagerService       service2.IFxManagerService
				},
			) error {
				params.Lifecycle.Append(
					fx.Hook{
						OnStart: func(ctx context.Context) error {
							serviceInstance, err := newService(
								params.ApplicationContext,
								params.OnData,
								params.Logger,
								params.PubSub,
								params.GoFunctionCounter,
							)
							if err != nil {
								return nil
							}
							params.Lifecycle.Append(
								fx.Hook{
									OnStart: serviceInstance.OnStart,
									OnStop:  serviceInstance.OnStop,
								},
							)
							return nil
						},
						OnStop: nil,
					})
				return nil
			},
		),
	)
}
