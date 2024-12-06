package main

import (
	//"code.rocketnine.space/tslocum/cview"
	"github.com/rivo/tview"
)

// Center returns a new primitive which shows the provided primitive in its
// center, given the provided primitive's size.
func Center(width, height int, p tview.Primitive) tview.Primitive {
	subFlex := tview.NewFlex()
	subFlex.SetDirection(tview.FlexRow)
	subFlex.AddItem(tview.NewBox(), 0, 1, false)
	subFlex.AddItem(p, height, 1, true)
	subFlex.AddItem(tview.NewBox(), 0, 1, false)

	flex := tview.NewFlex()
	flex.AddItem(tview.NewBox(), 0, 1, false)
	flex.AddItem(subFlex, width, 1, true)
	flex.AddItem(tview.NewBox(), 0, 1, false)

	return flex
}
