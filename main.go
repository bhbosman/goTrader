package main

import (
	"github.com/bhbosman/goCommonMarketData/fullMarketDataHelper"
	"github.com/bhbosman/goCommonMarketData/fullMarketDataManagerService"
	"github.com/bhbosman/goCommonMarketData/fullMarketDataManagerViewer"
	"github.com/bhbosman/goCommonMarketData/instrumentReference"
	"github.com/bhbosman/goFxApp"
	"github.com/bhbosman/goTrader/internal/lunoService"
	"github.com/bhbosman/goTrader/internal/marketDataConnection"
	"github.com/bhbosman/goTrader/internal/strategyStateManagerService"
	"github.com/bhbosman/goTrader/internal/strategyStateManagerView"
	"github.com/bhbosman/goTrader/internal/trackMarket"
)

func main() {
	app := goFxApp.NewFxMainApplicationServices(
		"Trader",
		false,
		strategyStateManagerService.Provide(),
		marketDataConnection.ProvideMarketDataDialer(1, "tcp4://127.0.0.1:4001"),
		fullMarketDataManagerViewer.Provide(),
		lunoService.Provide(),
		fullMarketDataManagerService.Provide(true),
		fullMarketDataHelper.Provide(),
		instrumentReference.Provide(),

		trackMarket.Provide(),
		strategyStateManagerView.Provide(),
		strategyStateManagerView.ProvideView(),
		lunoService.ProvideLunoKeys(
			false,
			&lunoService.LunoKeys{
				Key:    "e52n78axhy2j7",
				Secret: "4q00paAkXche01noiISYWsZQGtSOKe1kuMnQUk3m3Io",
			}),
		lunoService.ProvideLunoAPIKeyID(),
		lunoService.ProvideLunoAPIKeySecret(),
	)
	if app.FxApp.Err() != nil {
		println(app.FxApp.Err().Error())
		return
	}
	app.RunTerminalApp()
}
