package game

import "math/rand/v2"

func (gs *GameState) SpawnEnemy() {
	height, width := GetMapSize(*gs)
	for {
		line := rand.IntN(height)
		col := rand.IntN(width)
		if gs.Player.Line == line && gs.Player.Column == col {
			continue
		}
		if IsWall(*gs, line, col) {
			continue
		}
		enemy := Position{
			Line:   line,
			Column: col,
		}
		if len(gs.Enemies) < 1 {
			gs.Enemies = append(gs.Enemies, enemy)
		}
		return
	}
}
