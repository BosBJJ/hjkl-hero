package ui

import (
	"fmt"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/levels"
	"github.com/BosBJJ/hjkl-hero/internal/render"
	tea "github.com/charmbracelet/bubbletea"
)

func NewGameModel() GameModel {
	return GameModel{
		gameState: game.GameState{
			Player:  game.Position{Line: 1, Column: 1},
			MapInfo: GetMapInfo(),
		},
	}
}

func GetMapInfo() game.MapInfo {
	info := game.MapInfo{}
	info.Level = 1
	sMap, _ := levels.GetLevel(info.Level)
	info.LevelMap = sMap
	info.MapType = game.GetType(sMap)
	return info
}

type GameModel struct {
	gameState game.GameState
	width     int
	height    int
}

func (m GameModel) Init() tea.Cmd {
	return nil
}

type EditorMode int

const (
	NormalMode EditorMode = iota
	ReplaceMode
)

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "h", "j", "k", "l":
			m.gameState.Player.Move(msg.String(), m.gameState)
		case "x":
			m.gameState.MapInfo = game.DeleteAt(m.gameState)
		case "r":
			m.gameState.MapInfo = game.ReplaceAt(m.gameState, msg.String())
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m GameModel) View() string {
	currentMap := string(render.Render(m.gameState))
	return fmt.Sprintf("Current Terminal Size -- Width: %v   Height: %v\nPlayer Position --- %v %v\n%v\nGame Type: %v", m.width, m.height, m.gameState.Player.Line, m.gameState.Player.Column, currentMap, m.gameState.MapInfo.MapType)
}
