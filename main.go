package main

import (
	"log"
	//	"os"

	tea "github.com/charmbracelet/bubbletea"
	//	"github.com/charmbracelet/lipgloss"
	"github.com/kujtimiihoxha/vimtea"
)

func main() {
	//gameStatus := gameModel{}
	content := `Use HJKL to move to
	the exxtra lettersz and press x 
	to dellete
	themm`
	editor := vimtea.NewEditor(
		vimtea.WithContent(content),
		vimtea.WithFullScreen())
	editor.AddCommand("q", func(b vimtea.Buffer, _ []string) tea.Cmd {
		return tea.Quit
	})
	p := tea.NewProgram(editor)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type gameModel struct {
	level      int
	goal       string
	keystrokes int
}
