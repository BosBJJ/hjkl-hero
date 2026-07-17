package style

import "github.com/charmbracelet/lipgloss"

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("16")).
	Background(lipgloss.Color("87"))

var PlayerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("11"))

var ChaserStyle = lipgloss.NewStyle(). //Filler style for now
	Background(lipgloss.Color("87"))

var MeleerStyle = lipgloss.NewStyle(). //Filler style for now
	Background(lipgloss.Color("196"))

