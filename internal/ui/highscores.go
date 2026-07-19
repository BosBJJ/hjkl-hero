package ui

import (
	"fmt"

	"github.com/BosBJJ/hjkl-hero/internal/storage"
	"github.com/BosBJJ/hjkl-hero/internal/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type HighScoresModel struct {
	width    int
	height   int
	Cursor   int
	Offset   int
	Scores   []storage.Run
	Options  []string
	Selected int
	Page     int
}

func MakeHighScores() HighScoresModel {
	return HighScoresModel{
		Options:  []string{"Return To Main Menu"},
		Selected: -1,
		Page: 1,
	}
}

func (m HighScoresModel) UpdateHighScores(msg tea.Msg) (HighScoresModel, tea.Cmd) {
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
		case "h":
			if m.Offset >= 5 {
				m.Offset -= 5
				m.Page--
			}
		case "l":
			if m.Offset < len(m.Scores)-1 {
				m.Offset += 5
				m.Page++
			}
		case "enter":
			m.Selected = m.Cursor
		}
	}
	return m, nil
}

func parseScore(r storage.Run) string {
	return fmt.Sprintf("%-20s | %5d | %9d | %5d | %13d | %-8s | %-20s |", r.PlayerName, r.Kills, r.TotalXp, r.TotalMoves, r.MapLevel, r.GameMode, r.FinishedAt)
}

func (m HighScoresModel) ViewHighScores() string {
	const Title = `
██╗  ██╗██╗ ██████╗ ██╗  ██╗    ███████╗ ██████╗ ██████╗ ██████╗ ███████╗███████╗
██║  ██║██║██╔════╝ ██║  ██║    ██╔════╝██╔════╝██╔═══██╗██╔══██╗██╔════╝██╔════╝
███████║██║██║  ███╗███████║    ███████╗██║     ██║   ██║██████╔╝█████╗  ███████╗
██╔══██║██║██║   ██║██╔══██║    ╚════██║██║     ██║   ██║██╔══██╗██╔══╝  ╚════██║
██║  ██║██║╚██████╔╝██║  ██║    ███████║╚██████╗╚██████╔╝██║  ██║███████╗███████║
╚═╝  ╚═╝╚═╝ ╚═════╝ ╚═╝  ╚═╝    ╚══════╝ ╚═════╝ ╚═════╝ ╚═╝  ╚═╝╚══════╝╚══════╝
                                                                                 `
	title := style.MenuTitleStyle.
		Width(m.width).
		Align(lipgloss.Center).
		Render(Title) + "\n"
	header := fmt.Sprintf("%-20s | %5s | %9s | %5s | %13s | %-8s | %-20s |",
		"Name",
		"Kills",
		"XP Gained",
		"Moves",
		"Map Level",
		"Mode",
		"Date Logged",
	)
	scores := []string{}
	scores = append(scores,
		style.HeaderStyle.
			Width(110).
			Align(lipgloss.Center).
			Render(header))
	optionBoxes := []string{}
	const visible = 5
	end := min(m.Offset+visible, len(m.Scores))
	for _, score := range m.Scores[m.Offset:end] {
		scores = append(scores,
			style.HSStyle.
				Width(110).
				Align(lipgloss.Center).
				Render(parseScore(score)),
		)
	}

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
	content := []string{title}
	pgInfo := fmt.Sprintf("Page %v", m.Page)

	content = append(content, scores...)
	content = append(content, optionBoxes...)
	content = append(content, "Press H and L to scroll through pages")
	content = append(content, pgInfo)

	menu := lipgloss.JoinVertical(lipgloss.Center, content...)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top, "\n\n\n\n"+menu)
}
