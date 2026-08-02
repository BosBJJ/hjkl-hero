package style

import "github.com/charmbracelet/lipgloss"

// Shoutout to my favorite theme https://github.com/kbraggins/duskhaven.nvim/blob/main/lua/duskhaven/palette.lua
var (
	Bg      = lipgloss.Color("#0c1021")
	BgLight = lipgloss.Color("#1a1f35")
	BgDark  = lipgloss.Color("#0a0d1a")

	Fg    = lipgloss.Color("#fdfff1")
	FgDim = lipgloss.Color("#d8d9c8")

	Orange      = lipgloss.Color("#f25e01")
	Yellow      = lipgloss.Color("#b3f63a")
	YellowLight = lipgloss.Color("#ffec40")
	Golden      = lipgloss.Color("#f0a905")
	Magenta     = lipgloss.Color("#ff0cac")

	Blue      = lipgloss.Color("#6b8ab8")
	BlueLight = lipgloss.Color("#97c7e9")
	BlueDark  = lipgloss.Color("#3c5ea9")
	BabyBlue  = lipgloss.Color("#87CEEB")

	Red     = lipgloss.Color("#e04a5f")
	RedDark = lipgloss.Color("#b8041d")
	Green   = lipgloss.Color("#55ba30")

	Peach = lipgloss.Color("#f5c6b0")
	Cream = lipgloss.Color("#f8e9c7")

	Gray       = lipgloss.Color("#a4a7a7")
	GrayDark   = lipgloss.Color("#505257")
	GrayDarker = lipgloss.Color("#35384a")

	Black = lipgloss.Color("#272822")
	Brown = lipgloss.Color("#422305")
	White = lipgloss.Color("#ffffff")
)

var ColorNames = map[string]lipgloss.Color{
	"bg":       Bg,
	"bg_light": BgLight,
	"bg_dark":  BgDark,

	"fg":     Fg,
	"fg_dim": FgDim,

	"orange":       Orange,
	"yellow":       Yellow,
	"yellow_light": YellowLight,
	"golden":       Golden,
	"magenta":      Magenta,

	"blue":       Blue,
	"blue_light": BlueLight,
	"blue_dark":  BlueDark,
	"baby_blue":  BabyBlue,

	"red":      Red,
	"red_dark": RedDark,
	"green":    Green,

	"peach": Peach,
	"cream": Cream,

	"gray":        Gray,
	"gray_dark":   GrayDark,
	"gray_darker": GrayDarker,

	"black": Black,
	"brown": Brown,
	"white": White,
}

type NameColor struct {
	Name  string
	Color lipgloss.Color
}

var Colors = []NameColor{
	{"orange", Orange},
	{"yellow", Yellow},
	{"yellow_light", YellowLight},
	{"golden", Golden},
	{"magenta", Magenta},

	{"blue", Blue},
	{"blue_light", BlueLight},
	{"blue_dark", BlueDark},
	{"baby_blue", BabyBlue},

	{"red", Red},
	{"red_dark", RedDark},
	{"green", Green},

	{"peach", Peach},
	{"cream", Cream},

	{"gray", Gray},
	{"gray_dark", GrayDark},
	{"gray_darker", GrayDarker},

	{"black", Black},
	{"brown", Brown},
	{"white", White},
}
