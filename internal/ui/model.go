package ui

import (
	"fmt"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/levels"
	"github.com/BosBJJ/hjkl-hero/internal/render"
	tea "github.com/charmbracelet/bubbletea"
)

func NewGameModel() GameModel {
	sMap, _ := levels.GetLevel(1)

	return GameModel{
		gameState: game.GameState{
			Player: game.Position{Line: 1, Column: 1},
			Level:  1,
		},
		currentMap: sMap,
	}
}

type GameModel struct {
	gameState  game.GameState
	width      int
	height     int
	currentMap levels.LevelMap
}

func (m GameModel) Init() tea.Cmd {
	return nil
}

func tempTestDelete(sMap levels.LevelMap, gs game.GameState) levels.LevelMap {
	lines := levels.ToLines(sMap)
	newMap := levels.DeleteAt(lines, gs.Player.Line, gs.Player.Column)
	return levels.LevelMap(levels.ToText(newMap))
}

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "h", "j", "k", "l":
			m.gameState.Player.Move(msg.String(), m.width, m.height)
		case "x":
			m.currentMap = tempTestDelete(m.currentMap, m.gameState)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m GameModel) View() string {
	currentMap := string(render.Render(m.gameState, m.currentMap))
	return fmt.Sprintf("Current Terminal Size -- Width: %v   Height: %v\nPlayer Position --- %v %v\n%v", m.width, m.height, m.gameState.Player.Line, m.gameState.Player.Column, currentMap)
}
