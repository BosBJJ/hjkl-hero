package game

import "math/rand/v2"

type EnemyType int

const (
	Normal EnemyType = iota
	Chaser
	Tank
)

func (e EnemyType) String() string {
	switch e {
	case Normal:
		return "Brawler"
	case Chaser:
		return "Chaser"
	case Tank:
		return "Tank"
	default:
		return "Invalid Enemy Type"
	}
}

func (gs *GameState) SpawnEnemy(maxEnemies int) {
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
		if _, ok := gs.EnemyAt(line, col); ok {
			continue
		}
		roll := rand.IntN(100)
		switch {
		case roll == 99:
			newSpawn = MakeTank(line, col)
		case roll >= 20 && roll < 99:
			newSpawn = MakeChaser(line, col)
		default:
			newSpawn = MakeMeleer(line, col)
		}
		if len(gs.Enemies) < maxEnemies {
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
	newChaser.Health = 9
	newChaser.BaseDmg = 1
	newChaser.Location.Line = line
	newChaser.Location.Column = col
	newChaser.EnemyType = Chaser
	return newChaser
}

func MakeMeleer(line, col int) EnemyInfo {
	var newMob EnemyInfo
	newMob.Health = 12
	newMob.BaseDmg = 3
	newMob.Location.Line = line
	newMob.Location.Column = col
	newMob.EnemyType = Normal
	return newMob
}

//Zanth is a friend who loves playing tank, so this is a mob dedicated to him
func MakeTank(line, col int) EnemyInfo {
	var newMob EnemyInfo
	newMob.Health = 40
	newMob.BaseDmg = 1
	newMob.Location.Line = line
	newMob.Location.Column = col
	newMob.EnemyType = Tank
	return newMob
}

func (gs *GameState) ChasePlayer() string {
	combatMsg := ""
	for i := len(gs.Enemies) - 1; i >= 0; i-- {
		enemy := &gs.Enemies[i]
		clone := *enemy
		lineDiff := getDiff(gs.Player.Line, enemy.Location.Line)
		colDiff := getDiff(gs.Player.Column, enemy.Location.Column)
		trueLineDiff := enemy.Location.Line - gs.Player.Line
		trueColDiff := enemy.Location.Column - gs.Player.Column
		switch {
		case colDiff <= lineDiff: //move the longest path to player to prevent spamming two buttons and making enemy move back/forth without getting closer
			clone.moveLine(trueLineDiff)
			if gs.BlockedTile(clone.Location.Line, clone.Location.Column, i) {
				clone = *enemy
				clone.moveCol(trueColDiff)
				if gs.BlockedTile(clone.Location.Line, clone.Location.Column, i) {
					clone = *enemy
				}
			}
		case colDiff > lineDiff:
			clone.moveCol(trueColDiff)
			if gs.BlockedTile(clone.Location.Line, clone.Location.Column, i) {
				clone = *enemy
				clone.moveLine(trueLineDiff)
				if gs.BlockedTile(clone.Location.Line, clone.Location.Column, i) {
					clone = *enemy
				}
			}
		}
		switch enemy.EnemyType {
		case Chaser:
			enemy.Health--
			if enemy.Health <= 0 {
				gs.Enemies = append(gs.Enemies[:i], gs.Enemies[i+1:]...)
				continue
			}
		case Tank:
			enemy.MoveCount++
			if enemy.MoveCount < 5 {
				clone = *enemy
			} else {
				enemy.MoveCount = 0
			}
		}
		if clone.Location == gs.Player {
			combatMsg = gs.TryDamagePlayer()
			clone = *enemy
		}
		enemy.Location = clone.Location
	}
	return combatMsg
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

func (gs GameState) BlockedTile(line, col, idx int) bool {
	if IsWall(gs, line, col) {
		return true
	}
	for i, enemy := range gs.Enemies {
		if i == idx {
			continue
		}
		if enemy.Location.Line == line && enemy.Location.Column == col {
			return true
		}
	}
	return false
}
