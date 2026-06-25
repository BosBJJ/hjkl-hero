package game

type Position struct {
	Line   int
	Column int
}

type GameState struct {
	Player Position
	Level int
}
