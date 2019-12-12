package gui

import (
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type CommitPanel struct {
	*tview.Table
	commits []Commit
}

func NewCommitPanel() *CommitPanel {
	p := &CommitPanel{
		Table: tview.NewTable(),
	}

	p.SetBorder(true).SetTitle("commits").SetTitleAlign(tview.AlignLeft)
	p.SetFixed(1, 0).SetOffset(1, 1).SetSelectable(true, false)
	return p
}

func (c *CommitPanel) UpdateView(commits []Commit) {
	c.commits = commits

	table := c.Clear()

	headers := []string{
		"hash",
		"author",
		"message",
	}

	for i, h := range headers {
		table.SetCell(0, i, &tview.TableCell{
			Text:            h,
			NotSelectable:   true,
			Align:           tview.AlignLeft,
			Color:           tcell.ColorYellow,
			BackgroundColor: tcell.ColorDefault,
		})
	}

	for i, co := range commits {
		table.SetCell(i+1, 0, tview.NewTableCell(fmt.Sprintf("[red::]%s", co.Hash)))
		table.SetCell(i+1, 1, tview.NewTableCell(fmt.Sprintf("[green::]%s", co.Author)))
		table.SetCell(i+1, 2, tview.NewTableCell(co.Message))
	}
}

func (c *CommitPanel) Keybinding(g *Gui) {
	c.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			row, _ := c.GetSelection()
			if len(c.commits) == 0 || row > len(c.commits) {
				return event
			}

			commit := c.commits[row-1]

			g.DiffPanel.ShowDiff(commit, g.File)
			if g.Pages.HasPage("diff") {
				g.Pages.ShowPage("diff")
			} else {
				g.Pages.AddAndSwitchToPage("diff", g.DiffPanel, true)
			}
		}
		return event
	})
}
