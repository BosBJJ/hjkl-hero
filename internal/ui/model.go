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
	CmdCount   int
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
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			m.CmdCount = m.CmdCount*10 + int(msg.String()[0]-'0') //take first byte, remove '0' which is 48 and then it should be the normal value, make into int
		case "h", "j", "k", "l":
			direction := msg.String()
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				gs.Player.Move(direction, *gs)
			})
			m.CmdCount = 0
		case "x":
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				m.gameState.DeleteAt()
			})
			m.CmdCount = 0
		case "r":
			m.PendingCmd = true
			m.EditorMode = ReplaceMode
		case "d":
			m.PendingCmd = true
			m.EditorMode = DeleteMode
		case "u":
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				m.gameState.Undo()
			})
			m.CmdCount = 0
		case "ctrl+r":
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				m.gameState.Redo()
			})
			m.CmdCount = 0
		case "esc": //surprisingly doesn't seem like VIM has timer by default that resets count, only goes away with button press or esc
			m.CmdCount = 0
			m.EditorMode = NormalMode
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
			m.gameState.ReplaceAt(key)
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
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				m.gameState.DeleteDirection(key)
			})
			m.CmdCount = 0
			m.EditorMode = NormalMode
			m.PendingCmd = false
		}
	}
	return m, nil
}

func (m GameModel) View() string {
	currentMap := string(render.Render(m.gameState))
	return fmt.Sprintf("Current Terminal Size -- Width: %v   Height: %v\nPlayer Position --- %v %v\n%v\nGame Type: %v\n Editor Mode: %v, Times using next command: %v", m.width, m.height, m.gameState.Player.Line, m.gameState.Player.Column, currentMap, m.gameState.MapInfo.MapType, m.EditorMode, m.CmdCount)
}
