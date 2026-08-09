package ui

import (
	"fmt"
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/storage"
	"github.com/BosBJJ/hjkl-hero/internal/style"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Changes length to allow cursor to adapt to different sizes of each menu
func (m SettingsModel) currentOptionsCount() int {
	switch m.OptionsMode {
	case OptionMenuMode:
		return len(m.Options)
	case GameTypePickerMode:
		return 3
	case StylePickerMode:
		return 6
	default:
		return len(m.Options)
	}
}

func (m SettingsModel) currentValuesCount() int {
	switch m.EditorMenu.PropertyCursor {
	case 0, 1, 2:
		return len(style.Colors)
	case 3:
		return len(style.WallIcons)
	case 4:
		return len(style.FloorIcons)
	case 5:
		return len(style.PlayerIcons)
	default:
		return 0
	}
}

func (m *SettingsModel) updateGamePicker() {
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
}

func (m *SettingsModel) updateThemePicker() {
	switch m.Cursor {
	case 0:
		m.ThemeSelected = storage.DefaultTheme
	case 1:
		m.ThemeSelected = storage.RedTheme
	case 2:
		m.ThemeSelected = storage.WinterTheme
	case 3:
		m.ThemeSelected = storage.CyberTheme
	case 4:
		m.ThemeSelected = storage.CustomTheme
		m.OptionsMode = StyleEditorMode
	case 5:
		m.OptionsMode = OptionMenuMode
		storage.UpdateTheme(m.DB, m.ThemeSelected)
		m.Cursor = 1
	}
}

func (m SettingsModel) MakeBorder() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderBackground(style.GetColor(m.EditorMenu.Theme.WallColor)).
		BorderForeground(style.GetColor(m.EditorMenu.Theme.FloorColor))
}

func (m SettingsModel) getOptionsMenu() string {
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
	currMode := fmt.Sprintf("Current Game Mode: %v", m.ModeSelected)
	optionBoxes = append(optionBoxes, currMode)
	menu := lipgloss.JoinVertical(lipgloss.Center, append([]string{title}, optionBoxes...)...)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top, "\n\n\n\n"+menu)
}

func (m SettingsModel) getGamePickerMenu() string {
	const Title = `
 ██████╗  █████╗ ███╗   ███╗███████╗    ███╗   ███╗ ██████╗ ██████╗ ███████╗
██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ████╗ ████║██╔═══██╗██╔══██╗██╔════╝
██║  ███╗███████║██╔████╔██║█████╗      ██╔████╔██║██║   ██║██║  ██║█████╗  
██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║╚██╔╝██║██║   ██║██║  ██║██╔══╝  
╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ██║ ╚═╝ ██║╚██████╔╝██████╔╝███████╗
 ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝    ╚═╝     ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝
                                                                            `
	title := style.MenuTitleStyle.
		Width(m.width).
		Align(lipgloss.Center).
		Render(Title) + "\n"

	optionBoxes := []string{}
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
	currMode := fmt.Sprintf("Current Game Mode: %v", m.ModeSelected)
	optionBoxes = append(optionBoxes, currMode)
	menu := lipgloss.JoinVertical(lipgloss.Center, append([]string{title}, optionBoxes...)...)

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Top, "\n\n\n\n"+menu)
}

func renderDemo(theme storage.Theme) string {
	demoMap := `###################################
###################################
###...M.............M##############
###...^###########.......##########
###..#############..###.....C.#####
###.....+#########....#$.....######
###..$..........#....##############
###......$..@..M#.Z..##############
###....^........#....C........@####
###.....+.......+...#.....Z########
###@.M.Z.C..........#....##########
###################################
###################################
`
	lines := strings.Split(demoMap, "\n")
	styleTheme := style.MakeStyleFromTheme(theme)

	var renderedDemo strings.Builder

	runeMap := make([][]rune, len(lines))

	for l, line := range lines {
		runeMap[l] = []rune(line)
	}

	for _, line := range runeMap {
		for _, tile := range line {
			switch {
			case tile == '.':
				renderedDemo.WriteString(styleTheme.FloorStyle)
			case tile == '#':
				renderedDemo.WriteString(styleTheme.WallStyle)
			case tile == '@':
				renderedDemo.WriteString(styleTheme.PlayerStyle)
			case tile == 'M':
				renderedDemo.WriteString(style.MeleerStyle.Render("M"))
			case tile == 'Z':
				renderedDemo.WriteString(style.ZanthStyle.Render("Z"))
			case tile == 'C':
				renderedDemo.WriteString(style.ChaserStyle.Render("9"))
			case tile == '^':
				renderedDemo.WriteString(style.StairStyle.Render("^"))
			case tile == '+':
				renderedDemo.WriteString(style.ChestStyle.Render("\u2317"))
			case tile == '$':
				renderedDemo.WriteString(style.VendorStyle.Render("$"))
			default:
				renderedDemo.WriteRune(tile)
			}
		}
		renderedDemo.WriteByte('\n')
	}
	return renderedDemo.String()
}

