package game

type Position struct {
	Line   int
	Column int
}

type GameState struct {
	Player   Position
	Enemies  []EnemyInfo
	MapInfo  MapInfo
	undoSnap []SnapShot
	redoSnap []SnapShot
}

type SnapShot struct {
	PlayerSnapShot Position
	MapSnapShot    []string
}

type EnemyInfo struct {
	EnemyType
	Location Position
	Health   int
}
