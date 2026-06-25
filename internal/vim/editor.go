package vim

import (
	"fmt"

	"github.com/BosBJJ/hjkl-hero/internal/levels"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kujtimiihoxha/vimtea"
)

type Game struct {
	editor vimtea.Editor
	level  int
}

func NewGame() (*Game, error) {
	m, exists := levels.GetLevel(1)
	if !exists {
		fmt.Println("No map associated with current level yet")
		m = "No map yet"
	}
	editor := vimtea.NewEditor(
		vimtea.WithContent(string(m)),
		vimtea.WithFullScreen())
	editor.AddCommand("q", func(b vimtea.Buffer, _ []string) tea.Cmd {
		return tea.Quit
	})

	return &Game{editor: editor,}, nil
}

func (g *Game) Init() tea.Cmd {
	return nil
}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return g.editor.Update(msg)
}

func (g *Game) View() string {
	return g.editor.View()
}
