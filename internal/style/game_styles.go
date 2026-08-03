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

// Zanth is a friend who loves playing tank, so this is a mob dedicated to him
var ZanthStyle = lipgloss.NewStyle().
	Background(Green).
	Foreground(Black)

var PotionStyle = lipgloss.NewStyle().
	Foreground(RedDark).
	Background(Yellow)

var StairStyle = lipgloss.NewStyle().
	Background(White).
	Foreground(Brown)

var ChestStyle = lipgloss.NewStyle().
	Background(Brown).
	Foreground(Golden)

var VendorStyle = lipgloss.NewStyle().
	Background(Golden).
	Foreground(Black)

// Health Styles
var HealthBase = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	Width(30).
	Height(1).
	Align(lipgloss.Center).
	AlignVertical(lipgloss.Center).
	Bold(true)
var FullHealth = HealthBase.
	Background(Green).
	BorderForeground(Green).
	BorderBackground(Green).
	Foreground(Magenta)
var LowerHealth = HealthBase.
	Background(Yellow).
	BorderForeground(Yellow).
	BorderBackground(Yellow).
	Foreground(Magenta)
var LowHealth = HealthBase.
	Background(Red).
	BorderForeground(Red).
	BorderBackground(Red).
	Foreground(White)
