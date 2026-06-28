package game

func (p *Position) Move(direction string, maxLine, maxCol int) {
	switch direction {
	case "j":
		if p.Line < maxLine {
			p.Line++
		}
	case "k":
		if p.Line > 0 {
			p.Line--
		}
	case "h":
		if p.Column > 0 {
			p.Column--
		}
	case "l":
		if p.Column < maxCol-1 {
			p.Column++
		}
	}
}
