package marketDataConnection

import (
	"context"
	"github.com/bhbosman/goCommonMarketData/fullMarketDataHelper"
	"github.com/bhbosman/goCommonMarketData/fullMarketDataManagerService"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/model"
	"github.com/bhbosman/gocommon/services/interfaces"
	"github.com/bhbosman/gocomms/intf"
	"github.com/cskr/pubsub"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideConnectionReactor() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Target: func(
					params struct {
					fx.In
					CancelCtx              context.Context
					CancelFunc             context.CancelFunc
					ConnectionCancelFunc   model.ConnectionCancelFunc
					Logger                 *zap.Logger
					PubSub                 *pubsub.PubSub `name:"Application"`
					UniqueReferenceService interfaces.IUniqueReferenceService
					FullMarketDataHelper   fullMarketDataHelper.IFullMarketDataHelper
					GoFunctionCounter      GoFunctionCounter.IService
					FmdService             fullMarketDataManagerService.IFmdManagerService
				},
				) (intf.IConnectionReactor, error) {
					return NewConnectionReactor(
						params.Logger,
						params.CancelCtx,
						params.CancelFunc,
						params.ConnectionCancelFunc,
						params.PubSub,
						params.UniqueReferenceService,
						params.FullMarketDataHelper,
						params.GoFunctionCounter,
						params.FmdService,
					), nil
				},
			},
		),
	)
}
