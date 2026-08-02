package style

type NameIcon struct {
	Name   string
	Symbol string
}

var WallIcons = []NameIcon{
	{"Fully Shaded Wall", "\u2593"},
	{"Medium Shaded Wall", "\u2592"},
	{"Light Shaded Wall", "\u2591"},
	{"Hash", "#"},
	{"Trees", "\u234B"},
}

var FloorIcons = []NameIcon{
	{"Space", " "},
	{"Dot", "."},
	{"Center dot", "\u2219"},
	{"Grates", "\u253C"},
}

var PlayerIcons = []NameIcon{
	{"At", "@"},
	{"Star", "*"},
	{"Dagger", "\u2E38"},
	{"Arrow", "\u25B2"},
}
