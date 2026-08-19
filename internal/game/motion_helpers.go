package game

// Jumps to either next new word OR punctuation
func (gs *GameState) nextWordPos() (Position, bool) {
	wordPos := Position{}
	lines := ToLines(*gs)
	col := gs.Player.Column
	for line := gs.Player.Line; line < len(lines); line++ {
		runes := []rune(lines[line])
		for i := col + 1; i < len(runes); i++ {
			currRune := runes[i-1]
			nextRune := runes[i]
			if !isSpaceOrSymbol(currRune) && isSymbol(nextRune) {
				wordPos.Column = i
				wordPos.Line = line
				return wordPos, true
			}
			if isSpaceOrSymbol(currRune) && !isSpaceOrSymbol(nextRune) {
				wordPos.Column = i
				wordPos.Line = line
				return wordPos, true
			}
		}
		col = 0
	}
	return wordPos, false
}

// Only jumps to next word after a space
func (gs *GameState) nextWORDPos() (Position, bool) {
	wordPos := Position{}
	lines := ToLines(*gs)
	col := gs.Player.Column
	leftWord := false
	for line := gs.Player.Line; line < len(lines); line++ {
		runes := []rune(lines[line])
		for i := col; i < len(runes); i++ {
			currRune := runes[i]
			switch {
			case !leftWord && currRune == ' ':
				leftWord = true
			case leftWord && !isSpaceOrSymbol(currRune):
				wordPos.Column = i
				wordPos.Line = line
				return wordPos, true
			}
		}
		leftWord = true
		col = 0
	}
	return wordPos, false
}

// Jumps to end of next word
func (gs *GameState) endWORDPos() (Position, bool) {
	wordPos := Position{}
	lines := ToLines(*gs)
	col := gs.Player.Column
	for line := gs.Player.Line; line < len(lines); line++ {
		runes := []rune(lines[line])
		leftWord := runes[col] != ' '
		if leftWord && (col == len(runes)-1 || runes[col+1] == ' ') {
			leftWord = false
		}
		for i := col + 1; i < len(runes); i++ {
			currRune := runes[i-1]
			nextRune := runes[i]
			switch {
			case i == len(runes)-1:
				wordPos.Column = i
				wordPos.Line = line
				return wordPos, true
			case !leftWord && nextRune == ' ' && currRune != ' ':
				leftWord = true
			case !leftWord && currRune == ' ':
				leftWord = true
			case leftWord && nextRune == ' ':
				wordPos.Column = i - 1
				wordPos.Line = line
				return wordPos, true
			}
		}
		leftWord = true
		col = 1
	}
	return wordPos, false
}

// Jumps to end of word or next punct
func (gs *GameState) endWordPos() (Position, bool) {
	wordPos := Position{}
	lines := ToLines(*gs)
	col := gs.Player.Column
	for line := gs.Player.Line; line < len(lines); line++ {
		runes := []rune(lines[line])
		leftWord := runes[col] != ' '
		if leftWord && (col == len(runes)-1 || isSpaceOrSymbol(runes[col+1])) {
			leftWord = false
		}
		for i := col + 1; i < len(runes); i++ {
			currRune := runes[i-1]
			nextRune := runes[i]
			switch {
			case !leftWord && nextRune == ' ' && currRune != ' ':
				leftWord = true
			case !leftWord && currRune == ' ':
				leftWord = true
			case leftWord && nextRune == ' ':
				wordPos.Column = i - 1
				wordPos.Line = line
				return wordPos, true
			case leftWord && isSymbol(nextRune):
				wordPos.Column = i - 1
				wordPos.Line = line
				return wordPos, true
			case isSymbol(nextRune):
				wordPos.Column = i
				wordPos.Line = line
				return wordPos, true
			case i == len(runes)-1:
				wordPos.Column = i
				wordPos.Line = line
				return wordPos, true
			}
		}
		leftWord = true
		col = 1
	}
	return wordPos, false
}

// Jumps to last beginning of word or punctuation
func (gs *GameState) backWordPos() (Position, bool) {
	wordPos := Position{}
	lines := ToLines(*gs)
	col := gs.Player.Column
	for line := gs.Player.Line; line >= 0; line-- {
		runes := []rune(lines[line])
		for i := col - 1; i >= 0; i-- {
			if isSymbol(runes[i]) {
				wordPos.Column = i
				wordPos.Line = line
				return wordPos, true
			}
			if (i == 0 || isSpaceOrSymbol(runes[i-1])) && !isSpaceOrSymbol(runes[i]) {
				wordPos.Column = i
				wordPos.Line = line
				return wordPos, true
			}
		}
		if line > 0 {
			col = len([]rune(lines[line-1]))
		}
	}
	return wordPos, false
}

// Jump to last word
func (gs *GameState) backWORDPos() (Position, bool) {
	wordPos := Position{}
	lines := ToLines(*gs)
	col := gs.Player.Column
	for line := gs.Player.Line; line >= 0; line-- {
		runes := []rune(lines[line])
		for i := col - 1; i >= 0; i-- {
			if (i == 0 || runes[i-1] == ' ') && !isSpaceOrSymbol(runes[i]) {
				wordPos.Column = i
				wordPos.Line = line
				return wordPos, true
			}
		}
		if line > 0 {
			col = len([]rune(lines[line-1]))
		}
	}
	return wordPos, false
}
