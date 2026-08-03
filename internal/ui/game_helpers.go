package ui

import (
	"fmt"
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/levels"
	"github.com/BosBJJ/hjkl-hero/internal/storage"
	"github.com/BosBJJ/hjkl-hero/internal/style"
	"github.com/charmbracelet/lipgloss"
)

func (m *GameModel) LevelUp() {
	m.CmdText = ""
	m.MessageLog = nil
	m.EditorMode = NormalMode
	nextLevel := m.gameState.MapInfo.Level + 1
	switch m.GameType {
	case storage.TutorialMode:
		m.gameState.MapInfo = game.GetMapInfo(nextLevel)
	case storage.RogueLikeMode:
		m.gameState.MapInfo.Level++
		level := m.gameState.MapInfo.Level
		var height, width, rooms int
		height = 60 + (level-1)*2
		width = 80 + (level-1)*4
		rooms = 15 + level
		gameMap, specialEvent := levels.MakeMap(height, width, rooms)
		m.gameState.MapInfo.LevelMap = gameMap
		if specialEvent {
			m.AddMessage("Special event found on this floor!")
		}
		m.gameState.MapInfo.MapType = game.RoomMap
	}
	m.gameState.Enemies = nil
	m.gameState.Player = m.gameState.SpawnPlayer()
	m.AdjustCamera()
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
	if m.GameType == storage.TutorialMode {
		if m.gameState.MapInfo.Level == 9 {
			m.GameOver = true
		}
	}
	if m.gameState.MapInfo.Level > 15 {
		m.GameOver = true
	}
	if m.gameState.Stats.CurrentHealth <= 0 {
		m.GameOver = true
	}
}

