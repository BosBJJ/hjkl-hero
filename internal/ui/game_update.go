package ui

import (
	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/storage"
	tea "github.com/charmbracelet/bubbletea"
)

func (m GameModel) Update(msg tea.Msg) (GameModel, tea.Cmd) {
	switch msg := msg.(type) {
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
		m.camera.Height = m.height - 4
		sides := int(float64(m.width) * 0.20)
		m.camera.Width = m.width - sides - sides
		m.AdjustCamera()
	}
	return m, nil
}

func (m GameModel) updateNormal(msg tea.Msg) (GameModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.MerchantMode {
			switch msg.String() {
			case "1", "2", "3", "4", "5", "6", "7", "8", "9":
				merchMsg, closeShop := m.gameState.UseMerchant(int(msg.String()[0] - '0'))
				m.AddMessage(merchMsg)
				if closeShop {
					m.MerchantMode = false
				}
				if m.gameState.Stats.XPGained >= 10 {
					switch m.GameType {
					case storage.NoHitMode:
						m.LevelMsg = "Press r to level up! d- damage, c- crit chance, m- crit multiplier\n"
					default:
						m.LevelMsg = "Press r to level up! h- health, d- damage, c- crit chance, m- crit multiplier\n"
					}
				}
			}
			return m, nil
		}
		switch msg.String() {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			m.CmdCount = m.CmdCount*10 + int(msg.String()[0]-'0') //take first byte, remove '0' which is 48 and then it should be the normal value, make into int
		case "h", "j", "k", "l":
			direction := msg.String()
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				gs.Player.Move(direction, *gs)
			})
			m.AdjustCamera()
			m.TotalMoves += 1
			m.CmdCount = 0
			chase := m.gameState.ChasePlayer()
			m.AddMessage(chase)
			m.CheckGameState()
			if m.gameState.MapInfo.MapType == game.RoomMap && m.gameState.GetTile(m.gameState.Player.Line, m.gameState.Player.Column) == '^' {
				m.AddMessage("You have reached the stairs! Press SPACE to go to next floor!")
			}
			if m.gameState.MapInfo.MapType == game.RoomMap && m.gameState.GetTile(m.gameState.Player.Line, m.gameState.Player.Column) == '+' {
				m.AddMessage("You have found a chest! Press SPACE to go unlock it!")
			}
			if m.gameState.MapInfo.MapType == game.RoomMap && m.gameState.GetTile(m.gameState.Player.Line, m.gameState.Player.Column) == '$' {
				m.AddMessage("Press SPACE to trade merchant!")
			}
			grabbed := m.gameState.GrabItem()
			m.AddMessage(grabbed)
		case "w":
			m.gameState.JumpToNext()
		case "W":
			m.gameState.JumpToNextWord()
		case "b":
			m.gameState.JumpToPrev()
		case "B":
			m.gameState.JumpToPrevWord()
		case "e":
			m.gameState.JumpToEndOrPunct()
		case "E":
			m.gameState.JumpToEnd()
		case "0":
			if m.CmdCount > 0 {
				m.CmdCount = m.CmdCount * 10
			} else {
				m.gameState.JumpToStart()
			}
		case "$":
			m.gameState.JumpToLast()
		case "p":
			if m.GameType != storage.NoHitMode {
				m.gameState.UsePotion()
			}
		case "x":
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				if gs.MapInfo.MapType == game.EditorMap {
					m.gameState.DeleteAt()
				} else {
					combatLog := gs.MeleeAttack()
					cmbMsg := combatLog.ParseLog()
					m.AddMessage(cmbMsg)
					chase := gs.ChasePlayer()
					m.AddMessage(chase)
					m.CheckGameState()
					if m.gameState.Stats.XPGained >= 10 {
						switch m.GameType {
						case storage.NoHitMode:
							m.LevelMsg = "Press r to level up! d- damage, c- crit chance, m- crit multiplier\n"
						default:
							m.LevelMsg = "Press r to level up! h- health, d- damage, c- crit chance, m- crit multiplier\n"
						}
					}
				}
				m.TotalMoves += 1
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
			m.TotalMoves += 1
			m.CmdCount = 0
		case "esc": //surprisingly doesn't seem like VIM has timer by default that resets count, only goes away with button press or esc
			m.CmdCount = 0
			m.EditorMode = NormalMode
		case ":":
			m.EditorMode = CommandMode
		case " ":
			if m.gameState.MapInfo.MapType == game.RoomMap && m.gameState.GetTile(m.gameState.Player.Line, m.gameState.Player.Column) == '^' {
				m.LevelUp()
				m.CheckGameState()
				m.TotalMoves += 1
			}
			if m.gameState.MapInfo.MapType == game.RoomMap && m.gameState.GetTile(m.gameState.Player.Line, m.gameState.Player.Column) == '+' {
				obtained := m.gameState.OpenChest(m.gameState.Player.Line, m.gameState.Player.Column)
				m.AddMessage(obtained)
			}
			if m.gameState.MapInfo.MapType == game.RoomMap && m.gameState.GetTile(m.gameState.Player.Line, m.gameState.Player.Column) == '$' {
				m.MerchantMode = true
			}
		}
	}
	return m, nil
}

