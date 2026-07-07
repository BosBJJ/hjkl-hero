package game

type Position struct {
	Line   int
	Column int
}

type GameState struct {
	Player   Position
	MapInfo  MapInfo
	undoSnap []SnapShot
	redoSnap []SnapShot
}

type SnapShot struct {
	PlayerSnapShot Position
	MapSnapShot    []string
}
