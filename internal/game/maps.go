package game

import (
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/levels"
)

type MapInfo struct {
	Level     int
	LevelMap  levels.LevelMap
	AnswerMap levels.LevelMap
	MapType   MapType
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

func GetMapSize(gs GameState) (height, width int) {
	lines := ToLines(gs)
	height = len(lines)
	for _, line := range lines {
		if len([]rune(line)) > width {
			width = len([]rune(line))
		}
	}
	return height, width
}
