package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	// "terminalShop/structDefination"
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
			SetTextColor(tcell.ColorDarkGray).
			SetAlign(tview.AlignCenter).
			SetSelectable(false).
			SetExpansion(0)

		if i == 0 {
			cell.SetTextColor(tcell.ColorOrangeRed)
		}

		headerTable.SetCell(0, i, cell)
	}
	headerFrame := tview.NewFrame(headerTable).
		SetBorders(2, 2, 2, 2, 4, 4).
		AddText(fmt.Sprintf("Copyright © %d TermSite - Terminal & Website. All rights reserved.", yy), false, tview.AlignCenter, tcell.ColorGreen).
		AddText("(q) or (ctrl+c) to quit", false, tview.AlignCenter, tcell.ColorOrangeRed)

	pages.AddPage("home", headerFrame, true, true)
	pages.AddPage("post", headerFrame, true, false)
	pages.AddPage("shop", headerFrame, true, false)
	pages.AddPage("contactUs", headerFrame, true, false)
	pages.AddPage("about", headerFrame, true, false)
	pages.AddPage("faq", headerFrame, true, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		for i := range dynamicHeaderItems {
			headerTable.SetCell(0, i, tview.NewTableCell(dynamicHeaderItems[i]).
				SetTextColor(tcell.ColorDarkGray).
				SetAlign(tview.AlignCenter).
				SetSelectable(false).
				SetExpansion(0))
		}
		switch event.Rune() {
		case 'q':
			app.Stop()
			fmt.Println("Connection to terminal.site closed.")
		case 'p':
			activePage = 1
			headerTable.SetCell(0, 1, tview.NewTableCell("post (p)[*]").
				SetTextColor(tcell.ColorOrangeRed).
				SetAlign(tview.AlignCenter).
				SetSelectable(false).
				SetExpansion(0))
			pages.SwitchToPage("post")

		case 'h':
			activePage = 0
			headerTable.SetCell(0, 0, tview.NewTableCell("home (h)[*]").
				SetTextColor(tcell.ColorOrangeRed).
				SetAlign(tview.AlignCenter).
				SetSelectable(false).
				SetExpansion(0))
			pages.SwitchToPage("home")
		case 's':
			activePage = 2
			headerTable.SetCell(0, 2, tview.NewTableCell("shop (s)[*]").
				SetTextColor(tcell.ColorOrangeRed).
				SetAlign(tview.AlignCenter).
				SetSelectable(false).
				SetExpansion(0))
			pages.SwitchToPage("shop")
		case 'c':
			activePage = 3
			headerTable.SetCell(0, 3, tview.NewTableCell("contact us (c)[*]").
				SetTextColor(tcell.ColorOrangeRed).
				SetAlign(tview.AlignCenter).
				SetSelectable(false).
				SetExpansion(0))
			pages.SwitchToPage("contactUs")
		case 'a':
			activePage = 4
			headerTable.SetCell(0, 4, tview.NewTableCell("about (a)[*]").
				SetTextColor(tcell.ColorOrangeRed).
				SetAlign(tview.AlignCenter).
				SetSelectable(false).
				SetExpansion(0))
			pages.SwitchToPage("about")
		case 'f':
			activePage = 5
			headerTable.SetCell(0, 5, tview.NewTableCell("faq (f)[*]").
				SetTextColor(tcell.ColorOrangeRed).
				SetAlign(tview.AlignCenter).
				SetSelectable(false).
				SetExpansion(0))
			pages.SwitchToPage("faq")
		default:
			headerTable.SetCell(0, activePage, tview.NewTableCell(fmt.Sprintf("%s[*]", dynamicHeaderItems[activePage])).
				SetTextColor(tcell.ColorOrangeRed).
				SetAlign(tview.AlignCenter).
				SetSelectable(false).
				SetExpansion(0))
		}
		return event
	})

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
