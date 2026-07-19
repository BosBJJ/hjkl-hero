package ui

import (
	"fmt"

	"github.com/BosBJJ/hjkl-hero/internal/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GameOverModel struct {
	Cursor       int
	Options      []string
	Selected     int
	height       int
	width        int
	GameOverMode GameOverMode
	PlayerName   string
	Stats        RunStats
}
type GameOverMode int

const (
	MenuMode GameOverMode = iota
	EntryMode
)

func MakeGameOver() GameOverModel {
	return GameOverModel{
		Options:  []string{"Save to Leaderboard", "Quit"},
		Selected: -1,
	}
}

func (m GameOverModel) UpdateGameOver(msg tea.Msg) (GameOverModel, tea.Cmd) {
	switch m.GameOverMode {
	case MenuMode:
		return m.updateMenu(msg)
	case EntryMode:
		return m.updateNameEntry(msg)
	}
	return m, nil
}
func (m GameOverModel) updateMenu(msg tea.Msg) (GameOverModel, tea.Cmd) {
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
			if m.Cursor == 0 {
				m.GameOverMode = EntryMode
				return m, nil
			} else {
				m.Selected = m.Cursor
			}
		}
	}
	return m, nil
}

func (m GameOverModel) updateNameEntry(msg tea.Msg) (GameOverModel, tea.Cmd) {
	key := msg.(tea.KeyMsg)
	switch key.String() {
	case "enter":
		m.GameOverMode = MenuMode
		m.Selected = m.Cursor
	case "backspace":
		if len(m.PlayerName) > 0 {
			m.PlayerName = m.PlayerName[:len(m.PlayerName)-1]
		}
	default:
		m.PlayerName += key.String()
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
	var gameOverMessage string
	switch m.GameOverMode {
	case MenuMode:
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
		gameOverMessage = lipgloss.JoinVertical(lipgloss.Center, append([]string{title}, optionBoxes...)...)
	case EntryMode:
		name := fmt.Sprintf("Player Name: %v", m.PlayerName)
		stats := []string{
			fmt.Sprintf("Kills: %v", m.Stats.Kills),
			fmt.Sprintf("XP: %v", m.Stats.TotalXp),
			fmt.Sprintf("Moves: %v", m.Stats.TotalMoves),
			fmt.Sprintf("Level: %v", m.Stats.MapLevel),
		}
		for _, stat := range stats {
			optionBoxes = append(optionBoxes, style.HSStyle.
				Align(lipgloss.Center).
				AlignVertical(lipgloss.Center).
				Width(80).
				Height(3).
				Render(stat)+"\n")
		}
		optionBoxes = append(optionBoxes, style.CurrentOptionStyle.
			Align(lipgloss.Center).
			AlignVertical(lipgloss.Center).
			Width(80).
			Height(3).
			Render(name)+"\n")
		gameOverMessage = lipgloss.JoinVertical(lipgloss.Center, append([]string{title}, optionBoxes...)...)
	}

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top, "\n\n\n\n\n\n\n\n\n\n"+gameOverMessage)
}
