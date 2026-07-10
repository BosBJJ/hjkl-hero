package render

import (
	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/style"
)

func Render(gs game.GameState) string {
	lines := game.ToLines(gs)
	playerX := gs.Player.Line
	playerY := gs.Player.Column
	if playerX < 0 || playerX >= len(lines) {
		return string(gs.MapInfo.LevelMap)
	}
	currLine := lines[playerX]
	runes := []rune(currLine)
	if playerY < 0 || playerY >= len(runes) {
		return string(gs.MapInfo.LevelMap)
	}
	var line string
	if gs.MapInfo.MapType == game.EditorMap {
		line = string(runes[:playerY]) + style.CursorStyle.Render(string(runes[playerY])) + string(runes[playerY+1:])
	} else {
		line = string(runes[:playerY]) + style.PlayerStyle.Render("@") + string(runes[playerY+1:])
	}
	lines[playerX] = line
	return game.ToText(lines)
}
