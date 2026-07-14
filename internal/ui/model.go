package ui

import (
	"fmt"
	"time"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/render"
	tea "github.com/charmbracelet/bubbletea"
)

func NewGameModel() GameModel {
	stats := game.PlayerInfo{
		MaxHealth:     12,
		CurrentHealth: 12,
		BaseDmg:       4,
		CritChance:    10, //percent, start with 10%
		BaseCritMulti: 2,
		XPGained:      0,
	}
	return GameModel{
		gameState: game.GameState{
			Player:  game.Position{Line: 1, Column: 1},
			MapInfo: GetMapInfo(1),
			Stats:   stats,
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
	PendingCmd   bool
	CmdCount     int
	CmdText      string
	GameMessage  string
	EnemyMsg     string
	LevelPending string
}

type tickMsg time.Time

func doTick() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m GameModel) Init() tea.Cmd {
	return doTick()
}

func (m GameModel) View() string {
	currentMap := string(render.Render(m.gameState))
	return fmt.Sprintf(`Current Terminal Size -- Width: %v   Height: %v
	Player Position --- %v %v

%v

Game Type: %v
Enemies: %v
Editor Mode: %v 
Times using next command: %v
CommandText: %v
Game Message: %v
Current Health: %v
Max Health: %v
Combat Message: %v
%v
XP %v/10`,
		m.width, m.height, m.gameState.Player.Line, m.gameState.Player.Column, currentMap, m.gameState.MapInfo.MapType, len(m.gameState.Enemies), m.EditorMode, m.CmdCount, m.CmdText, m.GameMessage, m.gameState.Stats.CurrentHealth, m.gameState.Stats.MaxHealth, m.EnemyMsg, m.LevelPending, m.gameState.Stats.XPGained)
}
