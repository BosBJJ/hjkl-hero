package main

import (
	"log"
	//	"os"

	"github.com/BosBJJ/hjkl-hero/internal/vim"
	tea "github.com/charmbracelet/bubbletea"
	//	"github.com/charmbracelet/lipgloss"
)

func main() {
	game, err := vim.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	p := tea.NewProgram(game, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
