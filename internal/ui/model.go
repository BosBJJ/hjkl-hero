package ui

import (
	"github.com/BosBJJ/hjkl-hero/internal/game"
	tea "github.com/charmbracelet/bubbletea"
)

func NewGameModel() GameModel {
	return GameModel{
		gameState: game.GameState{
			Player: game.Position{Line: 1, Column: 1},
			Level:  1,
		},
	}
}

type GameModel struct {
	gameState game.GameState
	width     int
	height    int
}

func (m GameModel) Init() tea.Cmd {
	return nil
}

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m GameModel) View() string {
	return "hjkl is loading, press ctrl+c to quit\n"
}
