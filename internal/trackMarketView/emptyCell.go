package trackMarketView

import "github.com/rivo/tview"

type emptyCell struct {
	tview.TableContentReadOnly
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
