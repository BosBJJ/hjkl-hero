package ui

import (
	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/levels"
	"github.com/BosBJJ/hjkl-hero/internal/storage"
)

func (m *GameModel) LevelUp() {
	m.CmdText = ""
	m.GameMessage = ""
	m.EditorMode = NormalMode
	nextLevel := m.gameState.MapInfo.Level + 1
	switch m.GameType {
	case storage.TutorialMode:
		m.gameState.MapInfo = game.GetMapInfo(nextLevel)
	case storage.RogueLikeMode:
		m.gameState.MapInfo.Level++
		var height, width, rooms int
		height = 60
		width = 80
		rooms = 15
		for range m.gameState.MapInfo.Level {
			height += 20
			width += 20
			rooms += 1
		}
		m.gameState.MapInfo.LevelMap = levels.MakeMap(height, width, rooms)
		m.gameState.MapInfo.MapType = game.RoomMap
	}
	m.gameState.Enemies = nil
	m.gameState.Player = m.gameState.SpawnPlayer()
}

func (m *GameModel) AdjustCamera() {
	height, width := game.GetMapSize(m.gameState)
	m.camera.X = m.gameState.Player.Column - m.camera.Width/2
	m.camera.Y = m.gameState.Player.Line - m.camera.Height/2
	if m.camera.X < 0 {
		m.camera.X = 0
	}
	if m.camera.Y < 0 {
		m.camera.Y = 0
	}
	if m.camera.X > width-m.camera.Width {
		m.camera.X = width - m.camera.Width
	}
	if m.camera.Y > height-m.camera.Height {
		m.camera.Y = height - m.camera.Height
	}
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
		MaxHealth:     20,
		CurrentHealth: 20,
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
