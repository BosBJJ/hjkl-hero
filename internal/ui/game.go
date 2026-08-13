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
	gs.Stats.MaxHealth = 100
	gs.Stats.CurrentHealth = 100
	return GameModel{
		gameState:  gs,
		EditorMode: NormalMode,
		PendingCmd: false,
		GameType:   storage.TutorialMode,
	}
}

func MakeRogueLikeGameModel() GameModel {
	gameMap, _ := levels.MakeMap(60, 80, 15)
	info := game.MapInfo{
		Level:    1,
		LevelMap: gameMap,
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

func MakeHardCoreGameModel() GameModel {
	gameMap, _ := levels.MakeMap(60, 80, 15)
	info := game.MapInfo{
		Level:    1,
		LevelMap: gameMap,
		MapType:  game.RoomMap,
	}
	hero := makeBaseCharacter()
	hero.MaxHealth = 1
	hero.CurrentHealth = 1
	hero.HardCore = true
	gs := game.GameState{
		MapInfo: info,
		Stats:   hero,
	}
	gs.Player = gs.SpawnPlayer()
	return GameModel{
		gameState:  gs,
		EditorMode: NormalMode,
		GameType:   storage.NoHitMode,
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
	MessageLog    []string
	LevelMsg      string
	GameOver      bool
	TotalMoves    int
	GameType      storage.GameMode
	SelectedTheme storage.ThemeID
	HelpMenu      bool
	DebugMenu     bool
	HideUI        bool
	MerchantMode  bool
}

type RunStats struct {
	PlayerLevel int
	Kills       int
	TotalXp     int
	TotalMoves  int
	MapLevel    int
	DamageTaken int
}

func (m GameModel) ViewGameHorizontal() string {
	barSizeLeft := int(float64(m.width) * 0.18)
	barSizeRight := int(float64(m.width) * 0.24)
	lowBar := m.width - barSizeLeft - barSizeRight
	currentMap := render.Render(m.gameState, m.camera, m.SelectedTheme)
	editorInfo := fmt.Sprintf("Editor Mode: %v  %v\nCommandText: %v", m.EditorMode, m.CmdCount, m.CmdText)
	editorStyle := lipgloss.NewStyle().Width(30)
	healthInfo := DisplayHealth(m.gameState.Stats.CurrentHealth, m.gameState.Stats.MaxHealth)
	bottomBar := lipgloss.NewStyle().Width(lowBar).Render(lipgloss.JoinHorizontal(
		lipgloss.Left,
		editorStyle.Render(editorInfo),
		editorStyle.Render(""),
		healthInfo,
	))
	leftBar := m.displayLeftPanel()
	center := lipgloss.NewStyle().Width(m.camera.Width).Height(m.camera.Height - 2).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Center, currentMap, bottomBar))
	if m.gameState.MapInfo.MapType == game.EditorMap {
		bottomBar = lipgloss.NewStyle().Width(lowBar).Render(lipgloss.JoinHorizontal(
			lipgloss.Left,
			editorStyle.Render(editorInfo),
		))
		hint := m.displayStringHelper("Once your text above looks like the answer below use :w, make sure to always check the game messages")
		solution := levels.GetAnswer(m.gameState.MapInfo.Level)
		center = lipgloss.NewStyle().Width(m.camera.Width).Height(m.camera.Height).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Left, currentMap, bottomBar, hint, string(solution)))
	}
	rightBar := m.displayRightPanel()
	if m.MerchantMode {
		rightBar = m.displayShop()
	}
	if m.height <= 27 {
		return "Please increase your terminal height"
	}
	if m.width <= 156 {
		return "Please increase your terminal width"
	}
	if m.HideUI {
		leftBar = lipgloss.NewStyle().Width(barSizeLeft).Render("")
		rightBar = lipgloss.NewStyle().Width(barSizeRight).Render("use :hideUI to bring back panels")
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, leftBar, center, rightBar)
}

func (m GameModel) ViewGameVertical() string {
	currentMap := render.Render(m.gameState, m.camera, m.SelectedTheme)
	editorInfo := fmt.Sprintf("Editor Mode: %v  %v\nCommandText: %v", m.EditorMode, m.CmdCount, m.CmdText)
	editorStyle := lipgloss.NewStyle().Width(30)
	healthInfo := DisplayHealth(m.gameState.Stats.CurrentHealth, m.gameState.Stats.MaxHealth)
	bottomBar := lipgloss.JoinHorizontal(
		lipgloss.Left,
		editorStyle.Render(editorInfo),
		editorStyle.Render(""),
		healthInfo,
	)
	leftBar := m.displayLeftPanelVertical()
	center := lipgloss.NewStyle().Width(m.width).Height(m.camera.Height).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Center, currentMap, bottomBar))
	if m.gameState.MapInfo.MapType == game.EditorMap {
		bottomBar = lipgloss.NewStyle().Width(m.width).Render(lipgloss.JoinHorizontal(
			lipgloss.Left,
			editorStyle.Render(editorInfo),
		))
		hint := m.displayStringHelper("Once your text above looks like the answer below use :w, Make sure to always check the game messages")
		solution := levels.GetAnswer(m.gameState.MapInfo.Level)
		center = lipgloss.NewStyle().Width(m.width).Height(m.camera.Height).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Left, currentMap, bottomBar, hint, string(solution)))
	}
	rightBar := m.displayRightPanelVertical()
	if m.MerchantMode {
		rightBar = m.displayShop()
	}
	if m.height <= 78 {
		return "Please increase your terminal height"
	}
	if m.width <= 92 {
		return "Please increase your terminal width"
	}
	return lipgloss.JoinVertical(lipgloss.Center, leftBar, center, rightBar)
}

func (m GameModel) ViewGame() string {
	if m.IsVertical() {
		return m.ViewGameVertical()
	}
	return m.ViewGameHorizontal()
}
