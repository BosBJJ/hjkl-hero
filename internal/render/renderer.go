package render

import (
	"strconv"
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/storage"
	"github.com/BosBJJ/hjkl-hero/internal/style"
)

// Renders whats within cameras view
func Render(gs game.GameState, cam game.Camera, themeID storage.ThemeID) string {
	theme := style.MakeStyle(themeID)
	lines := game.ToLines(gs)
	playerY := gs.Player.Line
	playerX := gs.Player.Column
	if playerY < 0 || playerY >= len(lines) {
		return string(gs.MapInfo.LevelMap)
	}
	top := cam.Y
	bottom := cam.Y + cam.Height
	left := cam.X
	right := cam.X + cam.Width
	RuneMap := make([][]rune, len(lines))
	for i, line := range lines {
		RuneMap[i] = []rune(line)
	}
	if bottom > len(RuneMap) {
		bottom = len(RuneMap)
	}
	if top < 0 {
		top = 0
	}

	if left < 0 {
		left = 0
	}
	var rendered strings.Builder
	for y := top; y < bottom; y++ {
		row := RuneMap[y]
		rowRight := right
		if rowRight > len(row) {
			rowRight = len(row)
		}
		for x := left; x < rowRight; x++ {
			rune := RuneMap[y][x]
			enemy, isEnemy := gs.EnemyAt(y, x)
			switch {
			case y == playerY && x == playerX:
				if gs.MapInfo.MapType == game.EditorMap {
					rendered.WriteString(style.CursorStyle.Render(string(rune)))
				} else {
					rendered.WriteString(theme.PlayerStyle)
				}
			case isEnemy:
				if enemy.EnemyType == game.Chaser {
					rendered.WriteString(style.ChaserStyle.Render(strconv.Itoa(enemy.Health)))
				}
				if enemy.EnemyType == game.Normal {
					rendered.WriteString(style.MeleerStyle.Render("M"))
				}
				if enemy.EnemyType == game.Tank {
					rendered.WriteString(style.ZanthStyle.Render("Z"))
				}
			case rune == '.':
				if gs.MapInfo.MapType == game.EditorMap {
					rendered.WriteString(string(rune))
				} else {
					rendered.WriteString(theme.FloorStyle)
				}
			case rune == '^':
				rendered.WriteString(style.StairStyle.Render("^"))
			default:
				if gs.MapInfo.MapType == game.EditorMap {
					rendered.WriteString(string(rune))
				} else {
					rendered.WriteString(theme.WallStyle)
				}
			}
		}
		rendered.WriteByte('\n')
	}
	return rendered.String()
}
