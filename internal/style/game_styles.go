package style

import "github.com/charmbracelet/lipgloss"

var CursorStyle = lipgloss.NewStyle().
	Foreground(Black).
	Background(BabyBlue).
	Bold(true)

var ChaserStyle = lipgloss.NewStyle().
	Background(BabyBlue).
	Foreground(Black)

var MeleerStyle = lipgloss.NewStyle().
	Background(Red).
	Foreground(Peach)

var ZanthStyle = lipgloss.NewStyle(). //Zanth is a friend who loves playing tank, so this is a mob dedicated to him
	Background(Green).
	Foreground(Black)

var PotionStyle = lipgloss.NewStyle().
	Foreground(RedDark).
	Background(Yellow)

var StairStyle = lipgloss.NewStyle().
	Background(White).
	Foreground(Brown)

var StatsStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Background(Green).
	BorderForeground(Magenta).
	Foreground(Magenta).
	Bold(true)
