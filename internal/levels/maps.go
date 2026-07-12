package levels

func GetLevel(level int) (LevelMap, bool) {
	m, ok := Maps[level]
	return m, ok
}
func GetAnswer(level int) LevelMap {
	m, _ := AnswerMap[level]
	return m
}

type LevelMap string

var Maps = map[int]LevelMap{
	1: `
 This level is just to show basic functions
 such as x which will delete below your cursor
 r which will replace under the cursor.
 To start, go to the typos and fixx themm by pressinh x orr r
 Goof luCkK!`,
	2: `###################################
#.................................#
#.................................#
###############...#############..##
#.............#........#..........#
#.............#........#..........#
#.............#...................#
#.............####..###...........#
#.................................#
#.................................#
#.................................#
#.................................#
#.................................#
###################################`,
}

var AnswerMap = map[int]LevelMap{
	1: `
 This level is just to show basic functions
 such as x which will delete below your cursor
 r which will replace under the cursor.
 To start, go to the typos and fix them by pressing x or r
 Good luck!`,
}
