package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	//"code.rocketnine.space/tslocum/tview"
	"github.com/drewart/dough/data"
	"github.com/drewart/dough/util"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ProcessLineToRow(table *tview.Table, row int, line string) {
	for column, cell := range strings.Split(line, "|") {
		color := tcell.ColorWhite
		if row == 0 {
			color = tcell.ColorYellow
		} else if column == 0 {
			color = tcell.ColorDarkCyan
		}
		align := tview.AlignLeft
		if row == 0 {
			align = tview.AlignCenter
		} else if column == 0 || column >= 4 {
			align = tview.AlignRight
		}
		tableCell := tview.NewTableCell(cell).
			SetTextColor(color).
			SetAlign(align).
			SetSelectable(row != 0 && column != 0)
		if column >= 1 && column <= 3 {
			//tableCell.SetExpansion(1)
		}
		table.SetCell(row, column, tableCell)
	}
}

var (
	cats []*data.Catagory
)

func GetCategories() {

	if cats == nil {
		ds := data.NewDoughStorage()
		cats = ds.GetCategories()
	}
}

// Table demonstrates the Table.
func CatagoryReview(nextSlide func()) (title string, info string, content tview.Primitive) {
	log.Println("in Cat Review")
	GetCategories()
	catagoryLines := []string{"[green]Catagories:\n"}
	for i, c := range cats {
		color := "blue"
		if (i % 2) == 0 {
			color = "green"
		}
		line := fmt.Sprintf("[%s] %s - %s", color, c.Code, c.Name)
		catagoryLines = append(catagoryLines, line)

	}

	file, err := os.Open("checking.csv")
	if err != nil {
		log.Fatalf("checking %s", err)
	}

	entries, err := util.ImportCSVToAccount(file)
	if err != nil {
		log.Fatalf("cat %s", err)
	}
	log.Printf("csv loaded %d entries", len(entries))

	line := "Date|Payee|Memo|Check|Trans"

	table := tview.NewTable().
		SetFixed(1, 1)
	ProcessLineToRow(table, 0, line)
	for row, entry := range entries {
		line = fmt.Sprintf("%s|%s|%s|%s|%s", entry.Date.Format("01/02/2006"), entry.Payee, entry.Memo, entry.Check, entry.TransType)
		log.Printf("processing %s\n", line)
		ProcessLineToRow(table, row+1, line)
	}

	table.SetBorder(true).SetTitle("Table")

	code := tview.NewTextView().
		SetWrap(false).
		SetDynamicColors(true)
	code.SetBorderPadding(1, 1, 2, 0)

	//list := tview.NewList()

	selectRow := func() {
		table.SetBorders(false).
			SetSelectable(true, false).
			SetSeparator(' ')
		//code.Clear()
		//fmt.Fprint(code, tableSelectRow)
	}

	tranactionForm := tview.NewForm()

	navigate := func() {
		app.SetFocus(table)
		table.SetDoneFunc(func(key tcell.Key) {
			app.SetFocus(tranactionForm)
		}).SetSelectedFunc(func(row int, column int) {
			app.SetFocus(tranactionForm)
		})
	}

	selectRow()

	//tranactionForm.AddFormItem()
	//tranactionForm.AddInputField("Data:", "", 10, nil, nil)
	//tranactionForm.AddDropDown("Cat:", 0, nil, GetCatOptions())

	tranactionForm.AddInputField("CatCode:", "", 10, nil, nil)
	tranactionForm.AddInputField("Parts:", "", 10, nil, nil)

	onSave := func() {
		itemsCount := tranactionForm.GetFormItemCount()
		log.Println(itemsCount)
		navigate()
		/*var accountName string

		id := 0
		storage := data.NewDoughStorage()
		accounts := storage.GetAccounts()
		for _, account := range accounts {
			id = account.ID
			id = id + 1
		}
		storage.InsertAccount(id, accountName)
		*/
		//tranactionForm.
		//tranactionForm.Clear(true)

		//modal.SetVisible(true)

	}

	tranactionForm.AddButton("Save", onSave)
	tranactionForm.AddButton("Cancel", nil)
	tranactionForm.SetBorder(true)
	tranactionForm.SetTitle("Transactions")

	code.SetText(strings.Join(catagoryLines, "\n"))

	return "Cats", sliderInfo, tview.NewFlex().
		AddItem(tview.NewFlex().
			SetDirection(tview.FlexRow).
			AddItem(tranactionForm, 10, 1, true).
			AddItem(table, 0, 1, false), 0, 1, false).
		AddItem(code, codeWidth, 1, false)
}
