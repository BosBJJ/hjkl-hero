package render

import (
	"github.com/BosBJJ/hjkl-hero/internal/game"
)

func Render(gs game.GameState) string {
	lines := game.ToLines(gs)
	currLine := gs.Player.Line
	currRow := gs.Player.Column
	if currLine < 0 || currLine > len(lines) {
		return string(gs.MapInfo.LevelMap)
	}
	runes := []rune(lines[currLine])
	if currRow < 0 || currRow > len(runes) {
		return string(gs.MapInfo.LevelMap)
	}
	runes[currRow] = '@'
	lines[currLine] = string(runes)
	return game.ToText(lines)
}
