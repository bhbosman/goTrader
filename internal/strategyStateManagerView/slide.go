package strategyStateManagerView

import (
	"github.com/bhbosman/goTrader/internal/publish"
	"github.com/bhbosman/goTrader/internal/trackMarketView"
	ui2 "github.com/bhbosman/goUi/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"go.uber.org/fx"
)

type slide struct {
	next                   tview.Primitive
	canDraw                bool
	app                    *tview.Application
	listTable              *tview.Table
	TrackMarketViewService ITrackMarketViewService
	MainFlex               *tview.Flex
	CustomComponent        IAlgoViewer
	listPlate              *listPlate
	selectedItem           string
}

func (self *slide) Toggle(b bool) {
	self.canDraw = b
	switch b {
	case true:
		if self.selectedItem != "" {
			self.TrackMarketViewService.Subscribe(self.selectedItem)
		}
	case false:
		if self.selectedItem != "" {
			self.TrackMarketViewService.Unsubscribe(self.selectedItem)
		}
	}
	if b {
		self.app.ForceDraw()
	}
}

func (self *slide) Draw(screen tcell.Screen) {
	self.next.Draw(screen)
}

func (self *slide) GetRect() (int, int, int, int) {
	return self.next.GetRect()
}

func (self *slide) SetRect(x, y, width, height int) {
	self.next.SetRect(x, y, width, height)
}

func (self *slide) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return self.next.InputHandler()
}

func (self *slide) Focus(delegate func(p tview.Primitive)) {
	self.next.Focus(delegate)
}

func (self *slide) HasFocus() bool {
	return self.next.HasFocus()
}

func (self *slide) Blur() {
	self.next.Blur()
}

func (self *slide) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return self.next.MouseHandler()
}

func (self *slide) Close() error {
	return nil
}

func (self *slide) UpdateContent() error {
	return nil
}

func (self *slide) init() {
	self.listTable = tview.NewTable()
	self.listTable.SetBorder(true)
	self.listTable.SetSelectable(true, false)
	self.listTable.SetFixed(1, 1)
	self.listTable.SetSelectionChangedFunc(
		func(row, column int) {
			row, _ = self.listTable.GetSelection()
			if self.CustomComponent != nil {
				_ = self.CustomComponent.Close()
				self.MainFlex.RemoveItem(self.CustomComponent)
			}

			if item, ok := self.listPlate.GetItem(row); ok {
				self.CustomComponent = trackMarketView.NewAlgoViewer()
				self.MainFlex.AddItem(self.CustomComponent, 0, 3, false)
				if item != self.selectedItem {
					if self.selectedItem != "" {
						self.TrackMarketViewService.Unsubscribe(self.selectedItem)
					}
					self.selectedItem = item
					if self.canDraw {
						self.TrackMarketViewService.Subscribe(self.selectedItem)
					}
				}
			}
		},
	)
	self.listTable.SetSelectedFunc(
		func(row, column int) {

		},
	)
	//self.CustomComponent = self.table
	self.MainFlex = tview.NewFlex().
		SetDirection(tview.FlexColumn)
	self.MainFlex.AddItem(self.listTable, 30, 1, true)

	self.listTable.SetContent(&emptyCell{})
	flex := tview.NewFlex().
		AddItem(
			self.MainFlex,
			0,
			1,
			true)

	self.next = flex
}

func (self *slide) onStrategyDataChange(name string, strategy publish.IStrategy) bool {
	return self.app.QueueUpdate(
		func() {
			row, _ := self.listTable.GetSelection()
			if text, ok := self.listPlate.GetItem(row); ok {
				if text == strategy.GetStrategyName() {
					if self.CustomComponent != nil {
						self.CustomComponent.SetData(strategy)
					}
					if self.canDraw {
						self.app.ForceDraw()
					}
				}
			}
		},
	)
}

func (self *slide) onListChange(list []string) bool {
	return self.app.QueueUpdate(
		func() {
			if list != nil && len(list) > 0 {
				plateNil := self.listPlate == nil
				self.listPlate = newListPlate(list)
				self.listTable.SetContent(self.listPlate)
				if plateNil && self.listTable != nil {
					self.listTable.Select(1, 0)
				} else {
					row, column := self.listTable.GetSelection()
					self.listTable.Select(row, column)
				}
			} else {
				if self.selectedItem != "" {
					self.TrackMarketViewService.Unsubscribe(self.selectedItem)
				}
				self.selectedItem = ""
				self.listTable.SetContent(&emptyCell{})
				if self.CustomComponent != nil {
					self.MainFlex.RemoveItem(self.CustomComponent)
				}
			}
			if self.canDraw {
				self.app.ForceDraw()
			}
		},
	)
}

func newSlide(
	app *tview.Application,
	slideService ITrackMarketViewService,
) (*slide, error) {
	result := &slide{
		app:                    app,
		TrackMarketViewService: slideService,
	}
	result.init()
	slideService.SetListChange(result.onListChange)
	slideService.SetStrategyDataChange(result.onStrategyDataChange)
	return result, nil
}

type factory struct {
	Service ITrackMarketViewService
	app     *tview.Application
}

func (self *factory) OrderNumber() int {
	return 2
}

func (self *factory) Content(nextSlide func()) (string, ui2.IPrimitiveCloser, error) {
	slide, err := newSlide(
		self.app,
		self.Service,
	)
	if err != nil {
		return "", nil, err
	}
	return self.Title(), slide, nil
}

func (self *factory) Title() string {
	return "(some view name)"
}

func NewCoverSlideFactory(
	Service ITrackMarketViewService,
	app *tview.Application,
) *factory {
	return &factory{
		Service: Service,
		app:     app,
	}
}

func ProvideView() fx.Option {
	return fx.Options(
		fx.Provide(
			fx.Annotated{
				Group: "RegisteredMainWindowSlides",
				Target: func(
					params struct {
						fx.In
						Service ITrackMarketViewService
						App     *tview.Application
					},
				) (ui2.ISlideFactory, error) {
					return NewCoverSlideFactory(
						params.Service,
						params.App,
					), nil
				},
			},
		),
	)
}
