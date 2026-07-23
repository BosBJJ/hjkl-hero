package render

import (
	"strconv"
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/game"
	"github.com/BosBJJ/hjkl-hero/internal/style"
)

// Has to be outside of the func so it doesn't call style.X.Render 60 thousand times and lag the game
var (
	playerStyle = style.PlayerStyle.Render("@")
	wallStyle   = style.WallStyle.Render("\u2593") // #, \u2588, \u2593, \u2592
	floorStyle  = style.FloorStyle.Render(".")     //"." or " "
	stairStyle  = style.StairStyle.Render("^")
)

// Renders whats within cameras view
func Render(gs game.GameState, cam game.Camera) string {
	lines := game.ToLines(gs)
	playerX := gs.Player.Line
	playerY := gs.Player.Column
	if playerX < 0 || playerX >= len(lines) {
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
			case y == playerX && x == playerY:
				if gs.MapInfo.MapType == game.EditorMap {
					rendered.WriteString(style.CursorStyle.Render(string(rune)))
				} else {
					rendered.WriteString(playerStyle)
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
					rendered.WriteString(floorStyle)
				}
			case rune == '^':
				rendered.WriteString(stairStyle)
			default:
				if gs.MapInfo.MapType == game.EditorMap {
					rendered.WriteString(string(rune))
				} else {
					rendered.WriteString(wallStyle)
				}
			}
		}
		rendered.WriteByte('\n')
	}
	return rendered.String()
}
