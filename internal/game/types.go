package game

type Position struct {
	Line   int
	Column int
}

type GameState struct {
	Player   Position
	Stats    PlayerInfo
	Enemies  []EnemyInfo
	MapInfo  MapInfo
	undoSnap []SnapShot
	redoSnap []SnapShot
}

type PlayerInfo struct {
	MaxHealth     int
	CurrentHealth int
	BaseDmg       int
	CritChance    int
	BaseCritMulti int
	XPGained      int
}

type SnapShot struct {
	PlayerSnapShot Position
	MapSnapShot    []string
}

type EnemyInfo struct {
	EnemyType
	Location  Position
	BaseDmg   int
	Health    int
	MoveCount int
}

type CombatLog struct {
	EnemyType
	Hit         bool
	EnemyKilled bool
	Critical    bool
	DamageDealt int
	Experience  int
	AttackStyle AttackType
}
