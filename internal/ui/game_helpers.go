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

func GetRunStats(m GameModel) RunStats {
	return RunStats{
		Kills:      m.gameState.Stats.Kills,
		TotalXp:    m.gameState.Stats.TotalXP,
		TotalMoves: m.TotalMoves,
		MapLevel:   m.gameState.MapInfo.Level,
	}
}

func makeBaseCharacter() game.PlayerInfo {
	return game.PlayerInfo{
		MaxHealth:     12,
		CurrentHealth: 12,
		BaseDmg:       4,
		CritChance:    10, //percent, start with 10%
		BaseCritMulti: 2,
		XPGained:      0,
		TotalXP:       0,
		Kills:         0,
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
