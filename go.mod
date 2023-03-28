module github.com/bhbosman/goTrader

go 1.18

require (
	github.com/bhbosman/goCommonMarketData v0.0.0-20220802122727-698b9feba01e
	github.com/bhbosman/goCommsDefinitions v0.0.0-20230320101758-971a50fdbf8c
	github.com/bhbosman/goCommsNetDialer v0.0.0-20220726130315-bec9f09e45e7
	github.com/bhbosman/goCommsStacks v0.0.0-20220802130535-c36f51efcb47
	github.com/bhbosman/goFxApp v0.0.0-20220715185456-22d132c8b983
	github.com/bhbosman/goFxAppManager v0.0.0-20230320053418-125763bdf8d5
	github.com/bhbosman/goMessages v0.0.0-20230302063433-258339efe599
	github.com/bhbosman/goUi v0.0.0-20230327111003-b80e10f7c669
	github.com/bhbosman/gocommon v0.0.0-20230328140834-3b6f0ea2faaa
	github.com/bhbosman/gocomms v0.0.0-20230307212550-0918a992672c
	github.com/cskr/pubsub v1.0.2
	github.com/deepmap/oapi-codegen v1.11.0
	github.com/gdamore/tcell/v2 v2.5.1
	github.com/openlyinc/pointy v1.1.2
	github.com/reactivex/rxgo/v2 v2.5.0
	github.com/rivo/tview v0.0.0-20220709181631-73bf2902b59a
	go.uber.org/fx v1.19.2
	go.uber.org/multierr v1.10.0
	go.uber.org/zap v1.24.0
	golang.org/x/net v0.0.0-20220513224357-95641704303c
)

require github.com/bhbosman/goConn v0.0.0-20230327111455-7a39299fb0aa // indirect

require (
	github.com/bhbosman/goConnectionManager v0.0.0-20230327104600-4b1464715bb4 // indirect
	github.com/bhbosman/goerrors v0.0.0-20220623084908-4d7bbcd178cf // indirect
	github.com/bhbosman/gomessageblock v0.0.0-20230308173223-e8144f25444c // indirect
	github.com/bhbosman/goprotoextra v0.0.2 // indirect
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/icza/gox v0.0.0-20220321141217-e2d488ab2fbc // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/teivah/onecontext v0.0.0-20200513185103-40f981bfd775 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gdamore/tcell/v2 => github.com/bhbosman/tcell/v2 v2.5.2-0.20220624055704-f9a9454fab5b

replace github.com/golang/mock => github.com/bhbosman/gomock v1.6.1-0.20230302060806-d02c40b7514e

replace github.com/rivo/tview => ../tview

//replace github.com/rivo/tview => ../tview latest

replace github.com/bhbosman/gocomms => ../gocomms

replace github.com/bhbosman/goFxAppManager => ../goFxAppManager

replace github.com/bhbosman/gocommon => ../gocommon

replace github.com/bhbosman/goCommsStacks => ../goCommsStacks

replace github.com/bhbosman/goCommsNetDialer => ../goCommsNetDialer

replace github.com/bhbosman/goCommsNetListener => ../goCommsNetListener

replace github.com/bhbosman/goCommsDefinitions => ../goCommsDefinitions

replace github.com/bhbosman/goFxApp => ../goFxApp

replace github.com/bhbosman/goUi => ../goUi

replace github.com/bhbosman/goerrors => ../goerrors

replace github.com/bhbosman/goConnectionManager => ../goConnectionManager

//replace github.com/bhbosman/goprotoextra => ../goprotoextra

replace github.com/bhbosman/goMessages => ../goMessages

replace github.com/bhbosman/goCommonMarketData => ../goCommonMarketData

replace github.com/cskr/pubsub => github.com/bhbosman/pubsub v1.0.3-0.20220802200819-029949e8a8af
