package render

import (
	"github.com/BosBJJ/hjkl-hero/internal/game"
)

func Render(gs game.GameState) string {
	lines := game.ToLines(gs)
	playerX := gs.Player.Line
	playerY := gs.Player.Column
	if playerX <= 0 || playerX > len(lines) {
		return string(gs.MapInfo.LevelMap)
	}
	currLine := lines[playerX]
	runes := []rune(currLine)
	if playerY <= 0 || playerY > len(runes) {
		return string(gs.MapInfo.LevelMap)
	}
	runes[playerY] = '@'
	lines[playerX] = string(runes)
	return game.ToText(lines)
}
