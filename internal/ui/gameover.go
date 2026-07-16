package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)


type GameOverModel struct {
	Cursor   int
	Options  []string
	Selected int
}

func MakeGameOver() GameOverModel {
	return GameOverModel{
		Options:  []string{"Save to Leaderboard", "Quit"},
		Selected: -1,
	}
}

func (m GameOverModel) UpdateGameOver(msg tea.Msg) (GameOverModel, tea.Cmd) {
	switch msg := msg.(type) {
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
	s := "Game Over!\n"

	for i, option := range m.Options {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%v %v\n", cursor, option)
	}
	return s
}
