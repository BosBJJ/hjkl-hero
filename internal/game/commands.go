package game

import (
	"slices"
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/levels"
)

func CmdRepeater(gs *GameState, count int, cmd func(*GameState)) {
	if count == 0 {
		count = 1
	}
	if count > 9999 {
		count = 9999
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

// Only jumps to next word after a space
func (gs *GameState) JumpToNextWord() {
	if gs.MapInfo.MapType != EditorMap {
		return
	}
	wordPos, exists := gs.nextWORDPos()
	if !exists {
		return
	}
	gs.Player.Column = wordPos.Column
	gs.Player.Line = wordPos.Line
}

// Jumps to either next new word OR punctuation
func (gs *GameState) JumpToNext() {
	if gs.MapInfo.MapType != EditorMap {
		return
	}
	wordPos, exists := gs.nextWordPos()
	if !exists {
		return
	}
	gs.Player.Column = wordPos.Column
	gs.Player.Line = wordPos.Line
}

// Jumps to first index of line (shuffled to 1 because hardcoded maps start at index 1, normally is index 0)
func (gs *GameState) JumpToStart() {
	if gs.MapInfo.MapType != EditorMap {
		return
	}
	gs.Player.Column = 1
}

// Jumps to last index of line
func (gs *GameState) JumpToLast() {
	if gs.MapInfo.MapType != EditorMap {
		return
	}
	lines := ToLines(*gs)
	gs.Player.Column = len([]rune(lines[gs.Player.Line])) - 1
}

// Jumps to end of next word
func (gs *GameState) JumpToEnd() {
	if gs.MapInfo.MapType != EditorMap {
		return
	}
	wordPos, exists := gs.endWORDPos()
	if !exists {
		return
	}
	gs.Player.Line = wordPos.Line
	gs.Player.Column = wordPos.Column
}

// Jumps to end of word or next punct
func (gs *GameState) JumpToEndOrPunct() {
	if gs.MapInfo.MapType != EditorMap {
		return
	}
	wordPos, exists := gs.endWordPos()
	if !exists {
		return
	}
	gs.Player.Line = wordPos.Line
	gs.Player.Column = wordPos.Column
}

// Jumps to last beginning of word or punctuation
func (gs *GameState) JumpToPrev() {
	if gs.MapInfo.MapType != EditorMap {
		return
	}
	if wordPos, exists := gs.backWordPos(); exists {
		gs.Player.Column = wordPos.Column
		gs.Player.Line = wordPos.Line
	}
}

// Jump to last word
func (gs *GameState) JumpToPrevWord() {
	if gs.MapInfo.MapType != EditorMap {
		return
	}
	if wordPos, exists := gs.backWORDPos(); exists {
		gs.Player.Column = wordPos.Column
		gs.Player.Line = wordPos.Line
	}
}

func (gs *GameState) yank(text [][]rune) {
	gs.yanked = YankBuffer{
		Yanked: text,
	}
}

func (gs *GameState) YankWord() {
	pos, exists := gs.nextWordPos()
	if !exists {
		return
	}
	lines := ToLines(*gs)
	start := gs.Player
	end := pos
	runes := []rune(lines[start.Line])
	text := [][]rune{runes[start.Column:end.Column]}
	gs.yank(text)
}

func (gs *GameState) YankLine() {
	lines := ToLines(*gs)
	runes := []rune(lines[gs.Player.Line])
	gs.yank([][]rune{runes})
	gs.yanked.IsLine = true
}

func (gs *GameState) PasteYanked(after bool) {
	if gs.MapInfo.MapType != EditorMap {
		return
	}
	lines := ToLines(*gs)
	if gs.yanked.IsLine {
		gs.pasteYankedLine(lines, after)
		return
	}
	gs.pasteYankedText(lines, after)
}

func (gs *GameState) pasteYankedText(lines []string, after bool) {
	insertAt := gs.Player
	if after {
		insertAt.Column++
	}
	gs.TakeSnapShot(insertAt, lines)
	runes := []rune(lines[insertAt.Line])
	yanked := gs.yanked.Yanked[0]
	runes = append(runes[:insertAt.Column], append(yanked, runes[insertAt.Column:]...)...)
	gs.Player.Column = insertAt.Column + len(yanked) - 1
	lines[insertAt.Line] = string(runes)
	gs.MapInfo.LevelMap = levels.LevelMap(ToText(lines))
}

func (gs *GameState) pasteYankedLine(lines []string, after bool) {
	insertAt := gs.Player
	if after {
		insertAt.Line++
	}
	gs.TakeSnapShot(insertAt, lines)
	yanked := make([]string, len(gs.yanked.Yanked))
	for i, line := range gs.yanked.Yanked {
		yanked[i] = string(line)
	}
	lines = slices.Insert(lines, insertAt.Line, yanked...)
	gs.Player.Line = gs.Player.Line + len(yanked)
	gs.MapInfo.LevelMap = levels.LevelMap(ToText(lines))
}

func isSpaceOrSymbol(r rune) bool {
	return r == ' ' || isSymbol(r)
}

func isSymbol(r rune) bool {
	return r == '.' || r == ',' || r == '?' || r == '!' || r == ';'
}

// Test Func, delete later
func (gs *GameState) DisplayYank() string {
	var sb strings.Builder
	for _, line := range gs.yanked.Yanked {
		for _, rune := range line {
			sb.WriteString(string(rune))
		}
	}
	return sb.String()
}
