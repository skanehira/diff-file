package gui

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

type CommitPanel struct {
	*tview.Table
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
		table.SetCell(i+1, 0, tview.NewTableCell(co.Hash))
		table.SetCell(i+1, 1, tview.NewTableCell(co.Author))
		table.SetCell(i+1, 2, tview.NewTableCell(co.Message))
	}
}
