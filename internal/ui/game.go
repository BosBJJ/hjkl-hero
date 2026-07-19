package ui

import (
	"fmt"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/render"
)

func MakeDefaultGameModel() GameModel {
	return GameModel{
		gameState: game.GameState{
			Player:  game.Position{Line: 1, Column: 1},
			MapInfo: game.GetMapInfo(1),
			Stats:   makeBaseCharacter(),
		},
		EditorMode: NormalMode,
		PendingCmd: false,
	}
}

type GameModel struct {
	gameState game.GameState
	width     int
	height    int
	EditorMode
	PendingCmd  bool
	CmdCount    int
	CmdText     string
	GameMessage string
	EnemyMsg    string
	LevelMsg    string
	GameOver    bool
	TotalMoves  int
}

type RunStats struct {
	Kills      int
	TotalXp    int
	TotalMoves int
	MapLevel   int
}

func (m GameModel) ViewGame() string {
	currentMap := string(render.Render(m.gameState))
	debugInfo := fmt.Sprintf("Current Terminal Size -- Width: %v   Height: %v\nPlayer Position --- %v %v\nGame Type: %v\nEnemies: %v\n",
		m.width, m.height, m.gameState.Player.Line, m.gameState.Player.Column, m.gameState.MapInfo.MapType, len(m.gameState.Enemies))
	combatMessages := fmt.Sprintf("Combat Message: %v\n%v", m.GameMessage, m.EnemyMsg)
	characterInfo := fmt.Sprintf("Current Health: %v\nMax Health: %v\nXP %v/10\n%v\nTotalXP:%v", m.gameState.Stats.CurrentHealth, m.gameState.Stats.MaxHealth, m.gameState.Stats.XPGained, m.LevelMsg, m.gameState.Stats.TotalXP)
	editorInfo := fmt.Sprintf("Editor Mode: %v  %v\nCommandText: %v", m.EditorMode, m.CmdCount, m.CmdText)
	return fmt.Sprintf("%v\n\n%v\n\n%v\n%v\n\n%v", characterInfo, currentMap, combatMessages, editorInfo, debugInfo)
}