func (m SettingsModel) getStylePickerMenu() string {
	themeOptions := []string{"Default", "Red", "Wintery", "CyberPunk", "Custom", "Exit"}
	optionBoxes := []string{}
	for i, option := range themeOptions {
		if m.Cursor == i {
			optionBoxes = append(optionBoxes, style.EditorSelectedStyle.
				Align(lipgloss.Center).
				AlignVertical(lipgloss.Center).
				Width(30).
				Height(1).
				Render(option)+"\n")
		} else {
			optionBoxes = append(optionBoxes, style.EditorItemStyle.
				Align(lipgloss.Center).
				AlignVertical(lipgloss.Center).
				Width(30).
				Height(1).
				Render(option)+"\n")
		}
	}
	barSize := int(float64(m.width) * 0.25)
	styleView := renderDemo(style.Themes[m.ThemeSelected])
	leftBar := lipgloss.NewStyle().Width(barSize).Height(m.height).AlignVertical(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Left, optionBoxes...))
	center := lipgloss.NewStyle().Width(m.width - barSize - barSize).Height(m.height).AlignVertical(lipgloss.Center).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Left, styleView))
	info := "*, @, \u2E38,\u25B2 - Player\n$ - merchant\n\u233A - chest\n^ - stairs\nZ - tank\n9 - Chaser\nM - Meleer"
	rightBar := lipgloss.NewStyle().Width(barSize).Height(m.height).AlignVertical(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Center, "Legend", "--------------------", info))
	return lipgloss.JoinHorizontal(lipgloss.Top, leftBar, center, rightBar)
}

func (m SettingsModel) getCustomThemeEditor() string {
	barSize := int(float64(m.width) * 0.25)
	leftBar := m.showProperties()
	styleView := renderDemo(m.EditorMenu.Theme)
	panelHelper := ""
	if m.EditorMenu.LeftPanel {
		panelHelper = "Press Tab to switch panels. CURRENT PANEL - PROPERTIES\n\n\n"
	} else {
		panelHelper = "Press Tab to switch panels. CURRENT PANEL - VALUES\n\n\n"
	}
	panelDemo := lipgloss.NewStyle().Foreground(style.GetColor(m.EditorMenu.Theme.FloorColor)).Background(style.GetColor(m.EditorMenu.Theme.WallColor)).Render("This text is an example of how\nthe panel colors will look")
	info := "*, @, \u2E38,\u25B2 - Player\n$ - merchant\n\u233A - chest\n^ - stairs\nZ - tank\n9 - Chaser\nM - Meleer\n"
	center := lipgloss.NewStyle().Width(m.width - barSize - barSize).Height(m.height).AlignVertical(lipgloss.Center).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Center, panelHelper, panelDemo, styleView, "Legend", "--------------------", info))
	rightBar := m.showValues()
	return lipgloss.JoinHorizontal(lipgloss.Top, leftBar, center, rightBar)
}

type editorMenu struct {
	Theme          storage.Theme
	PropertyCursor int
	ValueCursor    int
	LeftPanel      bool
}

var customProperties = []string{"Wall Color", "Floor Color", "Cursor Color", "Wall Style", "Floor Style", "Cursor Style", "Exit without saving", "Save and quit"}

func (m SettingsModel) showProperties() string {
	barSize := int(float64(m.width) * 0.25)
	selectedBorder := m.MakeBorder()
	optionBoxes := []string{}
	for i, option := range customProperties {
		if m.EditorMenu.PropertyCursor == i {
			optionBoxes = append(optionBoxes, style.EditorSelectedStyle.
				Align(lipgloss.Center).
				AlignVertical(lipgloss.Center).
				Render(option)+"\n")
		} else {
			optionBoxes = append(optionBoxes, style.EditorItemStyle.
				Align(lipgloss.Center).
				AlignVertical(lipgloss.Center).
				Render(option)+"\n")
		}
	}
	if m.EditorMenu.LeftPanel {
		return selectedBorder.Width(barSize).Height(m.height - 3).AlignVertical(lipgloss.Center).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Left, optionBoxes...))
	}
	return style.EmptyBorder.Width(barSize).Height(m.height - 3).AlignVertical(lipgloss.Center).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Left, optionBoxes...))
}

