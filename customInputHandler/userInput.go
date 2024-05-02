package custominputhandler

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
)

func HandleInput(app *tview.Application, pages *tview.Pages, headerTable *tview.Table, dynamicHeaderItems []string, activePage *int, inputCh chan *tcell.EventKey) {
	for {
		select {
		case event := <-inputCh:
			for i := range dynamicHeaderItems {
				headerTable.SetCell(0, i, tview.NewTableCell(dynamicHeaderItems[i]).
					SetTextColor(tcell.ColorDarkGray).
					SetAlign(tview.AlignCenter).
					SetSelectable(false).
					SetExpansion(1))
			}
			switch event.Key() {
			case tcell.KeyCtrlC:
				app.Stop()
				fmt.Println("Connection to terminal.site closed.")
				os.Exit(0)
			case tcell.KeyRune:
				switch event.Rune() {
				case 'q':
					app.Stop()
					fmt.Println("Connection to terminal.site closed.")
				case 'p':
					*activePage = 1
					headerTable.SetCell(0, 1, tview.NewTableCell("post (p)[*]").
						SetTextColor(tcell.ColorOrangeRed).
						SetAlign(tview.AlignCenter).
						SetSelectable(false).
						SetExpansion(1))
					pages.SwitchToPage("post")

				case 'h':
					*activePage = 0
					headerTable.SetCell(0, 0, tview.NewTableCell("home (h)[*]").
						SetTextColor(tcell.ColorOrangeRed).
						SetAlign(tview.AlignCenter).
						SetSelectable(false).
						SetExpansion(1))
					pages.SwitchToPage("home")
				case 's':
					*activePage = 2
					headerTable.SetCell(0, 2, tview.NewTableCell("shop (s)[*]").
						SetTextColor(tcell.ColorOrangeRed).
						SetAlign(tview.AlignCenter).
						SetSelectable(false).
						SetExpansion(1))
					pages.SwitchToPage("shop")
				case 'c':
					*activePage = 3
					headerTable.SetCell(0, 3, tview.NewTableCell("contact us (c)[*]").
						SetTextColor(tcell.ColorOrangeRed).
						SetAlign(tview.AlignCenter).
						SetSelectable(false).
						SetExpansion(1))
					pages.SwitchToPage("contactUs")
				case 'a':
					*activePage = 4
					headerTable.SetCell(0, 4, tview.NewTableCell("about (a)[*]").
						SetTextColor(tcell.ColorOrangeRed).
						SetAlign(tview.AlignCenter).
						SetSelectable(false).
						SetExpansion(1))
					pages.SwitchToPage("about")
				case 'f':
					*activePage = 5
					headerTable.SetCell(0, 5, tview.NewTableCell("faq (f)[*]").
						SetTextColor(tcell.ColorOrangeRed).
						SetAlign(tview.AlignCenter).
						SetSelectable(false).
						SetExpansion(1))
					pages.SwitchToPage("faq")
				default:
					headerTable.SetCell(0, *activePage, tview.NewTableCell(fmt.Sprintf("%s[*]", dynamicHeaderItems[*activePage])).
						SetTextColor(tcell.ColorOrangeRed).
						SetAlign(tview.AlignCenter).
						SetSelectable(false).
						SetExpansion(1))
				}
			}
		}
	}
}

