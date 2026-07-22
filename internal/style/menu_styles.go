package style

import "github.com/charmbracelet/lipgloss"

var MenuTitleStyle = lipgloss.NewStyle().
	Foreground(Orange).
	Bold(true)

var CurrentOptionStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Background(Magenta).
	BorderForeground(Green).
	Foreground(Green).
	Bold(true)

var OptionsStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Background(Green).
	BorderForeground(Magenta).
	Foreground(Magenta).
	Bold(true)

var MenuBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Background(Black).
	BorderForeground(Green).
	Padding(1, 2).
	Align(lipgloss.Center)

var HSStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Background(BlueDark).
	BorderForeground(BlueDark).
	Foreground(Peach).
	Bold(true)

var HeaderStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Background(BlueLight).
	BorderForeground(Black).
	Foreground(Black).
	Bold(true)
