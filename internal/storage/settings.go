package storage

type ThemeID string

const (
	DefaultTheme ThemeID = "default"
	RedTheme     ThemeID = "red"
	WinterTheme  ThemeID = "christmas"
	CyberTheme   ThemeID = "cyberpunk"
	CustomTheme  ThemeID = "custom"
)

type GameMode string

const (
	TutorialMode  GameMode = "tutorial"
	RogueLikeMode GameMode = "rogue"
)

type Theme struct {
	WallColor   string
	FloorColor  string
	PlayerColor string

	WallIcon   string
	FloorIcon  string
	PlayerIcon string
}
