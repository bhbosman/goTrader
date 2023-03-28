package lunoService

import (
	"github.com/bhbosman/goTrader/internal/lunoApi/client"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/services/interfaces"
	"github.com/cskr/pubsub"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/net/context"
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
				Name: "Luno",
				Target: func(
					params struct {
						fx.In
						RequestDoer client.HttpRequestDoer `name:"Luno"`
					},
				) (client.ClientWithResponsesInterface, error) {
					lunoClient, err := client.NewClientWithResponses(
						"https://api.luno.com",
						func(c *client.Client) error {
							c.Client = params.RequestDoer
							return nil
						},
					)
					if err != nil {
						return nil, err
					}
					return lunoClient, nil
				},
			},
		),

		fx.Provide(
			fx.Annotated{
				Target: func(
					params struct {
						fx.In
						ClientWithResponses client.ClientWithResponsesInterface `name:"Luno"`
					},
				) (func() (ILunoServiceData, error), error) {
					return func() (ILunoServiceData, error) {
						return newData(
							params.ClientWithResponses,
							//params.RoundTripper,
							//params.RequestDoer,
						)
					}, nil
				},
			},
		),
		fx.Provide(
			fx.Annotated{
				Target: func(
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
					},
				) (ILunoServiceService, error) {
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
