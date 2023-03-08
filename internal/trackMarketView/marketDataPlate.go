package trackMarketView

import (
	"github.com/bhbosman/goTrader/internal/publish"
	"github.com/rivo/tview"
	"strconv"
)

type marketDataPlate struct {
	tview.TableContentReadOnly
	MarketData []*publish.MarketData
	emptyCell  *tview.TableCell
}

func newMarketDataPlate(marketData []*publish.MarketData) *marketDataPlate {
	return &marketDataPlate{
		MarketData: marketData,
		emptyCell:  tview.NewTableCell("").SetSelectable(false),
	}
}

func (self *marketDataPlate) GetCell(row, column int) *tview.TableCell {
	switch row {
	case 0:
		switch column {
		case 0:
			return tview.NewTableCell("*").SetSelectable(false).SetAlign(tview.AlignRight)
		case 1:
			return tview.NewTableCell("Instrument").SetSelectable(false)
		case 2:
			return tview.NewTableCell("BidVol").SetSelectable(false).SetAlign(tview.AlignRight)
		case 3:
			return tview.NewTableCell("Bid").SetSelectable(false).SetAlign(tview.AlignRight)
		case 4:
			return tview.NewTableCell("Ask").SetSelectable(false).SetAlign(tview.AlignRight)
		case 5:
			return tview.NewTableCell("AskVol").SetSelectable(false).SetAlign(tview.AlignRight)
		}
	default:
		row = row - 1
		index := row / 5
		res := row % 5
		m := self.MarketData[index]
		switch column {
		case 2:
			s := strconv.FormatFloat(m.Lines[res].Bid.Volume, 'f', 6, 64)
			return tview.NewTableCell(s).SetSelectable(true).SetMaxWidth(20).SetAlign(tview.AlignRight)
		case 3:
			s := strconv.FormatFloat(m.Lines[res].Bid.Price, 'f', 6, 64)
			return tview.NewTableCell(s).SetSelectable(true).SetMaxWidth(20).SetAlign(tview.AlignRight)
		case 4:
			s := strconv.FormatFloat(m.Lines[res].Ask.Price, 'f', 6, 64)
			return tview.NewTableCell(s).SetSelectable(true).SetMaxWidth(20).SetAlign(tview.AlignRight)
		case 5:
			s := strconv.FormatFloat(m.Lines[res].Ask.Volume, 'f', 6, 64)
			return tview.NewTableCell(s).SetSelectable(true).SetMaxWidth(20).SetAlign(tview.AlignRight)
		}

	}
	return self.emptyCell
}

func (self *marketDataPlate) GetRowCount() int {
	return len(self.MarketData)*5 + 1
}

func (self *marketDataPlate) GetColumnCount() int {
	return 6
}
