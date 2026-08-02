package style

import "github.com/BosBJJ/hjkl-hero/internal/storage"

var Themes = map[storage.ThemeID]storage.Theme{
	storage.DefaultTheme: {
		WallColor:   "magenta",
		FloorColor:  "green",
		PlayerColor: "red",

		WallIcon:   "\u2593",
		FloorIcon:  "\u2219",
		PlayerIcon: "@",
	},
	storage.RedTheme: {
		WallColor:   "red_dark",
		FloorColor:  "black",
		PlayerColor: "white",

		WallIcon:   "\u2591",
		FloorIcon:  "\u2219",
		PlayerIcon: "\u2E38",
	},
	storage.WinterTheme: {
		WallColor:   "green",
		FloorColor:  "white",
		PlayerColor: "yellow_light",

		WallIcon:   "\u234B",
		FloorIcon:  " ",
		PlayerIcon: "*",
	},
	storage.CyberTheme: {
		WallColor:   "blue_dark",
		FloorColor:  "cyan",
		PlayerColor: "magenta",

		WallIcon:   "\u2593",
		FloorIcon:  "\u253C",
		PlayerIcon: "\u25B2",
	},
}