func (m SettingsModel) showValues() string {
	colors := style.Colors
	selectedBorder := m.MakeBorder()
	barSize := int(float64(m.width)*0.25) - 3
	optionBoxes := []string{}
	switch m.EditorMenu.PropertyCursor {
	case 0, 1, 2:
		for i, color := range colors {
			if m.EditorMenu.ValueCursor == i {
				optionBoxes = append(optionBoxes, style.EditorSelectedStyle.
					Align(lipgloss.Center).
					AlignVertical(lipgloss.Center).
					Background(color.Color).
					Render(color.Name)+"\n")
			} else {
				optionBoxes = append(optionBoxes, lipgloss.NewStyle().
					Align(lipgloss.Center).
					AlignVertical(lipgloss.Center).
					Foreground(color.Color).
					Render(color.Name)+"\n")
			}
		}
	case 3:
		for i, icon := range style.WallIcons {
			if m.EditorMenu.ValueCursor == i {
				optionBoxes = append(optionBoxes, style.EditorPanelStyle.
					Align(lipgloss.Center).
					AlignVertical(lipgloss.Center).
					BorderForeground(style.White).
					Render(icon.Symbol)+"\n")
			} else {
				optionBoxes = append(optionBoxes, style.EditorPanelStyle.
					Align(lipgloss.Center).
					AlignVertical(lipgloss.Center).
					Render(icon.Symbol)+"\n")
			}
		}
	case 4:
		for i, icon := range style.FloorIcons {
			if m.EditorMenu.ValueCursor == i {
				optionBoxes = append(optionBoxes, style.EditorPanelStyle.
					Align(lipgloss.Center).
					AlignVertical(lipgloss.Center).
					BorderForeground(style.White).
					Render(icon.Name)+"\n")
			} else {
				optionBoxes = append(optionBoxes, style.EditorPanelStyle.
					Align(lipgloss.Center).
					AlignVertical(lipgloss.Center).
					Foreground(style.White).
					Render(icon.Name)+"\n")
			}
		}
	case 5:
		for i, icon := range style.PlayerIcons {
			if m.EditorMenu.ValueCursor == i {
				optionBoxes = append(optionBoxes, style.EditorPanelStyle.
					Align(lipgloss.Center).
					AlignVertical(lipgloss.Center).
					BorderForeground(style.White).
					Render(icon.Symbol)+"\n")
			} else {
				optionBoxes = append(optionBoxes, style.EditorPanelStyle.
					Align(lipgloss.Center).
					AlignVertical(lipgloss.Center).
					Foreground(style.White).
					Render(icon.Symbol)+"\n")
			}
		}
	}
	if !m.EditorMenu.LeftPanel {
		return selectedBorder.Width(barSize - 1).Height(m.height - 3).AlignVertical(lipgloss.Center).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Left, optionBoxes...))
	}
	return style.EmptyBorder.Width(barSize - 1).Height(m.height - 3).AlignVertical(lipgloss.Center).Align(lipgloss.Center).Render(lipgloss.JoinVertical(lipgloss.Left, optionBoxes...))
}

func (m *SettingsModel) updateThemeEditor(msg tea.KeyMsg) {
	switch msg.String() {
	case "tab":
		m.EditorMenu.LeftPanel = !m.EditorMenu.LeftPanel
		m.EditorMenu.ValueCursor = 0
	case "j":
		m.MoveDown()
	case "k":
		m.MoveUp()
	case "enter":
		if m.EditorMenu.PropertyCursor == 6 || m.EditorMenu.PropertyCursor == 7 {
			m.SelectOption()
		}
		if !m.EditorMenu.LeftPanel {
			m.SelectOption()
		}
	}
}

func (m *SettingsModel) MoveDown() {
	if m.EditorMenu.LeftPanel {
		if m.EditorMenu.PropertyCursor < len(customProperties)-1 {
			m.EditorMenu.PropertyCursor++
			return
		}
	}
	if m.EditorMenu.ValueCursor < m.currentValuesCount()-1 {
		m.EditorMenu.ValueCursor++
	}

}
func (m *SettingsModel) MoveUp() {
	if m.EditorMenu.LeftPanel {
		if m.EditorMenu.PropertyCursor > 0 {
			m.EditorMenu.PropertyCursor--
			return
		}
	}
	if m.EditorMenu.ValueCursor > 0 {
		m.EditorMenu.ValueCursor--
	}
}

func (m *SettingsModel) SelectOption() {
	switch m.EditorMenu.PropertyCursor {
	case 0: //WallColor
		m.EditorMenu.Theme.WallColor = style.Colors[m.EditorMenu.ValueCursor].Name
	case 1: //FloorColor
		m.EditorMenu.Theme.FloorColor = style.Colors[m.EditorMenu.ValueCursor].Name
	case 2: //CursorColor
		m.EditorMenu.Theme.PlayerColor = style.Colors[m.EditorMenu.ValueCursor].Name
	case 3: //WallIcon
		m.EditorMenu.Theme.WallIcon = style.WallIcons[m.EditorMenu.ValueCursor].Symbol
	case 4: //FloorIcon
		m.EditorMenu.Theme.FloorIcon = style.FloorIcons[m.EditorMenu.ValueCursor].Symbol
	case 5: //PlayerIcon
		m.EditorMenu.Theme.PlayerIcon = style.PlayerIcons[m.EditorMenu.ValueCursor].Symbol
	case 6: //Exit w/o saving
		m.OptionsMode = StylePickerMode
	case 7: //Save/Exit
		style.Themes[storage.CustomTheme] = m.EditorMenu.Theme
		storage.SaveCustomTheme(m.DB, m.EditorMenu.Theme)
		m.ThemeSelected = storage.CustomTheme
		m.OptionsMode = StylePickerMode
	}
}