func GetRunStats(m GameModel) RunStats {
	return RunStats{
		Kills:       m.gameState.Stats.Kills,
		TotalXp:     m.gameState.Stats.TotalXP,
		TotalMoves:  m.TotalMoves,
		MapLevel:    m.gameState.MapInfo.Level,
		DamageTaken: m.gameState.Stats.DamageTaken,
		PlayerLevel: m.gameState.Stats.PlayerLevel,
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
		DamageTaken:   0,
		PlayerLevel:   1,
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

func GetHelpMenu() string {
	sections := []string{
		"Movement",
		"  H J K L  - Move cursor",
		"  W        - Next WORD",
		"  B        - Previous WORD",
		"  E        - End of WORD",
		"  w        - Next word",
		"  b        - Previous word",
		"  e        - End of word",
		"  0        - Start of line",
		"  $        - End of line",

		"",
		"Actions",
		"  X        - Delete / Melee ATK",
		"  D        - Delete Mode / Ranged ATK",
		"  R        - Replace Mode / Level Up",
		"  P        - Drink Health Potion",

		"",
		"Counts",
		"  5j       - Move down 5 lines",
		"  3x       - Delete 3 characters",

		"",
		"Commands",
		"  :q       - End current game",
		"  :q!      - Quit immediately",
		"  :w       - Check level completion",
		"  :wq      - Complete and continue",
		"  :help    - Toggle help menu",
		"  :hideUI  - Hides side panels",
		"  :debug   - Toggle debug info"}

	return strings.Join(sections, "\n")

}
func DisplayHealth(currHP, maxHP int) string {
	healthInfo := fmt.Sprintf("Current Health: %v/%v", currHP, maxHP)
	if currHP <= 5 {
		return style.LowHealth.Render(healthInfo)
	}
	if currHP > 5 && currHP < 15 {
		return style.LowerHealth.Render(healthInfo)
	}
	return style.FullHealth.Render(healthInfo)
}
func (gs *GameModel) makePanel() lipgloss.Style {
	colors := style.MakePanelColor(gs.SelectedTheme)
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(style.GetColor(colors.WallColor)).
		BorderBackground(style.GetColor(colors.WallColor)).
		Background(style.GetColor(colors.WallColor)).
		Padding(1).
		Foreground(style.GetColor(colors.FloorColor)).
		Bold(true)
}

func (m GameModel) displayLeftPanel() string {
	panelcolors := m.makePanel()
	barSizeLeft := int(float64(m.width) * 0.18)
	stats := m.ShowStats()
	messages := m.ShowMessages()
	leftBar := panelcolors.Width(barSizeLeft).Height(m.height - 2).Render(lipgloss.JoinVertical(lipgloss.Left, stats, "Game Messages", "-----------------------------", messages))
	return leftBar
}

func (m GameModel) displayRightPanel() string {
	panelcolors := m.makePanel()
	height, width := game.GetMapSize(m.gameState)
	barSizeRight := int(float64(m.width) * 0.20)
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
	gameDebugInfo := fmt.Sprintf("\n\nPlayer Position - %v %v\nGame Type: %v\nEnemies: %v\nMoves: %v\n\n%v",
		m.gameState.Player.Line, m.gameState.Player.Column, m.gameState.MapInfo.MapType, len(m.gameState.Enemies), m.TotalMoves, termInfo)
	rightBar := panelcolors.Width(barSizeRight).Height(m.height - 2).Render(lipgloss.JoinVertical(lipgloss.Center, helpMenu))
	if m.DebugMenu {
		rightBar = panelcolors.Width(barSizeRight).Height(m.camera.Height).Render(lipgloss.JoinVertical(lipgloss.Left, helpMenu, gameDebugInfo))
	}
	return rightBar
}

func (m GameModel) displayShop() string {
	panelcolors := m.makePanel()
	barSizeRight := int(float64(m.width) * 0.20)
	options := []string{"1. Health Potion - 15 Gold", "2. 10 XP - 40 Gold", "3. Leave Shop"}
	rightBar := panelcolors.Width(barSizeRight).Height(m.height - 2).Render(lipgloss.JoinVertical(lipgloss.Left, options...))
	return rightBar
}

func (m GameModel) ShowStats() string {
	playerLevel := fmt.Sprintf("Player Level: %v\nXP: %v/10\n", m.gameState.Stats.PlayerLevel, m.gameState.Stats.XPGained)
	mapLevel := fmt.Sprintf("Floor: %v\n", m.gameState.MapInfo.Level)
	potionCount := 0
	for _, item := range m.gameState.Stats.Inventory {
		if item.Type == game.HealthPotion {
			potionCount++
		}
	}
	charStats := fmt.Sprintf("\nAttack Damage: %v\nCrit Chance: %v%%\nCrit Multi: %vx\n", m.gameState.Stats.BaseDmg, m.gameState.Stats.CritChance, m.gameState.Stats.BaseCritMulti)
	goldCount := fmt.Sprintf("Gold: %v\n", m.gameState.Stats.Gold)
	xpInfo := fmt.Sprintf("\n%v", m.LevelMsg)
	potionsAvailable := fmt.Sprintf("Potions Available: %v\n\n\n", potionCount)
	if potionCount > 0 {
		if m.gameState.Stats.CurrentHealth >= m.gameState.Stats.MaxHealth {
			potionsAvailable = fmt.Sprintf("Use P to drink potion and overheal for 3 health! Potions Available: %v\n", potionCount)
		} else {
			potionsAvailable = fmt.Sprintf("Use P to drink potion and heal for 6 health! Potions Available: %v\n", potionCount)
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, mapLevel, goldCount, playerLevel, charStats, potionsAvailable, xpInfo)

}
func (m *GameModel) AddMessage(msg string) {
	if msg == "" {
		return
	}
	m.MessageLog = append(m.MessageLog, msg+"\n")
	const maxMessages = 5
	if len(m.MessageLog) > maxMessages {
		m.MessageLog = m.MessageLog[len(m.MessageLog)-maxMessages:]
	}
}

func (m GameModel) ShowMessages() string {
	return strings.Join(m.MessageLog, "\n")
}
