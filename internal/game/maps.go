package game

import (
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/levels"
)

type MapInfo struct {
	Level    int
	LevelMap levels.LevelMap
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
	currLine := ToLines(gs)
	if gs.Player.Line < 0 || gs.Player.Line >= len(currLine) {
		return gs.MapInfo
	}
	runes := []rune(currLine[gs.Player.Line])
	if gs.Player.Column < 0 || gs.Player.Column >= len(runes) {
		return gs.MapInfo
	}
	lastIndex := len(runes) - 1
	if runes[0] == '#' && runes[lastIndex] == '#' {
		runes[gs.Player.Column] = '.'
	} else {
		runes = append(runes[:gs.Player.Column], runes[gs.Player.Column+1:]...)
	}
	currLine[gs.Player.Line] = string(runes)
	changedLine := ToText(currLine)
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
