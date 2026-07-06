package game

import (
	"slices"

	"github.com/BosBJJ/hjkl-hero/internal/levels"
)

// H and L aren't wrong or bugged, for some reason this is how actual VIM accepts these deletes based on position
// J and K also aren't bugged.. VIM doesn't seem to like trying to delete current + next if there isnt a next
func DeleteDirection(gs GameState, input string) GameState {
	if gs.MapInfo.MapType != EditorMap {
		return gs
	}
	mapLines := ToLines(gs)
	if gs.Player.Line < 0 || gs.Player.Line >= len(mapLines) {
		return gs
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
			return gs
		}
		if gs.Player.Column == 1 {
			runes = append(runes[:gs.Player.Column], runes[gs.Player.Column+1:]...)
		} else {
			runes = append(runes[:gs.Player.Column-1], runes[gs.Player.Column:]...)
		}
	case 'l':
		runes = append(runes[:gs.Player.Column], runes[gs.Player.Column+1:]...)
	case 'd':
		mapLines = slices.Delete(mapLines, gs.Player.Line, gs.Player.Line+1)
		gs.Player.AdjustPlayer(mapLines)
		runes = []rune(mapLines[gs.Player.Line])
	case 'j':
		remainingLines := len(mapLines) - gs.Player.Line
		if remainingLines >= 2 {
			mapLines = slices.Delete(mapLines, gs.Player.Line, gs.Player.Line+2)
			gs.Player.AdjustPlayer(mapLines)
			runes = []rune(mapLines[gs.Player.Line])
		}
	case 'k':
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
	return gs

}

func DeleteAt(gs GameState) GameState {
	mapLines := ToLines(gs)
	if gs.Player.Line < 0 || gs.Player.Line >= len(mapLines) {
		return gs
	}
	runes := []rune(mapLines[gs.Player.Line])
	if gs.Player.Column < 0 {
		gs.Player.Column = 0
	}
	if gs.Player.Column >= len(runes) {
		gs.Player.Column = len(runes) - 1
	}
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
	return gs
}

func ReplaceAt(gs GameState, input string) GameState {
	mapLines := ToLines(gs)
	if gs.Player.Line < 0 || gs.Player.Line >= len(mapLines) {
		return gs
	}
	runes := []rune(mapLines[gs.Player.Line])
	if gs.Player.Line < 0 || gs.Player.Column >= len(runes) {
		return gs
	}
	inputRune := []rune(input)[0]
	runes[gs.Player.Column] = inputRune
	mapLines[gs.Player.Line] = string(runes)
	changedLine := ToText(mapLines)
	gs.MapInfo.LevelMap = levels.LevelMap(changedLine)
	return gs
}
