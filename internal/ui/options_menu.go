package ui

import (
	"database/sql"

	"github.com/BosBJJ/hjkl-hero/internal/storage"
	"github.com/BosBJJ/hjkl-hero/internal/style"
	tea "github.com/charmbracelet/bubbletea"
)

type SettingsModel struct {
	width         int
	height        int
	Cursor        int
	Options       []string
	Selected      int
	ModeMenu      bool
	OptionsMode   OptionsMode
	ModeSelected  storage.GameMode
	ThemeSelected storage.ThemeID
	DB            *sql.DB
	EditorMenu    editorMenu
}

type OptionsMode int

const (
	OptionMenuMode OptionsMode = iota
	GameTypePickerMode
	StylePickerMode

	StyleEditorMode
)

func MakeSettingsModel(db *sql.DB, settings storage.Settings) SettingsModel {
	editorMenu := editorMenu{
		LeftPanel: true,
		Theme:     style.Themes[storage.CustomTheme],
	}
	return SettingsModel{
		DB:            db,
		Options:       []string{"Mode Type", "Theme Picker", "Exit"},
		Selected:      -1,
		ModeSelected:  settings.GameMode,
		ThemeSelected: settings.ThemeID,
		EditorMenu:    editorMenu,
	}
}

func (m SettingsModel) UpdateSettings(msg tea.Msg) (SettingsModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch m.OptionsMode {
		case StyleEditorMode:
			m.updateThemeEditor(msg)
			return m, nil
		}
		switch msg.String() {
		case "j":
			if m.Cursor < m.currentOptionsCount()-1 {
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
					m.Cursor = 0
					return m, nil
				}
				m.Selected = m.Cursor
			case GameTypePickerMode:
				m.updateGamePicker()
			case StylePickerMode:
				m.updateThemePicker()
			}
		}
	}
	return m, nil
}

func (m SettingsModel) ViewSettings() string {
	view := ""
	switch m.OptionsMode {
	case OptionMenuMode:
		view = m.getOptionsMenu()
	case GameTypePickerMode:
		view = m.getGamePickerMenu()
	case StylePickerMode:
		view = m.getStylePickerMenu()
	case StyleEditorMode:
		view = m.getCustomThemeEditor()
	}
	return view
}
