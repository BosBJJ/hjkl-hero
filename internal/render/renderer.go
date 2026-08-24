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
			item, isItem := gs.ItemAt(y, x)
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
			case isItem:
				switch item.Type {
				case game.HealthPotion:
					rendered.WriteString(style.PotionStyle.Render("P"))
				case game.Gold:
					rendered.WriteString(style.VendorStyle.Render("C"))
				}
			case rune == '.':
				if gs.MapInfo.MapType == game.EditorMap {
					rendered.WriteString(string(rune))
				} else {
					rendered.WriteString(theme.FloorStyle)
				}
			case rune == '^': //Stairs
				rendered.WriteString(style.StairStyle.Render("^"))
			case rune == '+': //Chest
				if gs.MapInfo.MapType == game.RoomMap {
					rendered.WriteString(style.ChestStyle.Render("\u233A"))
				} else {
					rendered.WriteString(string(rune))
				}

			case rune == '$': //Vendor
				if gs.MapInfo.MapType == game.RoomMap {
					rendered.WriteString(style.VendorStyle.Render("$"))
				} else {
					rendered.WriteString(string(rune))
				}

			default:
				if gs.MapInfo.MapType == game.EditorMap {
					rendered.WriteString(string(rune))
				} else {
					rendered.WriteString(theme.WallStyle)
				}
			}
		}
		if y == playerY && playerX == len(row) && gs.MapInfo.MapType == game.EditorMap {
			rendered.WriteString(style.CursorStyle.Render(string(" ")))
		}
		rendered.WriteByte('\n')
	}
	return rendered.String()
}
