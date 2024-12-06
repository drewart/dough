package main

import (
	"log"

	//"code.rocketnine.space/tslocum/cview"
	"github.com/rivo/tview"
)

/*
const form = `[green]Accounts:
[white] Checking| Savings
`
*/
/*
func GetCatOptions() []*tview.DropDown{

	var items []*tview.DropDown

	strs := []string{"A", "B", "C"}

	for _, s := range strs {
		items = append(items, tview.NewDropDownOption(s))

	}
	return items
}
*/

// Form demonstrates forms.
func Transactions(nextSlide func()) (title string, info string, content tview.Primitive) {

	onSave := func() {
		log.Println("save")
		nextSlide()
	}

	f := tview.NewForm().
		AddInputField("First name:", "", 0, nil, nil).
		AddInputField("Last name:", "", 0, nil, nil).
		AddDropDown("Role:", []string{"Engineer", "Manager", "Administration"}, 0, nil).
		AddCheckbox("On vacation:", false, nil).
		AddPasswordField("Password:", "", 10, '*', nil).
		AddTextArea("Notes:", "", 0, 2, 0, nil).
		AddButton("Save", onSave).
		AddButton("Cancel", nextSlide)
	f.SetBorder(true).SetTitle("Transaction")

	form := "[green]foo\n[white]bar\n"

	return "transactions", formInfo, Code(f, 36, 15, form)
}
