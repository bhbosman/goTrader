package trackMarketView

import (
	"fmt"
	"github.com/bhbosman/goTrader/internal/publish"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"time"
)

type AlgoViewer struct {
	flex            *tview.Flex
	marketDataTable *tview.Table
	tableList02     *tview.Table
	label           *tview.TextView
	marketDataPlate *marketDataPlate
}

func (self *AlgoViewer) Close() error {
	return nil
}

func (self *AlgoViewer) Draw(screen tcell.Screen) {
	self.flex.Draw(screen)
}

func (self *AlgoViewer) GetRect() (int, int, int, int) {
	return self.flex.GetRect()
}

func (self *AlgoViewer) SetRect(x, y, width, height int) {
	self.flex.SetRect(x, y, width, height)
}

func (self *AlgoViewer) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return self.flex.InputHandler()
}

func (self *AlgoViewer) Focus(delegate func(p tview.Primitive)) {
	self.flex.Focus(delegate)
}

func (self *AlgoViewer) HasFocus() bool {
	return self.flex.HasFocus()
}

func (self *AlgoViewer) Blur() {
	self.flex.Blur()
}

func (self *AlgoViewer) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return self.flex.MouseHandler()
}

func (self *AlgoViewer) SetData(data interface{}) {
	if unk, ok := data.(*publish.PublishData); ok {
		self.marketDataPlate = newMarketDataPlate(unk.MarketData)
		self.marketDataTable.SetContent(self.marketDataPlate)
		self.setLabel(unk.Date, unk.State)
	}
}

func (self *AlgoViewer) init() {
	self.flex.SetDirection(tview.FlexRow)
	self.marketDataTable = tview.NewTable()
	self.marketDataTable.SetTitle("MarketData")
	self.marketDataTable.SetBorder(true)
	self.tableList02 = tview.NewTable()
	self.tableList02.SetTitle("dddd")
	self.tableList02.SetBorder(true)

	self.label = tview.NewTextView()
	self.setLabel(time.Now(), "(Unknown)")

	self.flex.AddItem(self.label, 4, 1, true)
	self.flex.AddItem(self.marketDataTable, 0, 40, false)
	self.flex.AddItem(self.tableList02, 0, 40, false)
}

func (self *AlgoViewer) setLabel(dt time.Time, state string) {
	self.label.Clear()
	_, _ = fmt.Fprintf(self.label, "Update Time: %v \nState: %v", dt, state)
}

func NewAlgoViewer() *AlgoViewer {
	flex := tview.NewFlex()

	result := &AlgoViewer{
		flex: flex,
	}
	result.init()
	return result
}
