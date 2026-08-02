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
	MessageLog    []string
	LevelMsg      string
	GameOver      bool
	TotalMoves    int
	GameType      storage.GameMode
	SelectedTheme storage.ThemeID
	HelpMenu      bool
	DebugMenu     bool
}

type RunStats struct {
	PlayerLevel int
	Kills       int
	TotalXp     int
	TotalMoves  int
	MapLevel    int
	DamageTaken int
}

func (m GameModel) ViewGame() string {
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
		editorStyle.Render("   "),
		healthInfo,
	))
	leftBar := m.displayLeftPanel()
	center := lipgloss.NewStyle().Width(m.camera.Width).Height(m.camera.Height).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Center, currentMap, bottomBar))
	rightBar := m.displayRightPanel()
	if m.height <= 27 {
		leftBar = ""
		rightBar = ""
		center = "Please increase your terminal height"
	}
	if m.width <= 156 {
		leftBar = ""
		rightBar = ""
		center = "Please increase your terminal width"
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, leftBar, center, rightBar)

}
