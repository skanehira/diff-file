package gui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type DiffPanel struct {
	*tview.TextView
}

func NewDiffPanel() *DiffPanel {
	d := &DiffPanel{
		TextView: tview.NewTextView().SetDynamicColors(true),
	}

	d.SetBorder(true).SetTitle("diff").SetTitleAlign(tview.AlignLeft)

	return d
}

func (d *DiffPanel) ShowDiff(commit Commit, file string) {
	contents, err := DiffContents(commit.Hash, file)
	if err != nil {
		contents = err.Error()
	}

	d.Clear().SetText(tview.TranslateANSI(contents))
}

func (d *DiffPanel) Keybinding(g *Gui) {
	d.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'q':
			g.Pages.HidePage("diff").ShowPage("main")
		}

		return event
	})
}
