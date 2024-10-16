package main

import ( 
	"code.rocketnine.space/tslocum/cview"
	"strings"
)
const accountData = `ID|Account|Type
0|Us Bank Chccking|Checking
1|Us Bank Savings|Saving
2|Us Bank Credit Card|Credit Card
3|Us Bank HELOC|Loan`

func Accounts(nextSlide func()) (title, info string, content cview.Primitive) {
	table := cview.NewTable()
	table.SetFixed(1,1)
	table.SetBorder(true)
	table.SetTitle("Accounts")
	for row, line := range strings.Split(accountData, "\n") {
		for column, cell := range strings.Split(line, "|") {
			color := cview.Styles.PrimaryTextColor
			if row == 0 {
				color = cview.Styles.SecondaryTextColor
			} else if column == 0 {
				color = cview.Styles.TertiaryTextColor
			}
			align := cview.AlignLeft
			if row == 0 {
				align = cview.AlignCenter
			} else if column == 0 || column >= 4 {
				align = cview.AlignRight
			}
			tableCell := cview.NewTableCell(cell)
			tableCell.SetTextColor(color)
			tableCell.SetAlign(align)
			tableCell.SetSelectable(row != 0 && column != 0)
			if column >= 1 && column <= 3 {
				tableCell.SetExpansion(1)
			}
			table.SetCell(row, column, tableCell)
		}

	}

	flex := cview.NewFlex()
	flex.SetDirection(cview.FlexRow)
	flex.AddItem(table,0, 1, true)

	return "Accounts", "", flex
	// data pull data
}