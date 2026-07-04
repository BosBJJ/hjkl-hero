package game

import (
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/levels"
)

type MapInfo struct {
	Level    int
	LevelMap levels.LevelMap
	LinesMap []string
	MapType  MapType
}
type MapType string

const (
	RoomMap   MapType = "room"
	EditorMap MapType = "editor"
)

func GetType(sMap levels.LevelMap) MapType {
	currMap := string(sMap)
	if strings.HasPrefix(currMap, "#") {
		return RoomMap
	}
	return EditorMap
}

func ToLines(gs GameState) []string {
	return strings.Split(string(gs.MapInfo.LevelMap), "\n")
}

func DeleteAt(gs GameState) MapInfo {
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
	lastIndex := len(runes) - 1
	if runes[0] == '#' && runes[lastIndex] == '#' {
		runes[playerY] = '.'
	} else {
		runes = append(runes[:playerY], runes[playerY+1:]...)
	}
	mapLines[playerX] = string(runes)
	changedLine := ToText(mapLines)
	gs.MapInfo.LevelMap = levels.LevelMap(changedLine)
	if playerY >= len(runes) {
		playerY = len(runes) - 1
	}
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

func ToText(lines []string) string {
	return strings.Join(lines, "\n")
}

func IsWall(gs GameState, line, col int) bool {
	lines := ToLines(gs)
	if line <= 0 || line > len(lines) {
		return true
	}
	runes := []rune(lines[line])
	if col <= 0 || col > len(runes) {
		return true
	}
	return runes[col] == '#'
}
