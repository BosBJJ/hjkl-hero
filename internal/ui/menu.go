package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type MenuModel struct {
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
	s := "Welcome to HJKL Hero!\n"

	for i, option := range m.Options {
		cursor := " "
		if m.Cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%v %v\n", cursor, option)
	}
	s += "Press ctrl+c to quit\n"
	return s
}
