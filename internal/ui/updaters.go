package ui

import (
	"github.com/BosBJJ/hjkl-hero/internal/game"
	tea "github.com/charmbracelet/bubbletea"
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
			m.EnemyMsg = m.gameState.ChasePlayer()
			if m.gameState.GetTile(m.gameState.Player.Line, m.gameState.Player.Column) == '^' {
				m.GameMessage = "You have reached the stairs! Press SPACE to go to next floor!"
			}
		case "x":
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				if gs.MapInfo.MapType == game.EditorMap {
					m.gameState.DeleteAt()
				} else {
					combatLog := gs.MeleeAttack()
					cmbMsg := combatLog.ParseLog()
					m.GameMessage = cmbMsg
					m.EnemyMsg = gs.ChasePlayer()
					if m.gameState.Stats.XPGained >= 10 {
						m.LevelPending = "Press r to level up! h- health, d- damage, c- crit chance, m- crit multiplier"
					}
				}
			})
			m.CmdCount = 0
		case "r":
			m.PendingCmd = true
			m.EditorMode = ReplaceMode
		case "d":
			m.PendingCmd = true
			m.EditorMode = DeleteMode
		case "u":
			if m.gameState.MapInfo.MapType == game.EditorMap {
				game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
					m.gameState.Undo()
				})
			}
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
		case " ":
			if m.gameState.MapInfo.MapType == game.RoomMap && m.gameState.GetTile(m.gameState.Player.Line, m.gameState.Player.Column) == '^' {
				m.LevelUp()
			}
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
			if m.gameState.MapInfo.MapType == game.EditorMap {
				m.gameState.ReplaceAt(key)
			}
			if m.gameState.MapInfo.MapType == game.RoomMap {
				m.gameState.LevelStats(key)
				if m.gameState.Stats.XPGained < 10 {
					m.LevelPending = ""
				}
			}
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
				if m.gameState.MapInfo.MapType == game.EditorMap {
					m.gameState.DeleteDirection(key)
				} else {
					combatLog := gs.RangedAttack(key)
					cmbMsg := combatLog.ParseLog()
					m.GameMessage = cmbMsg
					m.EnemyMsg = gs.ChasePlayer()
					if m.gameState.Stats.XPGained >= 10 {
						m.LevelPending = "Press r to level up! h- health, d- damage, c- crit chance, m- crit multiplier"
					}
				}
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
				m.LevelUp()
				return m, nil
			} else {
				m.GameMessage = `Mistakes still found, keep trying and use ":w" to check status`
				m.CmdText = ""
				m.EditorMode = NormalMode
			}
		case "GoUpALevel": //REMOVE LATER JUST FOR TESTING
			m.LevelUp()
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