func (m GameModel) updateReplace(msg tea.Msg) (GameModel, tea.Cmd) {
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
				m.TotalMoves += 1
			}
			if m.gameState.MapInfo.MapType == game.RoomMap {
				m.gameState.LevelStats(key)
				if m.gameState.Stats.XPGained < 10 {
					m.LevelMsg = ""
				}
			}
			m.EditorMode = NormalMode
			m.PendingCmd = false
		}
	}
	return m, nil
}
func (m GameModel) updateDelete(msg tea.Msg) (GameModel, tea.Cmd) {
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
					m.AddMessage(cmbMsg)
					chase := gs.ChasePlayer()
					m.AddMessage(chase)
					m.CheckGameState()
					if m.gameState.Stats.XPGained >= 10 {
						switch m.GameType {
						case storage.NoHitMode:
							m.LevelMsg = "Press r to level up! d- damage, c- crit chance, m- crit multiplier\n"
						default:
							m.LevelMsg = "Press r to level up! h- health, d- damage, c- crit chance, m- crit multiplier\n"
						}
					}
				}
				m.TotalMoves += 1
			})
			m.CmdCount = 0
			m.EditorMode = NormalMode
			m.PendingCmd = false
		}
	}
	return m, nil
}

func (m GameModel) updateCommand(msg tea.Msg) (GameModel, tea.Cmd) {
	key := msg.(tea.KeyMsg)
	switch key.String() {
	case "esc":
		m.EditorMode = NormalMode
		m.CmdText = ""
		return m, nil
	case "enter":
		switch m.CmdText {
		case "q":
			m.GameOver = true
		case "q!":
			return m, tea.Quit
		case "w":
			if m.gameState.MapComplete() {
				levelComplete := `Level Completed! Please use ":wq" to close the level!`
				m.AddMessage(levelComplete)
				m.CmdText = ""
				m.EditorMode = NormalMode
			} else {
				keepTrying := "Mistakes still found, keep trying"
				m.AddMessage(keepTrying)
				m.CmdText = ""
				m.EditorMode = NormalMode
			}
		case "wq":
			if m.gameState.MapComplete() {
				m.LevelUp()
				m.CheckGameState()
				return m, nil
			} else {
				keepTrying := `Mistakes still found, keep trying and use ":w" to check status`
				m.AddMessage(keepTrying)
				m.CmdText = ""
				m.EditorMode = NormalMode
			}
		case "g": //REMOVE LATER JUST FOR TESTING
			m.LevelUp()
			m.CheckGameState()
			return m, nil
		case "cheat":
			m.gameState.Stats.XPGained += 10
			m.gameState.Stats.CurrentHealth += 10
			m.gameState.Stats.Gold += 100
		case "money":
			m.gameState.Stats.Gold += 100
		case "help":
			m.HelpMenu = !m.HelpMenu
			m.CmdText = ""
			m.EditorMode = NormalMode
		case "debug":
			m.DebugMenu = !m.DebugMenu
			m.CmdText = ""
			m.EditorMode = NormalMode
		case "hideUI":
			m.HideUI = !m.HideUI
			m.CmdText = ""
			m.EditorMode = NormalMode
		case "guide":
			if m.gameState.MapInfo.MapType == game.EditorMap {
				m.printInstructionsEditor()
			}
			if m.gameState.MapInfo.MapType == game.RoomMap {
				m.printInstructionsRogue()
			}
			m.CmdText = ""
			m.EditorMode = NormalMode
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
