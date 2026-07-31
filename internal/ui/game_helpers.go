package ui

import (
	"fmt"
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/levels"
	"github.com/BosBJJ/hjkl-hero/internal/storage"
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
		var height, width, rooms int
		height = 60
		width = 80
		rooms = 15
		for range m.gameState.MapInfo.Level {
			height += 10
			width += 20
			rooms += 2
		}
		m.gameState.MapInfo.LevelMap = levels.MakeMap(height, width, rooms)
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
		"  X        - Delete / Melee Attack",
		"  D        - Delete Mode / Ranged Attack",
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
		"  :wq      - Complete level and continue",
		"  :help    - Toggle help menu",
		"  :debug   - Toggle debug info"}

	return strings.Join(sections, "\n")

}

func (m GameModel) ShowStats() string {
	mapLevel := fmt.Sprintf("Map Level: %v\n", m.gameState.MapInfo.Level)
	healthInfo := fmt.Sprintf("Current Health: %v/%v\n", m.gameState.Stats.CurrentHealth, m.gameState.Stats.MaxHealth)
	potionCount := 0
	for _, item := range m.gameState.Stats.Inventory {
		if item.Type == game.HealthPotion {
			potionCount++
		}
	}
	charStats := fmt.Sprintf("Gold: %v\n\nAttack Damage: %v\nCrit Chance: %v%%\nCrit Multi: %vx\n\n\n", m.gameState.Stats.Gold, m.gameState.Stats.BaseDmg, m.gameState.Stats.CritChance, m.gameState.Stats.BaseCritMulti)
	xpInfo := fmt.Sprintf("XP %v/10\n\n%v", m.gameState.Stats.XPGained, m.LevelMsg)
	potionsAvailable := fmt.Sprintf("Potions Available: %v\n\n\n", potionCount)
	if potionCount > 0 {
		potionsAvailable = fmt.Sprintf("Use P to drink potion and heal for 5 health! Potions Available: %v\n\n\n", potionCount)
	}
	return lipgloss.JoinVertical(lipgloss.Left, mapLevel, healthInfo, charStats, xpInfo, potionsAvailable)

}
func (m *GameModel) AddMessage(msg string) {
	if msg == "" {
		return
	}
	m.MessageLog = append(m.MessageLog, msg)
	const maxMessages = 6
	if len(m.MessageLog) > maxMessages {
		m.MessageLog = m.MessageLog[len(m.MessageLog)-maxMessages:]
	}
}

func (m GameModel) ShowMessages() string {
	return strings.Join(m.MessageLog, "\n")
}
