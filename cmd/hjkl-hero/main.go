package main

import (
	"log"

	"github.com/BosBJJ/hjkl-hero/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(ui.NewGameModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
