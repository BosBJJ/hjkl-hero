package style

import (
	"github.com/BosBJJ/hjkl-hero/internal/storage"
	"github.com/charmbracelet/lipgloss"
)

type GameStyle struct {
	WallStyle   string
	FloorStyle  string
	PlayerStyle string
}

type PanelStyle struct {
	WallColor  string
	FloorColor string
}

func MakeStyle(themeID storage.ThemeID) GameStyle {
	theme, ok := Themes[themeID]
	if !ok {
		theme = Themes[storage.DefaultTheme]
	}
	return MakeStyleFromTheme(theme)
}

func MakeStyleFromTheme(theme storage.Theme) GameStyle {
	return GameStyle{
		WallStyle: lipgloss.NewStyle().
			Foreground(GetColor(theme.WallColor)).Render(theme.WallIcon),
		FloorStyle: lipgloss.NewStyle().
			Foreground(GetColor(theme.FloorColor)).Render(theme.FloorIcon),
		PlayerStyle: lipgloss.NewStyle().
			Foreground(GetColor(theme.PlayerColor)).Render(theme.PlayerIcon),
	}
}

func GetColor(name string) lipgloss.Color {
	return ColorNames[name]
}

func MakePanelColor(themeID storage.ThemeID) PanelStyle {
	theme := Themes[themeID]
	return PanelStyle{
		WallColor:  theme.WallColor,
		FloorColor: theme.FloorColor,
	}
}
