package ui

import (
	"github.com/BosBJJ/hjkl-hero/internal/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GameOverModel struct {
	Cursor   int
	Options  []string
	Selected int
	height   int
	width    int
}

func MakeGameOver() GameOverModel {
	return GameOverModel{
		Options:  []string{"Save to Leaderboard", "Quit"},
		Selected: -1,
	}
}

func (m GameOverModel) UpdateGameOver(msg tea.Msg) (GameOverModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "j":
			if m.Cursor < len(m.Options)-1 {
				m.Cursor++
			}
		case "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "enter":
			m.Selected = m.Cursor
		}
	}
	return m, nil
}

func (m GameOverModel) ViewGameOver() string {
	const Title = `
 ██████╗  █████╗ ███╗   ███╗███████╗     ██████╗ ██╗   ██╗███████╗██████╗ 
██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ██╔═══██╗██║   ██║██╔════╝██╔══██╗
██║  ███╗███████║██╔████╔██║█████╗      ██║   ██║██║   ██║█████╗  ██████╔╝
██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║   ██║╚██╗ ██╔╝██╔══╝  ██╔══██╗
╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ╚██████╔╝ ╚████╔╝ ███████╗██║  ██║
 ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝   ╚═══╝  ╚══════╝╚═╝  ╚═╝
                                                                          `
	title := style.MenuTitleStyle.
		Width(m.width).
		Align(lipgloss.Center).
		Render(Title) + "\n"

	optionBoxes := []string{}

	for i, option := range m.Options {
		if m.Cursor == i {
			optionBoxes = append(optionBoxes, style.CurrentOptionStyle.
				Align(lipgloss.Center).
				AlignVertical(lipgloss.Center).
				Width(80).
				Height(3).
				Render(option)+"\n")
		} else {
			optionBoxes = append(optionBoxes, style.OptionsStyle.
				Align(lipgloss.Center).
				AlignVertical(lipgloss.Center).
				Width(80).
				Height(3).
				Render(option)+"\n")
		}
	}
	gameOverMessage := lipgloss.JoinVertical(lipgloss.Center, append([]string{title}, optionBoxes...)...)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top, "\n\n\n\n\n\n\n\n\n\n"+gameOverMessage)
}
