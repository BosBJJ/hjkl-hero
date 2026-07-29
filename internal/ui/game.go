package ui

import (
	"fmt"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/levels"
	"github.com/BosBJJ/hjkl-hero/internal/render"
	"github.com/BosBJJ/hjkl-hero/internal/storage"
	"github.com/charmbracelet/lipgloss"
)

func MakeDefaultGameModel() GameModel {
	gs := game.GameState{
		MapInfo: game.GetMapInfo(1),
		Stats:   makeBaseCharacter(),
	}
	gs.Player = gs.SpawnPlayer()
	return GameModel{
		gameState:  gs,
		EditorMode: NormalMode,
		PendingCmd: false,
		GameType:   storage.TutorialMode,
	}
}

func MakeRogueLikeGameModel() GameModel {
	info := game.MapInfo{
		Level:    1,
		LevelMap: levels.MakeMap(60, 80, 15),
		MapType:  game.RoomMap,
	}
	gs := game.GameState{
		MapInfo: info,
		Stats:   makeBaseCharacter(),
	}
	gs.Player = gs.SpawnPlayer()
	return GameModel{
		gameState:  gs,
		EditorMode: NormalMode,
		GameType:   storage.RogueLikeMode,
	}
}

type GameModel struct {
	gameState game.GameState
	camera    game.Camera
	width     int
	height    int
	EditorMode
	PendingCmd    bool
	CmdCount      int
	CmdText       string
	GameMessage   string
	EnemyMsg      string
	LevelMsg      string
	GameOver      bool
	TotalMoves    int
	GameType      storage.GameMode
	SelectedTheme storage.ThemeID
	HelpMenu      bool
	DebugMenu     bool
}

type RunStats struct {
	Kills      int
	TotalXp    int
	TotalMoves int
	MapLevel   int
}

func (m GameModel) ViewGame() string {
	height, width := game.GetMapSize(m.gameState)
	helpMenu := "Type :help to display help menu"
	if m.HelpMenu == true {
		helpMenu = GetHelpMenu()
	}
	termInfo := fmt.Sprintf(
		"Terminal: %dx%d\nCamera: %dx%d\nMap: %dx%d",
		m.width, m.height,
		m.camera.Width, m.camera.Height,
		height, width,
	)
	gameDebugInfo := fmt.Sprintf("\n\nPlayer Position - %v %v\nGame Type: %v\nEnemies: %v\n\n%v",
		m.gameState.Player.Line, m.gameState.Player.Column, m.gameState.MapInfo.MapType, len(m.gameState.Enemies), termInfo)
	barSize := int(float64(m.width) * 0.20)
	currentMap := render.Render(m.gameState, m.camera, m.SelectedTheme)
	editorInfo := fmt.Sprintf("Editor Mode: %v  %v\nCommandText: %v", m.EditorMode, m.CmdCount, m.CmdText)
	characterInfo := fmt.Sprintf("Current Health: %v/%v\n\nXP %v/10\n\n%v", m.gameState.Stats.CurrentHealth, m.gameState.Stats.MaxHealth, m.gameState.Stats.XPGained, m.LevelMsg)
	combatMessages := fmt.Sprintf("Game Message: %v\n\n%v", m.GameMessage, m.EnemyMsg)
	leftBar := lipgloss.NewStyle().Width(barSize).Render(lipgloss.JoinVertical(lipgloss.Left, characterInfo, combatMessages))
	center := lipgloss.NewStyle().Width(m.camera.Width).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Left, currentMap, editorInfo))
	rightBar := lipgloss.NewStyle().Width(barSize).Render(lipgloss.JoinVertical(lipgloss.Center, helpMenu))
	if m.DebugMenu {
		rightBar = lipgloss.NewStyle().Width(barSize).Render(lipgloss.JoinVertical(lipgloss.Left, helpMenu, gameDebugInfo))
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, leftBar, center, rightBar)

}
