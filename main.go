/*
A presentation of the cview package, implemented with cview.

# Navigation

The presentation will advance to the next slide when the primitive demonstrated
in the current slide is left (usually by hitting Enter or Escape). Additionally,
the following shortcuts can be used:

  - Ctrl-N: Jump to next slide
  - Ctrl-P: Jump to previous slide
*/
package main

import (
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"strconv"

	// "code.rocketnine.space/tslocum/cview"
	"github.com/drewart/dough/data"
	"github.com/gdamore/tcell/v2"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rivo/tview"
)

const (
	appInfo      = "Next slide: Ctrl-N  Previous: Ctrl-P  Exit: Ctrl-C  (Navigate with your keyboard and mouse)"
	listInfo     = "Next item: J, Down  Previous item: K, Up  Open context menu: Alt+Enter"
	textViewInfo = "Scroll down: J, Down, PageDown  Scroll up: K, Up, PageUp"
	sliderInfo   = "Decrease: H, J, Left, Down  Increase: K, L, Right, Up"
	formInfo     = "Next field: Tab  Previous field: Shift+Tab  Select: Enter"
	windowInfo   = "Windows may be dragged and resized using the mouse."
)

// Slide is a function which returns the slide's title, any applicable
// information and its main primitive, its. It receives a "nextSlide" function
// which can be called to advance the presentation to the next slide.
type Page func(nextPage func()) (title string, info string, content tview.Primitive)

// The application.
var app = tview.NewApplication()

func InitDatabase() {
	data.InitSchema(nil)
}

func SetupLogger() {
	logFile, err := os.OpenFile("dough.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error open: %v ", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println("Logging Setup for dough")
}

// Starting point for the presentation.
func main() {
	SetupLogger()
	//defer app.

	slides := []Page{
		Accounts,
		PickFileView,
		Transactions,
		Table,
		Form,
		CatagoryReview,
	}

	pages := tview.NewPages()

	//panels := NewTabbedPanels()

	// The bottom row has some info on where we are.
	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			if len(added) == 0 {
				return
			}

			pages.SwitchToPage(added[0])
		})

	//subInfo := tview.NewTextView().SetWrap(false)

	// Create the pages for all slides.
	previousSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide - 1 + len(slides)) % len(slides)
		//panels.SetCurrentTab(strconv.Itoa(slide))
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		//panels.SetCurrentTab(strconv.Itoa(slide))
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	intSlide := func(i int) {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (i) % len(slides)
		//panels.SetCurrentTab(strconv.Itoa(slide))
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}

	for index, slide := range slides {
		title, slideInfo, primitive := slide(nextSlide)
		h := tview.NewTextView()
		if slideInfo != "" {
			h.SetDynamicColors(true)
			h.SetText("  [darkcyan]Info:[-]  " + slideInfo)
		}
		f := tview.NewFlex()
		f.SetDirection(tview.FlexRow)
		f.AddItem(h, 1, 1, false)
		f.AddItem(primitive, 0, 1, true)
		//panels.AddTab(strconv.Itoa(index), title, f)

		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		fmt.Fprintf(info, `%d ["%d"][darkcyan]%s[white][""] `, index+1, index, title)
	}
	info.Highlight("0")
	//panels.SetCurrentTab("0")

	// Create the main layout.
	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(info, 1, 1, false).
		//AddItem(panels, 1, 1, false).
		AddItem(pages, 0, 1, true)

	// Shortcuts to navigate the slides.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
			return nil
		} else if event.Key() == tcell.KeyCtrlA {
			intSlide(0)
		} else if event.Key() == tcell.KeyCtrlB {
			intSlide(1)
		}
		return event
	})

	// Start the application.
	if err := app.SetRoot(layout, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
