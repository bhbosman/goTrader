package trackMarketView

import (
	ui2 "github.com/bhbosman/goUi/ui"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"go.uber.org/fx"
)

type slide struct {
	next            tview.Primitive
	canDraw         bool
	app             *tview.Application
	table           *tview.Table
	listTable       *tview.Table
	slideService    ITrackMarketViewService
	MainFlex        *tview.Flex
	CustomComponent tview.Primitive
}

func (self *slide) Toggle(b bool) {
	self.canDraw = b
	switch b {
	case true:
		break
	case false:
		break
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
	self.table = tview.NewTable()
	self.table.SetSelectable(true, false)
	self.table.SetBorder(true)
	self.table.SetFixed(1, 1)
	self.table.SetTitle("Full Market Data Viewer")
	self.table.SetContent(&emptyCell{})

	self.listTable = tview.NewTable()
	self.listTable.SetBorder(true)
	self.listTable.SetSelectable(true, false)
	self.listTable.SetFixed(1, 1)
	self.listTable.SetSelectionChangedFunc(
		func(row, column int) {
			row, _ = self.listTable.GetSelection()
		},
	)
	self.listTable.SetSelectedFunc(
		func(row, column int) {

		},
	)
	//self.CustomComponent = self.table
	self.MainFlex = tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(self.listTable, 30, 1, true).
		AddItem(self.CustomComponent, 0, 3, false)

	self.listTable.SetContent(&emptyCell{})
	flex := tview.NewFlex().
		AddItem(
			self.MainFlex,
			0,
			1,
			true)

	self.next = flex
}

func (self *slide) onListChange(data []string) bool {
	return self.app.QueueUpdate(
		func() {
		},
	)
}

func newSlide(
	app *tview.Application,
	slideService ITrackMarketViewService,
) (*slide, error) {
	result := &slide{
		app:          app,
		slideService: slideService,
	}
	result.init()
	slideService.SetListChange(result.onListChange)
	return result, nil
}

type emptyCell struct {
}

func (self *emptyCell) GetCell(row, column int) *tview.TableCell {
	return tview.NewTableCell("")
}

func (self *emptyCell) GetRowCount() int {
	return 0
}

func (self *emptyCell) GetColumnCount() int {
	return 0
}

func (self *emptyCell) SetCell(row, column int, cell *tview.TableCell) {
}

func (self *emptyCell) RemoveRow(row int) {
}

func (self *emptyCell) RemoveColumn(column int) {
}

func (self *emptyCell) InsertRow(row int) {
}

func (self *emptyCell) InsertColumn(column int) {
}

func (self *emptyCell) Clear() {
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
