module github.com/bhbosman/goTrader

go 1.23.0

toolchain go1.24.0

require (
	github.com/bhbosman/goCommonMarketData v0.0.0-20220802122727-698b9feba01e
	github.com/bhbosman/goCommsDefinitions v0.0.0-20250307125227-bfc368fdddf9
	github.com/bhbosman/goCommsNetDialer v0.0.0-20220726130315-bec9f09e45e7
	github.com/bhbosman/goCommsStacks v0.0.0-20220802130535-c36f51efcb47
	github.com/bhbosman/goFxApp v0.0.0-20220715185456-22d132c8b983
	github.com/bhbosman/goFxAppManager v0.0.0-20250307145515-bda0fa4d9959
	github.com/bhbosman/goMessages v0.0.0-20250307224348-83ddb4c19467
	github.com/bhbosman/goUi v0.0.0-20250307150712-d06325af4877
	github.com/bhbosman/gocommon v0.0.0-20250307235859-f370cb0a3bac
	github.com/bhbosman/gocomms v0.0.0-20230730212408-04ba72ddb372
	github.com/cskr/pubsub v1.0.2
	github.com/deepmap/oapi-codegen v1.11.0
	github.com/gdamore/tcell/v2 v2.8.1
	github.com/openlyinc/pointy v1.1.2
	github.com/reactivex/rxgo/v2 v2.5.0
	github.com/rivo/tview v0.0.0-20241227133733-17b7edb88c57
	go.uber.org/fx v1.23.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.27.0
	golang.org/x/net v0.37.0
)

replace github.com/gdamore/tcell/v2 => github.com/bhbosman/tcell/v2 v2.5.2-0.20220624055704-f9a9454fab5b

replace github.com/golang/mock => github.com/bhbosman/gomock v1.6.1-0.20230302060806-d02c40b7514e

replace github.com/cskr/pubsub => github.com/bhbosman/pubsub v1.0.3-0.20220802200819-029949e8a8af

replace github.com/rivo/tview => github.com/bhbosman/tview v0.0.0-20230310100135-f8b257a85d36

//replace github.com/bhbosman/gocomms => ../gocomms
//
//replace github.com/bhbosman/goFxAppManager => ../goFxAppManager
//
//replace github.com/bhbosman/goCommsStacks => ../goCommsStacks
//
//replace github.com/bhbosman/goCommsNetDialer => ../goCommsNetDialer
//
//replace github.com/bhbosman/goCommsNetListener => ../goCommsNetListener
//
//replace github.com/bhbosman/goCommsDefinitions => ../goCommsDefinitions
//
//replace github.com/bhbosman/goFxApp => ../goFxApp

//replace github.com/bhbosman/goUi => ../goUi

//replace github.com/bhbosman/goerrors => ../goerrors
//
//replace github.com/bhbosman/goConnectionManager => ../goConnectionManager
//
//replace github.com/bhbosman/goprotoextra => ../goprotoextra
//
//replace github.com/bhbosman/goMessages => ../goMessages
//
//replace github.com/bhbosman/goCommonMarketData => ../goCommonMarketData
