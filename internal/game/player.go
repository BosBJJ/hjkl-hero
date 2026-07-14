package game

func (p *Position) Move(direction string, gs GameState) {
	lines := ToLines(gs)
	newPos := *p
	initLine := []rune(lines[p.Line])
	switch direction {
	case "j":
		if newPos.Line < len(lines)-1 {
			newPos.Line++
			updatedLine := []rune(lines[newPos.Line])
			if len(initLine) > len(updatedLine) && newPos.Column > len(updatedLine)-1 {
				newPos.Column = len(updatedLine) - 1
			}
		}
	case "k":
		if newPos.Line > 1 {
			newPos.Line--
			updatedLine := []rune(lines[newPos.Line])
			if len(initLine) > len(updatedLine) && newPos.Column > len(updatedLine)-1 {
				newPos.Column = len(updatedLine) - 1
			}
		}
	case "h":
		if newPos.Column > 1 {
			newPos.Column--
		}
	case "l":
		if newPos.Column < len(initLine)-1 {
			newPos.Column++
		}
	}
	if gs.MapInfo.MapType == RoomMap {
		if IsWall(gs, newPos.Line, newPos.Column) {
			return
		}
	}
	*p = newPos
}
func (p *Position) AdjustPlayer(lines []string) {
	if len(lines) == 0 {
		p.Line = 0
		p.Column = 0
		return
	}
	if p.Line >= len(lines) {
		p.Line = len(lines) - 1
	}
	if p.Line < 0 {
		p.Line = 0
	}
	runes := []rune(lines[p.Line])
	if len(runes) == 0 {
		p.Column = 0
		return
	}
	if p.Column >= len(runes) {
		p.Column = len(runes) - 1
	}
	if p.Column < 0 {
		p.Column = 0
	}
}

func (gs *GameState) SpawnPlayer() Position {
	var pos Position
	if gs.MapInfo.MapType == EditorMap {
		pos.Line = 1
		pos.Column = 1
		return pos
	}
	lines := ToLines(*gs)
	for lineNum, line := range lines {
		for col := range line {
			_, hasEnemy := gs.EnemyAt(lineNum, col)
			if hasEnemy {
				continue
			}
			if IsWall(*gs, lineNum, col) {
				continue
			}
			pos.Line = lineNum
			pos.Column = col
			return pos
		}
	}
	return pos
}

func (gs *GameState) LevelStats(input string) {
	if gs.Stats.XPGained <= 9 {
		return
	}
	switch input {
	case "h":
		gs.Stats.MaxHealth += 2
		gs.Stats.CurrentHealth += 1
	case "d":
		gs.Stats.BaseDmg += 3
	case "c":
		gs.Stats.CritChance += 2
	case "m":
		gs.Stats.BaseCritMulti += 1
	}
	gs.Stats.XPGained -= 10
}
