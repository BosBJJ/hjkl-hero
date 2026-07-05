package game

import "github.com/BosBJJ/hjkl-hero/internal/levels"

// H and L aren't wrong or bugged, for some reason this is how actual VIM accepts these deletes based on position
func DeleteDirection(gs GameState, input string) MapInfo {
	if gs.MapInfo.MapType != EditorMap {
		return gs.MapInfo
	}
	mapLines := ToLines(gs)
	playerX := gs.Player.Line
	playerY := gs.Player.Column
	if playerX < 0 || playerX >= len(mapLines) {
		return gs.MapInfo
	}
	runes := []rune(mapLines[playerX])
	if playerY < 0 {
		playerY = 0
	}
	if playerY >= len(runes) {
		playerY = len(runes) - 1
	}
	inputRune := []rune(input)[0]
	switch inputRune {
	case 'h':
		if playerY == 0 {
			return gs.MapInfo
		}
		if playerY == 1 {
			runes = append(runes[:playerY], runes[playerY+1:]...)
		} else {
			runes = append(runes[:playerY-1], runes[playerY:]...)
		}
	case 'l':
		runes = append(runes[:playerY], runes[playerY+1:]...)
	}
	mapLines[playerX] = string(runes)
	if len(mapLines[playerX]) == 1 {
		mapLines[playerX] = "  "
	}
	changedLine := ToText(mapLines)
	gs.MapInfo.LevelMap = levels.LevelMap(changedLine)
	return gs.MapInfo

}

func DeleteAt(gs GameState) MapInfo {
	mapLines := ToLines(gs)
	playerX := gs.Player.Line
	playerY := gs.Player.Column

	if playerX < 0 || playerX >= len(mapLines) {
		return gs.MapInfo
	}
	runes := []rune(mapLines[playerX])
	if playerY < 0 || playerY >= len(runes) {
		return gs.MapInfo
	}
	lastIndex := len(runes) - 1
	if runes[0] == '#' && runes[lastIndex] == '#' {
		runes[playerY] = '.'
	} else {
		runes = append(runes[:playerY], runes[playerY+1:]...)
	}
	mapLines[playerX] = string(runes)
	if len(mapLines[playerX]) == 1 {
		mapLines[playerX] = "  "
	}
	changedLine := ToText(mapLines)
	gs.MapInfo.LevelMap = levels.LevelMap(changedLine)
	return gs.MapInfo
}

func ReplaceAt(gs GameState, input string) MapInfo {
	mapLines := ToLines(gs)
	playerX := gs.Player.Line
	playerY := gs.Player.Column

	if playerX < 0 || playerX >= len(mapLines) {
		return gs.MapInfo
	}
	runes := []rune(mapLines[playerX])
	if playerX < 0 || playerY >= len(runes) {
		return gs.MapInfo
	}
	inputRune := []rune(input)[0]
	runes[playerY] = inputRune
	mapLines[playerX] = string(runes)
	changedLine := ToText(mapLines)
	gs.MapInfo.LevelMap = levels.LevelMap(changedLine)
	return gs.MapInfo
}
