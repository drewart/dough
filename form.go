package main

import (
	"log"

	//"code.rocketnine.space/tslocum/tview"
	"github.com/drewart/dough/data"
	"github.com/rivo/tview"
)

const form = `[green]Accounts:
[white] Checking| Savings
`

var accountForm *tview.Form

// Form demonstrates forms.
func Form(nextSlide func()) (title string, info string, content tview.Primitive) {

	accountForm = tview.NewForm()
	accountForm.AddInputField("Account:", "", 30, nil, nil)
	accountForm.AddDropDown("Type:", []string{"Checking", "Savings", "Credit", "Loan"}, 0, nil)
	//accountForm.AddDropDownSimple("Type:", 0, nil, )
	accountForm.AddCheckbox("Budget:", true, nil)

	modal := tview.NewModal()
	modal.SetText("Saved!")
	modal.AddButtons([]string{"Add Another?", "Done"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLable string) {
		if buttonIndex == 0 {
			accountForm.Clear(true)
		} else if buttonIndex == 1 {
			nextSlide()
		}
	})
	//modal.Set

	if modal != nil {
		log.Println("model not nil")
	}

	onSave := func() {
		itemsCount := accountForm.GetFormItemCount()
		var accountName string

		for i := 0; i < itemsCount; i++ {
			formItem := accountForm.GetFormItem(i)
			if formItem.GetLabel() == "Account:" {
				inputField := formItem.(*tview.InputField)
				if inputField.GetLabel() == "Account:" {
					accountName = inputField.GetText()
					log.Printf("Account name saved %s", accountName)
				}
			} else if formItem.GetLabel() == "Type:" {
				op := formItem.(*tview.DropDown)
				i, opt := op.GetCurrentOption()
				log.Printf("%d %s", i, opt)
			} else if formItem.GetLabel() == "Budget:" {
				b := formItem.(*tview.Checkbox)
				log.Printf("is budget %s", b.IsChecked())
			}
		}
		id := 0
		storage := data.NewDoughStorage()
		accounts := storage.GetAccounts()
		for _, account := range accounts {
			id = account.ID
			id = id + 1
		}
		storage.InsertAccount(id, accountName)
		//modal.SetVisible(true)
	}

	accountForm.AddButton("Save", onSave)
	accountForm.AddButton("Cancel", nextSlide)
	accountForm.SetBorder(true)
	accountForm.SetTitle("Account")
	return "Form", formInfo, Code(accountForm, 36, 15, form)
}
