package trackMarketView

import (
	"github.com/rivo/tview"
	"strconv"
)

type listPlate struct {
	tview.TableContentReadOnly
	data      []string
	emptyCell *tview.TableCell
}

func (self *listPlate) GetCell(row, column int) *tview.TableCell {
	switch row {
	case 0:
		switch column {
		case 0:
			return tview.NewTableCell("*").SetSelectable(false).SetAlign(tview.AlignRight)
		case 1:
			return tview.NewTableCell("Name").SetSelectable(false)
		}
	default:
		index := row - 1
		count := len(self.data)
		if index >= 0 && index < count {
			switch column {
			case 0:
				return tview.NewTableCell(strconv.Itoa(row)).SetSelectable(false).SetAlign(tview.AlignRight)
			case 1:
				return tview.NewTableCell(self.data[index])
			}
		}
	}
	return self.emptyCell
}

func (self *listPlate) GetRowCount() int {
	return len(self.data) + 1
}

func (self *listPlate) GetColumnCount() int {
	return 2
}

func (self *listPlate) GetItem(row int) (string, bool) {
	if row == -1 {
		return "", false
	}

	index := row - 1
	count := len(self.data)
	if index >= 0 && count > index {
		return self.data[index], true
	}
	return "", false

}

func newListPlate(data []string) *listPlate {
	return &listPlate{
		data:      data,
		emptyCell: tview.NewTableCell("").SetSelectable(false),
	}
}
