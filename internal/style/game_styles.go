package style

import "github.com/charmbracelet/lipgloss"

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("16")).
	Background(BabyBlue)

var PlayerStyle = lipgloss.NewStyle().
	Foreground(Red)

var ChaserStyle = lipgloss.NewStyle().
	Background(BabyBlue).
	Foreground(Black)

var MeleerStyle = lipgloss.NewStyle(). //Filler style for now
	Background(Red).
	Foreground(Peach)

var ZanthStyle = lipgloss.NewStyle(). //Zanth is a friend who loves playing tank, so this is a mob dedicated to him
	Background(Green).
	Foreground(Black)

var WallStyle = lipgloss.NewStyle().
	Foreground(GrayDark) // I like Magenta here

var FloorStyle = lipgloss.NewStyle().
	Foreground(Black) // I like Green here

var StairStyle = lipgloss.NewStyle().
	Background(Brown)
