package gui

import "github.com/rivo/tview"

type Gui struct {
	App         *tview.Application
	Pages       *tview.Pages
	CommitPanel *CommitPanel
}

func New() *Gui {
	return &Gui{
		App:         tview.NewApplication(),
		Pages:       tview.NewPages(),
		CommitPanel: NewCommitPanel(),
	}
}

func (g *Gui) Run(file string) error {
	commits, err := Commits(file)
	if err != nil {
		return err
	}

	g.CommitPanel.UpdateView(commits)

	grid := tview.NewGrid().SetRows(0).
		AddItem(g.CommitPanel, 0, 0, 1, 1, 0, 0, true)

	g.Pages.AddAndSwitchToPage("main", grid, true)

	return g.App.SetRoot(g.Pages, true).Run()
}
