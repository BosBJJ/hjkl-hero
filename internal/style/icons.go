package style

type NameIcon struct {
	Name   string
	Symbol string
}

var WallIcons = []NameIcon{
	{"Fully Shaded Wall", "\u2593"},
	{"3/4 Shaded Wall", "\u2592"},
	{"1/2 Shaded Wall", "\u2591"},
	{"Hash", "#"},
	{"Trees", "\u25B2"},
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
	{"Skull", "\u2620"},
	{"Arrow", "\u25B2"},
}
