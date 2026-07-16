package ui

import "github.com/BosBJJ/hjkl-hero/internal/game"

func (m *GameModel) LevelUp() {
	m.CmdText = ""
	m.GameMessage = ""
	m.EditorMode = NormalMode
	nextLevel := m.gameState.MapInfo.Level + 1
	m.gameState.MapInfo = game.GetMapInfo(nextLevel)
	m.gameState.Enemies = nil
	m.gameState.Player = m.gameState.SpawnPlayer()
}

func (m *GameModel) CheckGameState() {
	if m.gameState.MapInfo.Level > 15 {
		m.GameOver = true
	}
	if m.gameState.Stats.CurrentHealth <= 0 {
		m.GameOver = true
	}
}

type EditorMode int

const (
	NormalMode EditorMode = iota
	ReplaceMode
	DeleteMode
	CommandMode
)

func (m EditorMode) String() string {
	switch m {
	case NormalMode:
		return "Normal Mode"
	case ReplaceMode:
		return "Replace Mode"
	case DeleteMode:
		return "Delete Mode"
	case CommandMode:
		return "Command Mode"
	default:
		return "InvalidMode"
	}
}
