package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/levels"
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
		case m.EditorMode == YankMode:
			return m.updateYank(msg)
		case m.EditorMode == TypingMode:
			return m.updateTyping(msg)
		case m.EditorMode == ChangeMode:
			return m.updateChangeMode(msg)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		if m.IsVertical() {
			m.camera.Height = m.height - 30 - 30
			m.camera.Width = m.width
		} else {
			m.camera.Height = m.height - 4
			sides := int(float64(m.width) * 0.20)
			m.camera.Width = m.width - sides - sides
		}
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
						m.LevelMsg = "Press r to level up!\nd- damage\nc- crit chance\nm- crit multiplier\n"
					default:
						m.LevelMsg = "Press r to level up!\nh- health\nd- damage\nc- crit chance\nm- crit multiplier\n"
					}
				}
			}
			return m, nil
		}
		switch msg.String() {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			m.CmdCount = m.CmdCount*10 + int(msg.String()[0]-'0') //take first byte, remove '0' which is 48 and then it should be the normal value, make into int
			if m.CmdCount > 9999 {
				m.CmdCount = 9999
			}
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
				m.AddMessage("Press SPACE to go to next floor!")
			}
			if m.gameState.MapInfo.MapType == game.RoomMap && m.gameState.GetTile(m.gameState.Player.Line, m.gameState.Player.Column) == '+' {
				m.AddMessage("Press SPACE to unlock chest!")
			}
			if m.gameState.MapInfo.MapType == game.RoomMap && m.gameState.GetTile(m.gameState.Player.Line, m.gameState.Player.Column) == '$' {
				m.AddMessage("Press SPACE to trade merchant!")
			}
			item, hasItem := m.gameState.ItemAt(m.gameState.Player.Line, m.gameState.Player.Column)
			if m.gameState.MapInfo.MapType == game.RoomMap && hasItem {
				m.AddMessage(fmt.Sprintf(`Press "y" to collect %v`, item.Type))
			}
		case "w":
			m.gameState.JumpToNext()
		case "W":
			m.gameState.JumpToNextWord()
		case "b":
			m.gameState.JumpToPrev()
		case "z": //TEST, DELETE
			curr := m.gameState.PrintCurrWord()
			m.AddMessage(curr)
		case "B":
			m.gameState.JumpToPrevWord()
		case "e":
			m.gameState.JumpToEndOrPunct()
		case "E":
			m.gameState.JumpToEnd()
		case "0":
			if m.CmdCount > 0 {
				m.CmdCount = m.CmdCount * 10
				if m.CmdCount > 9999 {
					m.CmdCount = 9999
				}
			} else {
				m.gameState.JumpToStart()
			}
		case "$":
			m.gameState.JumpToLast()
		case "p":
			if m.GameType == storage.TutorialMode {
				m.gameState.PasteYanked(true)
			}
			if m.GameType != storage.NoHitMode && m.GameType != storage.TutorialMode {
				m.gameState.UsePotion()
			}
		case "P":
			if m.GameType == storage.TutorialMode {
				m.gameState.PasteYanked(false)
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
							m.LevelMsg = "Press r to level up!\nd- damage\nc- crit chance\nm- crit multiplier\n"
						default:
							m.LevelMsg = "Press r to level up!\nh- health\nd- damage\nc- crit chance\nm- crit multiplier\n"
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
		case "c":
			m.PendingCmd = true
			m.EditorMode = ChangeMode
		case "y":
			if m.gameState.MapInfo.MapType != game.EditorMap {
				grabbed := m.gameState.GrabItem()
				m.AddMessage(grabbed)
			}
			if m.gameState.MapInfo.MapType == game.EditorMap {
				m.PendingCmd = true
				m.EditorMode = YankMode
			}
		case "o":
			if m.gameState.MapInfo.MapType == game.EditorMap {
				m.gameState.InsertNewLine()
				m.AdjustCamera()
				m.EditorMode = TypingMode
			}
		case "a":
			if m.gameState.MapInfo.MapType == game.EditorMap {
				lines := game.ToLines(m.gameState)
				m.gameState.TakeSnapShot(m.gameState.Player, lines)
				m.gameState.Player.Column++
				m.TypeAfter = true
				m.EditorMode = TypingMode
			}
		case "i":
			if m.gameState.MapInfo.MapType == game.EditorMap {
				lines := game.ToLines(m.gameState)
				m.gameState.TakeSnapShot(m.gameState.Player, lines)
				m.EditorMode = TypingMode
			}
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
			if key == "i" && m.gameState.MapInfo.MapType == game.EditorMap {
				m.Inner = true
				return m, nil
			}
			game.CmdRepeater(&m.gameState, m.CmdCount, func(gs *game.GameState) {
				if m.gameState.MapInfo.MapType == game.EditorMap {
					m.gameState.DeleteDirection(key, m.Inner)
					m.Inner = false
					m.AdjustCamera()
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
							m.LevelMsg = "Press r to level up!\nd- damage\nc- crit chance\nm- crit multiplier\n"
						default:
							m.LevelMsg = "Press r to level up!\nh- health\nd- damage\nc- crit chance\nm- crit multiplier\n"
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
		if levelRequested, found := strings.CutPrefix(m.CmdText, "goto-"); found {
			if m.GameType != storage.TutorialMode {
				m.AddMessage("Command only allowed in tutorial mode")
				m.CmdText = ""
				m.EditorMode = NormalMode
				return m, nil
			}
			level, err := strconv.Atoi(levelRequested)
			availableLevels := levels.GetLevelsCount()
			if err != nil || level <= 0 || level > availableLevels-1 {
				m.AddMessage("Invalid tutorial level")
				m.CmdText = ""
				m.EditorMode = NormalMode
				return m, nil
			}
			m.GoToLevel(level)
		}
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

func (m GameModel) updateTyping(msg tea.Msg) (GameModel, tea.Cmd) {
	key := msg.(tea.KeyMsg)
	switch key.String() {
	case "esc":
		if m.TypeAfter {
			m.gameState.Player.Column--
		}
		m.EditorMode = NormalMode
		m.TypeAfter = false
		return m, nil
	case "enter":
		m.gameState.InsertNewLine()
		m.AdjustCamera()
	case "backspace":
		m.gameState.DeleteLeft()
	default:
		m.gameState.InsertInto(key.Runes[0])
		m.AdjustCamera()
	}
	return m, nil
}

func (m GameModel) updateYank(msg tea.Msg) (GameModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.PendingCmd {
			key := msg.String()
			if key == "esc" {
				m.EditorMode = NormalMode
				m.PendingCmd = false
				return m, nil
			}
			if key == "i" && m.gameState.MapInfo.MapType == game.EditorMap {
				m.Inner = true
				return m, nil
			}
		}
		switch msg.String() {
		case "y":
			m.gameState.YankLine()
		case "w":
			m.gameState.YankWord(m.Inner)
			m.Inner = false
		}
		m.EditorMode = NormalMode
		m.PendingCmd = false
	}
	return m, nil
}

func (m GameModel) updateChangeMode(msg tea.Msg) (GameModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		if m.PendingCmd {
			if key == "esc" {
				m.EditorMode = NormalMode
				m.PendingCmd = false
				return m, nil
			}
			if key == "i" && m.gameState.MapInfo.MapType == game.EditorMap {
				m.Inner = true
				return m, nil
			}
		}
		switch key {
		case "w":
			m.gameState.ChangeText(m.Inner)
			m.Inner = false
			m.EditorMode = TypingMode
			return m, nil
		}
		m.EditorMode = NormalMode
		m.PendingCmd = false
	}
	return m, nil
}
