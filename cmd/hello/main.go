package main

import (
	"fmt"
	"time"

	"code.rocketnine.space/tslocum/cview"
)

func main() {
	app := cview.NewApplication()

	tv := cview.NewTextView()
	tv.SetBorder(true)
	tv.SetTitle("Hello, world!")
	tv.SetDynamicColors(true)
	n := time.Now()

	nUTC := n.UTC()

	t, _ := time.Parse("2006-01-02", "2024-10-12")

	unixTime := t.UTC().Unix()

	t2 := time.Unix(unixTime, 0)

	t3 := t2.UTC()

	strTime := fmt.Sprintf("\nTime: %s ,unix: %d \n[green]t2:  %s\n[white]t3: %s\n[yellow]now:%s\n[red]now UTC:%s", t, unixTime, t2, t3, n, nUTC)
	tv.SetText("Lorem ipsum dolor sit amet" + strTime)

	app.SetRoot(tv, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
