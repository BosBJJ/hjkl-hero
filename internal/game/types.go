package game

type Position struct {
	Line   int
	Column int
}

type GameState struct {
	Player   Position
	SnapShot Position
	MapInfo  MapInfo
}
