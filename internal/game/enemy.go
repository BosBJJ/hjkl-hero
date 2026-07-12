package game

import "math/rand/v2"

type EnemyType int

const (
	Normal EnemyType = iota
	Chaser
)

func (gs *GameState) SpawnEnemy() {
	height, width := GetMapSize(*gs)
	var newSpawn EnemyInfo
	for {
		line := rand.IntN(height)
		col := rand.IntN(width)
		if gs.Player.Line == line && gs.Player.Column == col {
			continue
		}
		if IsWall(*gs, line, col) {
			continue
		}
		roll := rand.IntN(100)
		if roll < 80 {
			newSpawn = MakeChaser(line, col)
		} else {
			newSpawn = MakeMeleer(line, col)
		}
		if len(gs.Enemies) < 5 {
			gs.Enemies = append(gs.Enemies, newSpawn)
		}
		return
	}
}

func (gs GameState) EnemyAt(line, col int) (EnemyInfo, bool) {
	for _, enemy := range gs.Enemies {
		if enemy.Location.Line == line && enemy.Location.Column == col {
			return enemy, true
		}
	}
	return EnemyInfo{}, false
}

func MakeChaser(line, col int) EnemyInfo {
	var newChaser EnemyInfo
	newChaser.Location.Line = line
	newChaser.Location.Column = col
	newChaser.Health = 9
	newChaser.EnemyType = Chaser
	return newChaser
}

func MakeMeleer(line, col int) EnemyInfo {
	var newMob EnemyInfo
	newMob.Health = 12
	newMob.Location.Line = line
	newMob.Location.Column = col
	newMob.EnemyType = Normal
	return newMob
}

func (gs *GameState) ChasePlayer() {
	for i := len(gs.Enemies) - 1; i >= 0; i-- {
		enemy := &gs.Enemies[i]
		clone := *enemy
		lineDiff := getDiff(gs.Player.Line, enemy.Location.Line)
		colDiff := getDiff(gs.Player.Column, enemy.Location.Column)
		trueLineDiff := enemy.Location.Line - gs.Player.Line
		trueColDiff := enemy.Location.Column - gs.Player.Column
		switch {
		case colDiff < lineDiff: //move the longest path to player to prevent spamming two buttons and making enemy move back/forth without getting closer
			clone.moveLine(trueLineDiff)
			if IsWall(*gs, clone.Location.Line, clone.Location.Column) {
				clone = *enemy
				clone.moveCol(trueColDiff)
				if IsWall(*gs, clone.Location.Line, clone.Location.Column) {
					clone = *enemy
				}
			}
		case colDiff > lineDiff:
			clone.moveCol(trueColDiff)
			if IsWall(*gs, clone.Location.Line, clone.Location.Column) {
				clone = *enemy
				clone.moveLine(trueLineDiff)
				if IsWall(*gs, clone.Location.Line, clone.Location.Column) {
					clone = *enemy
				}
			}
		}
		enemy.Location = clone.Location
		if enemy.EnemyType == Chaser {
			enemy.Health--
			if enemy.Health == 0 {
				gs.Enemies = append(gs.Enemies[:i], gs.Enemies[i+1:]...)
			}
		}
	}
}

func (e *EnemyInfo) moveLine(diff int) {
	if diff < 0 {
		e.Location.Line++
	} else {
		e.Location.Line--
	}
}

func (e *EnemyInfo) moveCol(diff int) {
	if diff < 0 {
		e.Location.Column++
	} else {
		e.Location.Column--
	}
}
