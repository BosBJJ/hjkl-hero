package storage

type Theme string

const (
	DefaultTheme Theme = "default"
	RedTheme     Theme = "red"
)

type GameMode string

const (
	TutorialMode  GameMode = "tutorial"
	RogueLikeMode GameMode = "rogue"
)
