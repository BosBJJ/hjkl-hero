package game

func (p *Position) Move(direction string, gs GameState) {
	lines := ToLines(gs)
	newPos := *p
	currLineLen := []rune(lines[p.Line])
	switch direction {
	case "j":
		if newPos.Line < len(lines)-1 {
			newPos.Line++
		}
	case "k":
		if newPos.Line > 1 {
			newPos.Line--
		}
	case "h":
		if newPos.Column > 1 {
			newPos.Column--
		}
	case "l":
		if newPos.Column < len(currLineLen)-1 {
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
