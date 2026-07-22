package ui

import (
	"fmt"

	"database/sql"
	"github.com/BosBJJ/hjkl-hero/internal/storage"
	"github.com/BosBJJ/hjkl-hero/internal/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SettingsModel struct {
	width        int
	height       int
	Cursor       int
	Options      []string
	Selected     int
	ModeMenu     bool
	OptionsMode  OptionsMode
	ModeSelected storage.GameMode
	DB           *sql.DB
}

type OptionsMode int

const (
	OptionMenuMode OptionsMode = iota
	GameTypePickerMode
	StylePickerMode
)

func MakeSettingsModel(db *sql.DB, settings storage.Settings) SettingsModel {
	return SettingsModel{
		DB:           db,
		Options:      []string{"Mode Type", "Cursor Style", "Exit"},
		Selected:     -1,
		ModeSelected: settings.GameMode,
	}
}

func (m SettingsModel) UpdateSettings(msg tea.Msg) (SettingsModel, tea.Cmd) {
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
			switch m.OptionsMode {
			case OptionMenuMode:
				if m.Cursor == 0 {
					m.OptionsMode = GameTypePickerMode
					return m, nil
				}
				if m.Cursor == 1 {
					m.OptionsMode = StylePickerMode
					return m, nil
				}
				m.Selected = m.Cursor
			case GameTypePickerMode:
				if m.Cursor == 0 {
					m.ModeSelected = storage.TutorialMode
					storage.UpdateGameMode(m.DB, storage.TutorialMode)
				}
				if m.Cursor == 1 {
					m.ModeSelected = storage.RogueLikeMode
					storage.UpdateGameMode(m.DB, storage.RogueLikeMode)
				}
				if m.Cursor == 2 {
					m.OptionsMode = OptionMenuMode
					m.Cursor = 0
				}
			case StylePickerMode:
				if m.Cursor == 2 {
					m.OptionsMode = OptionMenuMode
					m.Cursor = 0
				}
			}
		}
	}
	return m, nil
}

func (m SettingsModel) ViewSettings() string {
	const Title = `
 ██████╗ ██████╗ ████████╗██╗ ██████╗ ███╗   ██╗███████╗
██╔═══██╗██╔══██╗╚══██╔══╝██║██╔═══██╗████╗  ██║██╔════╝
██║   ██║██████╔╝   ██║   ██║██║   ██║██╔██╗ ██║███████╗
██║   ██║██╔═══╝    ██║   ██║██║   ██║██║╚██╗██║╚════██║
╚██████╔╝██║        ██║   ██║╚██████╔╝██║ ╚████║███████║
 ╚═════╝ ╚═╝        ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚══════╝
                                                        `
	title := style.MenuTitleStyle.
		Width(m.width).
		Align(lipgloss.Center).
		Render(Title) + "\n"

	optionBoxes := []string{}
	switch m.OptionsMode {
	case OptionMenuMode:
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
	case GameTypePickerMode:
		options := []string{"Tutorial", "Rogue", "Back"}
		for i, option := range options {
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
	}
	currMode := fmt.Sprintf("Currently Selected Game Mode: %v", m.ModeSelected)
	optionBoxes = append(optionBoxes, currMode)
	menu := lipgloss.JoinVertical(lipgloss.Center, append([]string{title}, optionBoxes...)...)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top, "\n\n\n\n"+menu)
}
