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
		EditorMode: NormalMode,
		PendingCmd: false,
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
	EditorMode
	PendingCmd bool
}

func (m GameModel) Init() tea.Cmd {
	return nil
}

type EditorMode int

const (
	NormalMode EditorMode = iota
	ReplaceMode
	DeleteMode
)

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case msg.String() == "ctrl+c":
			return m, tea.Quit
		case m.EditorMode == ReplaceMode:
			return m.updateReplace(msg)
		case m.EditorMode == NormalMode:
			return m.updateNormal(msg)
		case m.EditorMode == DeleteMode:
			return m.updateDelete(msg)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m GameModel) updateNormal(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h", "j", "k", "l":
			m.gameState.Player.Move(msg.String(), m.gameState)
		case "x":
			m.gameState = game.DeleteAt(m.gameState)
		case "r":
			m.PendingCmd = true
			m.EditorMode = ReplaceMode
		case "d":
			m.PendingCmd = true
			m.EditorMode = DeleteMode
		case "u":
			m.gameState.Undo()
		case "ctrl+r":
			m.gameState.Redo()
		}
	}
	return m, nil
}

func (m GameModel) updateReplace(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.PendingCmd {
			key := msg.String()
			if key == "esc" {
				m.EditorMode = NormalMode
				m.PendingCmd = false
				return m, nil
			}
			m.gameState = game.ReplaceAt(m.gameState, key)
			m.EditorMode = NormalMode
			m.PendingCmd = false
		}
	}
	return m, nil
}
func (m GameModel) updateDelete(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.PendingCmd {
			key := msg.String()
			if key == "esc" {
				m.EditorMode = NormalMode
				m.PendingCmd = false
				return m, nil
			}
			m.gameState = game.DeleteDirection(m.gameState, key)
			m.EditorMode = NormalMode
			m.PendingCmd = false
		}
	}
	return m, nil
}
func (m GameModel) View() string {
	currentMap := string(render.Render(m.gameState))
	return fmt.Sprintf("Current Terminal Size -- Width: %v   Height: %v\nPlayer Position --- %v %v\n%v\nGame Type: %v\n Editor Mode: %v", m.width, m.height, m.gameState.Player.Line, m.gameState.Player.Column, currentMap, m.gameState.MapInfo.MapType, m.EditorMode)
}
