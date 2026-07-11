package game

import (
	"slices"

	"github.com/BosBJJ/hjkl-hero/internal/levels"
)

func CmdRepeater(gs *GameState, count int, cmd func(*GameState)) {
	if count == 0 {
		count = 1
	}
	for range count {
		cmd(gs)
	}
}

// H and L aren't wrong or bugged, for some reason this is how actual VIM accepts these deletes based on position
// J and K also aren't bugged.. VIM doesn't seem to like trying to delete current + next if there isnt a next
func (gs *GameState) DeleteDirection(input string) {
	if gs.MapInfo.MapType != EditorMap {
		return
	}
	mapLines := ToLines(*gs)
	if gs.Player.Line < 0 || gs.Player.Line >= len(mapLines) {
		return
	}
	runes := []rune(mapLines[gs.Player.Line])
	if gs.Player.Column < 0 {
		gs.Player.Column = 0
	}
	if gs.Player.Column >= len(runes) {
		gs.Player.Column = len(runes) - 1
	}
	inputRune := []rune(input)[0]
	switch inputRune {
	case 'h':
		if gs.Player.Column == 0 {
			return
		}
		gs.TakeSnapShot(gs.Player, mapLines)
		if gs.Player.Column == 1 {
			runes = append(runes[:gs.Player.Column], runes[gs.Player.Column+1:]...)
		} else {
			runes = append(runes[:gs.Player.Column-1], runes[gs.Player.Column:]...)
		}
	case 'l':
		gs.TakeSnapShot(gs.Player, mapLines)
		runes = append(runes[:gs.Player.Column], runes[gs.Player.Column+1:]...)
	case 'd':
		gs.TakeSnapShot(gs.Player, mapLines)
		mapLines = slices.Delete(mapLines, gs.Player.Line, gs.Player.Line+1)
		gs.Player.AdjustPlayer(mapLines)
		runes = []rune(mapLines[gs.Player.Line])
	case 'j':
		gs.TakeSnapShot(gs.Player, mapLines)
		remainingLines := len(mapLines) - gs.Player.Line
		if remainingLines >= 2 {
			mapLines = slices.Delete(mapLines, gs.Player.Line, gs.Player.Line+2)
			gs.Player.AdjustPlayer(mapLines)
			runes = []rune(mapLines[gs.Player.Line])
		}
	case 'k':
		gs.TakeSnapShot(gs.Player, mapLines)
		if gs.Player.Line >= 2 {
			mapLines = slices.Delete(mapLines, gs.Player.Line-1, gs.Player.Line+1)
			gs.Player.AdjustPlayer(mapLines)
			runes = []rune(mapLines[gs.Player.Line])
		}
	}
	if len(mapLines) == 1 {
		mapLines = []string{
			"  ",
			"  ",
		}
		gs.Player.Line = 1
	}
	if len(runes) == 0 {
		runes = []rune{
			' ',
			' ',
		}
		gs.Player.Column = 1
	}
	gs.Player.AdjustPlayer(mapLines)
	mapLines[gs.Player.Line] = string(runes)
	changedLine := ToText(mapLines)
	gs.MapInfo.LevelMap = levels.LevelMap(changedLine)
}

func (gs *GameState) DeleteAt() {
	mapLines := ToLines(*gs)
	if gs.Player.Line < 0 || gs.Player.Line >= len(mapLines) {
		return
	}
	runes := []rune(mapLines[gs.Player.Line])
	if gs.Player.Column < 0 {
		gs.Player.Column = 0
	}
	if gs.Player.Column >= len(runes) {
		gs.Player.Column = len(runes) - 1
	}
	gs.TakeSnapShot(gs.Player, mapLines)
	lastIndex := len(runes) - 1
	if runes[0] == '#' && runes[lastIndex] == '#' {
		runes[gs.Player.Column] = '.'
	} else {
		runes = append(runes[:gs.Player.Column], runes[gs.Player.Column+1:]...)
	}
	mapLines[gs.Player.Line] = string(runes)
	if len(mapLines[gs.Player.Line]) == 1 {
		mapLines[gs.Player.Line] = "  "
	}
	changedLine := ToText(mapLines)
	gs.Player.AdjustPlayer(mapLines)
	gs.MapInfo.LevelMap = levels.LevelMap(changedLine)
}

func (gs *GameState) ReplaceAt(input string) {
	mapLines := ToLines(*gs)
	if gs.Player.Line < 0 || gs.Player.Line >= len(mapLines) {
		return
	}
	runes := []rune(mapLines[gs.Player.Line])
	if gs.Player.Line < 0 || gs.Player.Column >= len(runes) {
		return
	}
	gs.TakeSnapShot(gs.Player, mapLines)
	inputRune := []rune(input)[0]
	runes[gs.Player.Column] = inputRune
	mapLines[gs.Player.Line] = string(runes)
	changedLine := ToText(mapLines)
	gs.MapInfo.LevelMap = levels.LevelMap(changedLine)
}

func (gs *GameState) Undo() {
	if len(gs.undoSnap) == 0 {
		return
	}
	lastSnap := gs.undoSnap[len(gs.undoSnap)-1]
	gs.redoSnap = append(gs.redoSnap, gs.CurrentSnapShot())
	previousMap := ToText(lastSnap.MapSnapShot)
	gs.MapInfo.LevelMap = levels.LevelMap(previousMap)
	gs.Player = lastSnap.PlayerSnapShot
	gs.undoSnap = gs.undoSnap[:len(gs.undoSnap)-1]
}

func (gs *GameState) Redo() {
	if len(gs.redoSnap) == 0 {
		return
	}
	lastSnap := gs.redoSnap[len(gs.redoSnap)-1]
	gs.undoSnap = append(gs.undoSnap, gs.CurrentSnapShot())
	previousMap := ToText(lastSnap.MapSnapShot)
	gs.MapInfo.LevelMap = levels.LevelMap(previousMap)
	gs.Player = lastSnap.PlayerSnapShot
	gs.redoSnap = gs.redoSnap[:len(gs.redoSnap)-1]

}

func (gs *GameState) MapComplete() bool {
	return gs.MapInfo.LevelMap == gs.MapInfo.AnswerMap
}
