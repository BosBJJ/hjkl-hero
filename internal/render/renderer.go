package render

import (
	"github.com/BosBJJ/hjkl-hero/internal/game"
)

func Render(gs game.GameState, m game.MapInfo) string {
	lines := game.ToLines(m)
	currLine := gs.Player.Line
	currRow := gs.Player.Column
	if currLine < 0 || currLine > len(lines) {
		return string(m.LevelMap)
	}
	runes := []rune(lines[currLine])
	if currRow < 0 || currRow > len(runes) {
		return string(m.LevelMap)
	}
	runes[currRow] = '@'
	lines[currLine] = string(runes)
	return game.ToText(lines)
}
