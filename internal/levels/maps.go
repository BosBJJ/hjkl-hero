package levels

func GetLevel(level int) (LevelMap, bool) {
	m, ok := Maps[level]
	return m, ok
}

type LevelMap string

var Maps = map[int]LevelMap{
	1: `#################################
#.................................#
#...............@.................#
#.................................#
#################################`,
}
