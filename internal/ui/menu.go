package ui

import (
	"github.com/BosBJJ/hjkl-hero/internal/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MenuModel struct {
	width    int
	height   int
	Cursor   int
	Options  []string
	Selected int
}

func MakeMenu() MenuModel {
	return MenuModel{
		Options:  []string{"Play", "Leaderboards", "Options", "Quit"},
		Selected: -1,
	}
}

func (m MenuModel) UpdateMenu(msg tea.Msg) (MenuModel, tea.Cmd) {
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

func (m MenuModel) ViewMenu() string {
	const Title = `
██╗  ██╗     ██╗██╗  ██╗██╗         ██╗  ██╗███████╗██████╗  ██████╗ 
██║  ██║     ██║██║ ██╔╝██║         ██║  ██║██╔════╝██╔══██╗██╔═══██╗
███████║     ██║█████╔╝ ██║         ███████║█████╗  ██████╔╝██║   ██║
██╔══██║██   ██║██╔═██╗ ██║         ██╔══██║██╔══╝  ██╔══██╗██║   ██║
██║  ██║╚█████╔╝██║  ██╗███████╗    ██║  ██║███████╗██║  ██║╚██████╔╝
╚═╝  ╚═╝ ╚════╝ ╚═╝  ╚═╝╚══════╝    ╚═╝  ╚═╝╚══════╝╚═╝  ╚═╝ ╚═════╝ 
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
	menu := lipgloss.JoinVertical(lipgloss.Center, append([]string{title}, optionBoxes...)...)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top, "\n\n\n\n"+menu)
}
