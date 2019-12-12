package gui

import "github.com/rivo/tview"

type Gui struct {
	App         *tview.Application
	Pages       *tview.Pages
	CommitPanel *CommitPanel
	DiffPanel   *DiffPanel
	File        string
}

func New(file string) *Gui {
	return &Gui{
		App:         tview.NewApplication(),
		Pages:       tview.NewPages(),
		CommitPanel: NewCommitPanel(),
		DiffPanel:   NewDiffPanel(),
		File:        file,
	}
}

func (g *Gui) Run() error {
	commits, err := Commits(g.File)
	if err != nil {
		return err
	}

	g.CommitPanel.UpdateView(commits)
	g.Keybinding()

	grid := tview.NewGrid().SetRows(0).
		AddItem(g.CommitPanel, 0, 0, 1, 1, 0, 0, true)

	g.Pages.AddAndSwitchToPage("main", grid, true)

	return g.App.SetRoot(g.Pages, true).Run()
}

func (g *Gui) Keybinding() {
	g.CommitPanel.Keybinding(g)
	g.DiffPanel.Keybinding(g)
}
