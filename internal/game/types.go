package game

type Position struct {
	Line   int
	Column int
}

type GameState struct {
	Player   Position
	Stats    PlayerInfo
	Enemies  []EnemyInfo
	Items    []Item
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
	TotalXP       int
	Kills         int
	DamageTaken   int
	Inventory     []Item
	Gold          int
	PlayerLevel   int
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

type ItemType int

const (
	HealthPotion ItemType = iota
	Gold         ItemType = iota
)

type Item struct {
	Type   ItemType
	Amount int
	Line   int
	Col    int
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

type Camera struct {
	X      int
	Y      int
	Width  int
	Height int
}
