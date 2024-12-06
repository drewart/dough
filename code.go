package main

import (
	"fmt"

	//"code.rocketnine.space/tslocum/cview"
	"github.com/rivo/tview"
)

// The width of the code window.
const codeWidth = 56

// Code returns a primitive which displays the given primitive (with the given
// size) on the left side and its source code on the right side.
func Code(p tview.Primitive, width, height int, code string) tview.Primitive {
	// Set up code view.
	codeView := tview.NewTextView()
	codeView.SetWrap(false)
	codeView.SetDynamicColors(true)
	codeView.SetBorderPadding(1,2,2,0)
	//codeView.SetPadding(1, 1, 2, 0)
	fmt.Fprint(codeView, code)

	f := tview.NewFlex()
	f.AddItem(Center(width, height, p), 0, 1, true)
	f.AddItem(codeView, codeWidth, 1, false)
	return f
}
