package render

import (
	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/levels"
)

func Render(state game.GameState, sMap levels.LevelMap) string {
	lines := levels.ToLines(sMap)
	currLine := state.Player.Line
	currRow := state.Player.Column
	if currLine < 0 || currLine > len(lines) {
		return string(sMap)
	}
	runes := []rune(lines[currLine])
	if currRow < 0 || currRow > len(runes) {
		return string(sMap)
	}
	runes[currRow] = '@'
	lines[currLine] = string(runes)
	return levels.ToText(lines)
}
