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
