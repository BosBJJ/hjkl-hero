package game

import (
	"fmt"
	"math/rand/v2"
)

func (gs *GameState) RangedAttack(direction string) CombatLog {
	var log CombatLog
	for i := len(gs.Enemies) - 1; i >= 0; i-- {
		enemy := &gs.Enemies[i]
		if enemy.EnemyType == Chaser {
			continue
		}
		enemyLine := enemy.Location.Line
		enemyCol := enemy.Location.Column
		playerLine := gs.Player.Line
		playerCol := gs.Player.Column
		lineDiff := enemyLine - playerLine
		colDiff := enemyCol - playerCol
		inputRune := []rune(direction)[0]
		switch inputRune {
		case 'h':
			if colDiff >= -4 && colDiff < 0 && lineDiff == 0 {
				log = gs.AttackMob(enemy, Ranged)
			}
			if log.Hit == false {
				continue
			}
		case 'l':
			if colDiff > 0 && colDiff <= 4 && lineDiff == 0 {
				log = gs.AttackMob(enemy, Ranged)
			}
			if log.Hit == false {
				continue
			}
		case 'j':
			if lineDiff <= 4 && lineDiff > 0 && colDiff == 0 {
				log = gs.AttackMob(enemy, Ranged)
			}
			if log.Hit == false {
				continue
			}
		case 'k':
			if lineDiff >= -4 && lineDiff < 0 && colDiff == 0 {
				log = gs.AttackMob(enemy, Ranged)
			}
			if log.Hit == false {
				continue
			}
		}
		if enemy.Health <= 0 {
			log.EnemyKilled = true
			xp := enemy.GetExperience(log.AttackStyle)
			log.Experience = xp
			gs.Stats.XPGained += xp
			gs.Stats.TotalXP += xp
			gs.Stats.Kills += 1
			gs.Enemies = append(gs.Enemies[:i], gs.Enemies[i+1:]...)
			break
		}
	}
	return log
}

func (gs *GameState) MeleeAttack() CombatLog {
	var log CombatLog
	for i := len(gs.Enemies) - 1; i >= 0; i-- {
		enemy := &gs.Enemies[i]
		if enemy.EnemyType == Chaser {
			continue
		}
		enemyLine := enemy.Location.Line
		enemyCol := enemy.Location.Column
		playerLine := gs.Player.Line
		playerCol := gs.Player.Column
		lineDiff := getDiff(enemyLine, playerLine)
		colDiff := getDiff(enemyCol, playerCol)
		if lineDiff <= 1 && colDiff == 0 || lineDiff == 0 && colDiff <= 1 {
			log = gs.AttackMob(enemy, Melee)
			if log.Hit == false {
				continue
			}
			if enemy.Health <= 0 {
				log.EnemyKilled = true
				xp := enemy.GetExperience(log.AttackStyle)
				gs.Stats.XPGained += xp
				gs.Stats.TotalXP += xp
				gs.Stats.Kills += 1
				log.Experience = xp
				gs.Enemies = append(gs.Enemies[:i], gs.Enemies[i+1:]...)
			}
			break
		}
	}
	return log
}

func (l CombatLog) ParseLog() string {
	switch {
	case !l.Hit:
		return "Attack missed!"
	case l.EnemyKilled && l.Critical:
		return fmt.Sprintf("Enemy %v obliterated from %v damage crit! (+%v XP)", l.EnemyType, l.DamageDealt, l.Experience)
	case l.EnemyKilled:
		return fmt.Sprintf("Killed enemy %v with %v damage (+%v XP)", l.EnemyType, l.DamageDealt, l.Experience)
	case l.Critical:
		return fmt.Sprintf("Massive critical hit! %v damage dealt to %v", l.DamageDealt, l.EnemyType)
	default:
		return fmt.Sprintf("%v damage dealt to %v", l.DamageDealt, l.EnemyType)
	}
}

type AttackType int

const (
	Melee AttackType = iota
	Ranged
)

func (e EnemyInfo) GetExperience(attack AttackType) int {
	xpDrop := 0
	switch e.EnemyType {
	case Normal:
		xpDrop = 4
	case Tank:
		xpDrop = 20
	}
	if attack == Ranged {
		xpDrop /= 2
	}
	return xpDrop
}

func (gs *GameState) AttackMob(e *EnemyInfo, style AttackType) CombatLog {
	var log CombatLog
	if !successfulHit() {
		log.Hit = false
		return log
	}
	damage := gs.Stats.BaseDmg
	log.AttackStyle = style
	if style == Ranged {
		damage /= 2
	}
	if crit, mulitplier := isCrit(*gs); crit {
		log.Critical = true
		damage = damage * mulitplier
	}
	e.Health -= damage
	log.EnemyType = e.EnemyType
	log.Hit = true
	log.DamageDealt = damage
	return log
}

func isCrit(gs GameState) (bool, int) {
	roll := rand.IntN(100)
	if roll <= gs.Stats.CritChance {
		return true, gs.Stats.BaseCritMulti
	}
	return false, 1
}

func (gs *GameState) TryDamagePlayer() string {
	hitMsg := "Enemy hit missed!"
	for i := len(gs.Enemies) - 1; i >= 0; i-- {
		enemy := &gs.Enemies[i]
		enemyLine := enemy.Location.Line
		enemyCol := enemy.Location.Column
		playerLine := gs.Player.Line
		playerCol := gs.Player.Column
		lineDiff := getDiff(enemyLine, playerLine)
		colDiff := getDiff(enemyCol, playerCol)
		if lineDiff <= 1 && colDiff <= 1 {
			switch enemy.EnemyType { //SWITCH SO LATER CAN ADD DIFFERENT TYPES OF DAMAGES
			case Normal:
				if successfulHit() {
					gs.Stats.CurrentHealth -= enemy.BaseDmg
					hitMsg = fmt.Sprintf("Enemy %v hit player for %v", enemy.EnemyType, enemy.BaseDmg)
				}
			case Chaser:
				if successfulHit() {
					gs.Stats.CurrentHealth -= enemy.BaseDmg
					hitMsg = fmt.Sprintf("You were hit by a %v for %v damage! Make sure to dodge!", enemy.EnemyType, enemy.BaseDmg)
					enemy.Health = 0
					gs.Enemies = append(gs.Enemies[:i], gs.Enemies[i+1:]...)
				}
			case Tank:
				if successfulHit() {
					gs.Stats.CurrentHealth -= enemy.BaseDmg
					hitMsg = fmt.Sprintf("Enemy %v hit player for %v", enemy.EnemyType, enemy.BaseDmg)
				}
			}
		}
	}
	return hitMsg
}

func successfulHit() bool {
	hitChance := rand.IntN(100) //70% chance to hit, 30% chance to miss
	if hitChance < 30 {
		return false
	}
	return true
}
