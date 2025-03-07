module github.com/bhbosman/goTrader

go 1.18

require (
	github.com/bhbosman/goCommonMarketData v0.0.0-20220802122727-698b9feba01e
	github.com/bhbosman/goCommsDefinitions v0.0.0-20230329100608-a6a24c060ad8
	github.com/bhbosman/goCommsNetDialer v0.0.0-20220726130315-bec9f09e45e7
	github.com/bhbosman/goCommsStacks v0.0.0-20220802130535-c36f51efcb47
	github.com/bhbosman/goFxApp v0.0.0-20220715185456-22d132c8b983
	github.com/bhbosman/goFxAppManager v0.0.0-20230328220050-a5b50e43977e
	github.com/bhbosman/goMessages v0.0.0-20230329104216-4906969c1e61
	github.com/bhbosman/goUi v0.0.0-20230329104221-220650220e7d
	github.com/bhbosman/gocommon v0.0.0-20230329101749-40db0f52d859
	github.com/bhbosman/gocomms v0.0.0-20230329110556-946ebc6ff5f4
	github.com/cskr/pubsub v1.0.2
	github.com/deepmap/oapi-codegen v1.11.0
	github.com/gdamore/tcell/v2 v2.5.1
	github.com/openlyinc/pointy v1.1.2
	github.com/reactivex/rxgo/v2 v2.5.0
	github.com/rivo/tview v0.0.0-20230621164836-6cc0565babaf
	go.uber.org/fx v1.20.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.24.0
	golang.org/x/net v0.5.0
)


replace github.com/gdamore/tcell/v2 => github.com/bhbosman/tcell/v2 v2.5.2-0.20220624055704-f9a9454fab5b

replace github.com/golang/mock => github.com/bhbosman/gomock v1.6.1-0.20230302060806-d02c40b7514e

replace github.com/cskr/pubsub => github.com/bhbosman/pubsub v1.0.3-0.20220802200819-029949e8a8af

replace github.com/rivo/tview => github.com/bhbosman/tview v0.0.0-20230310100135-f8b257a85d36

replace github.com/bhbosman/gocomms => ../gocomms

replace github.com/bhbosman/goFxAppManager => ../goFxAppManager

replace github.com/bhbosman/goCommsStacks => ../goCommsStacks

replace github.com/bhbosman/goCommsNetDialer => ../goCommsNetDialer

replace github.com/bhbosman/goCommsNetListener => ../goCommsNetListener

replace github.com/bhbosman/goCommsDefinitions => ../goCommsDefinitions

replace github.com/bhbosman/goFxApp => ../goFxApp

//replace github.com/bhbosman/goUi => ../goUi

replace github.com/bhbosman/goerrors => ../goerrors

replace github.com/bhbosman/goConnectionManager => ../goConnectionManager

replace github.com/bhbosman/goprotoextra => ../goprotoextra

replace github.com/bhbosman/goMessages => ../goMessages

replace github.com/bhbosman/goCommonMarketData => ../goCommonMarketData
