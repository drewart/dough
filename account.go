package main

import (
	//"code.rocketnine.space/tslocum/cview"
	"strings"
	"github.com/rivo/tview"

)

const accountData = `ID|Account|Type
0|Us Bank Chccking|Checking
1|Us Bank Savings|Saving
2|Us Bank Credit Card|Credit Card
3|Us Bank HELOC|Loan`

func Accounts(nextSlide func()) (title, info string, content tview.Primitive) {
	table := tview.NewTable()
	table.SetFixed(1, 1)
	table.SetBorder(true)
	table.SetTitle("Accounts")
	for row, line := range strings.Split(accountData, "\n") {
		for column, cell := range strings.Split(line, "|") {
			color := tview.Styles.PrimaryTextColor
			if row == 0 {
				color = tview.Styles.SecondaryTextColor
			} else if column == 0 {
				color = tview.Styles.TertiaryTextColor
			}
			align := tview.AlignLeft
			if row == 0 {
				align = tview.AlignLeft
			} else if column == 0 || column >= 4 {
				align = tview.AlignRight
			}
			tableCell := tview.NewTableCell(cell)
			tableCell.SetTextColor(color)
			tableCell.SetAlign(align)
			tableCell.SetSelectable(row != 0 && column != 0)
			//if column >= 1 && column <= 3 {
			//	tableCell.SetExpansion(1)
			//}
			table.SetCell(row, column, tableCell)
		}
	}

	flex := tview.NewFlex()
	flex.SetDirection(tview.FlexRow)
	table.SetSelectable(true, false)
	table.SetSeparator(' ')
	flex.AddItem(table, 0, 1, true)

	return "Accounts", "", flex
	// data pull data
}
