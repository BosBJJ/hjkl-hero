package ui

import (
	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/levels"
)

func (m *GameModel) LevelUp() {
	m.CmdText = ""
	m.GameMessage = ""
	m.EditorMode = NormalMode
	nextLevel := m.gameState.MapInfo.Level + 1
	m.gameState.MapInfo = GetMapInfo(nextLevel)
	m.gameState.Enemies = nil
	m.gameState.Player = m.gameState.SpawnPlayer()
}

func GetMapInfo(level int) game.MapInfo {
	info := game.MapInfo{}
	info.Level = level
	sMap, ok := levels.GetLevel(info.Level)
	if !ok {
		sMap = "No map available at this level"
	}
	info.LevelMap = sMap
	info.MapType = game.GetType(sMap)
	if info.MapType == game.EditorMap {
		aMap := levels.GetAnswer(info.Level)
		info.AnswerMap = aMap
	}
	return info
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
