package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"terminalShop/customInputHandler"
)

func main() {
	app := tview.NewApplication()
	yy, _, _ := time.Now().Date()
	pages := tview.NewPages()

	activePage := 0

	headerTable := tview.NewTable().
		SetBorders(true).
		SetSeparator('|').
		SetSelectable(false, false).
		SetEvaluateAllRows(true)

	headerItems := []string{"home (h)[*]", "post (p)", "shop (s)", "contact us (c)", "about (a)", "faq (f)"}
	dynamicHeaderItems := []string{"home (h)", "post (p)", "shop (s)", "contact us (c)", "about (a)", "faq (f)"}
	for i, item := range headerItems {
		cell := tview.NewTableCell(item).
			SetTextColor(tcell.ColorWhiteSmoke).
			SetAlign(tview.AlignCenter).
			SetSelectable(false).
			SetExpansion(1)

		if i == 0 {
			cell.SetTextColor(tcell.ColorOrangeRed)
		}

		headerTable.SetCell(0, i, cell)
	}
	headerFrame := tview.NewFrame(headerTable).
		SetBorders(2, 2, 2, 2, 4, 4).
		AddText(fmt.Sprintf("Copyright © %d TermSite - Terminal & Website. All rights reserved.", yy), false, tview.AlignCenter, tcell.ColorWhiteSmoke).
		AddText("(q) or (ctrl+c) to quit", false, tview.AlignCenter, tcell.ColorOrangeRed)

	pages.AddPage("home", headerFrame, true, true)
	pages.AddPage("post", headerFrame, true, false)
	pages.AddPage("shop", headerFrame, true, false)
	pages.AddPage("contactUs", headerFrame, true, false)
	pages.AddPage("about", headerFrame, true, false)
	pages.AddPage("faq", headerFrame, true, false)

	inputCh := make(chan *tcell.EventKey)

	go func() {
		for {
			event := <-inputCh
			app.QueueEvent(event)
		}
	}()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		inputCh <- event
		return nil
	})

	go custominputhandler.HandleInput(app, pages, headerTable, dynamicHeaderItems, &activePage, inputCh)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
