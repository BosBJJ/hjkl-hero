package ui

import (
	"fmt"
	"time"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/levels"
	"github.com/BosBJJ/hjkl-hero/internal/render"
	tea "github.com/charmbracelet/bubbletea"
)

func NewGameModel() GameModel {
	return GameModel{
		gameState: game.GameState{
			Player:  game.Position{Line: 1, Column: 1},
			MapInfo: GetMapInfo(1),
		},
		EditorMode: NormalMode,
		PendingCmd: false,
	}
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

type GameModel struct {
	gameState game.GameState
	width     int
	height    int
	EditorMode
	PendingCmd  bool
	CmdCount    int
	CmdText     string
	GameMessage string
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

type EditorMode int

const (
	NormalMode EditorMode = iota
	ReplaceMode
	DeleteMode
	CommandMode
)

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.gameState.MapInfo.MapType == game.RoomMap {
			m.gameState.SpawnEnemy()
		}
		return m, doTick()
	case tea.KeyMsg:
		switch {
		case msg.String() == "ctrl+c":
			return m, tea.Quit
		case m.EditorMode == ReplaceMode:
			return m.updateReplace(msg)
		case m.EditorMode == NormalMode:
			return m.updateNormal(msg)
		case m.EditorMode == DeleteMode:
			return m.updateDelete(msg)
		case m.EditorMode == CommandMode:
			return m.updateCommand(msg)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m GameModel) updateNormal(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			m.CmdCount = m.CmdCount*10 + int(msg.String()[0]-'0') //take first byte, remove '0' which is 48 and then it should be the normal value, make into int
		case "h", "j", "k", "l":
			direction := msg.String()
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				gs.Player.Move(direction, *gs)
			})
			m.CmdCount = 0
			m.GameMessage = ""
			m.gameState.ChasePlayer()
		case "x":
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				m.gameState.DeleteAt()
			})
			m.CmdCount = 0
		case "r":
			m.PendingCmd = true
			m.EditorMode = ReplaceMode
		case "d":
			m.PendingCmd = true
			m.EditorMode = DeleteMode
		case "u":
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				m.gameState.Undo()
			})
			m.CmdCount = 0
		case "ctrl+r":
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				m.gameState.Redo()
			})
			m.CmdCount = 0
		case "esc": //surprisingly doesn't seem like VIM has timer by default that resets count, only goes away with button press or esc
			m.CmdCount = 0
			m.EditorMode = NormalMode
		case ":":
			m.EditorMode = CommandMode
		}
	}
	return m, nil
}

func (m GameModel) updateReplace(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.PendingCmd {
			key := msg.String()
			if key == "esc" {
				m.EditorMode = NormalMode
				m.PendingCmd = false
				return m, nil
			}
			m.gameState.ReplaceAt(key)
			m.EditorMode = NormalMode
			m.PendingCmd = false
		}
	}
	return m, nil
}
func (m GameModel) updateDelete(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.PendingCmd {
			key := msg.String()
			if key == "esc" {
				m.EditorMode = NormalMode
				m.PendingCmd = false
				return m, nil
			}
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				m.gameState.DeleteDirection(key)
			})
			m.CmdCount = 0
			m.EditorMode = NormalMode
			m.PendingCmd = false
		}
	}
	return m, nil
}

func (m GameModel) updateCommand(msg tea.Msg) (tea.Model, tea.Cmd) {
	key := msg.(tea.KeyMsg)
	switch key.String() {
	case "esc":
		m.EditorMode = NormalMode
		m.CmdText = ""
		return m, nil
	case "enter":
		switch m.CmdText {
		case "q!":
			return m, tea.Quit
		case "w":
			if m.gameState.MapComplete() {
				m.GameMessage = `Level Completed! Please use ":wq" to close the level!`
				m.CmdText = ""
				m.EditorMode = NormalMode
			} else {
				m.GameMessage = "Mistakes still found, keep trying"
				m.CmdText = ""
				m.EditorMode = NormalMode
			}
		case "wq":
			if m.gameState.MapComplete() {
				m.CmdText = ""
				m.EditorMode = NormalMode
				nextLevel := m.gameState.MapInfo.Level + 1
				m.gameState.MapInfo = GetMapInfo(nextLevel)
				m.gameState.Player = game.Position{Line: 1, Column: 1}
				return m, nil
			} else {
				m.GameMessage = `Mistakes still found, keep trying and use ":w" to check status`
				m.CmdText = ""
				m.EditorMode = NormalMode
			}
		case "GoUpALevel": //REMOVE LATER JUST FOR TESTING
			m.CmdText = ""
			m.EditorMode = NormalMode
			nextLevel := m.gameState.MapInfo.Level + 1
			m.gameState.MapInfo = GetMapInfo(nextLevel)
			m.gameState.Player = game.Position{Line: 1, Column: 1}
			return m, nil
		default:
			m.CmdText = ""
			m.EditorMode = NormalMode
		}
	case "backspace":
		if len(m.CmdText) > 0 {
			m.CmdText = m.CmdText[:len(m.CmdText)-1]
		}
	default:
		m.CmdText += key.String()
	}
	return m, nil
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
Game Message: %v`,
		m.width, m.height, m.gameState.Player.Line, m.gameState.Player.Column, currentMap, m.gameState.MapInfo.MapType, len(m.gameState.Enemies), m.EditorMode, m.CmdCount, m.CmdText, m.GameMessage)
}
