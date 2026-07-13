package render

import (
	"strconv"
	"strings"

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
	height, _ := game.GetMapSize(gs)
	RuneMap := make([][]rune, height)
	for i, line := range lines {
		RuneMap[i] = []rune(line)
	}
	var rendered strings.Builder
	for h, row := range RuneMap {
		for w, rune := range row {
			enemy, isEnemy := gs.EnemyAt(h, w)
			switch {
			case h == playerX && w == playerY:
				if gs.MapInfo.MapType == game.EditorMap {
					rendered.WriteString(style.CursorStyle.Render(string(rune)))
				} else {
					rendered.WriteString(style.PlayerStyle.Render(string('@')))
				}
			case isEnemy:
				if enemy.EnemyType == game.Chaser {
					rendered.WriteString(style.ChaserStyle.Render(strconv.Itoa(enemy.Health)))
				}
				if enemy.EnemyType == game.Normal {
					rendered.WriteString(style.MeleerStyle.Render("M"))
				}
				if enemy.EnemyType == game.Tank{
					rendered.WriteString(style.MeleerStyle.Render("Z"))
				}
			default:
				rendered.WriteString(string(rune))
			}
		}
		rendered.WriteString("\n")
	}

	return rendered.String()
}
