package levels

import "strings"

func GetLevel(level int) (LevelMap, bool) {
	m, ok := Maps[level]
	return m, ok
}

type LevelMap string

var Maps = map[int]LevelMap{
	1: `#################################
#.................................#
#.....................p...........#
#.................................#
#################################`,
	2: `This level is just to show basic functions
	such as x which will delete below your cursor
	r which will replace under the cursor
	to start, go to the typos and fixx themm by pressinh x orr r
	Goof luCkK!`,
}

func ToLines(sMap LevelMap) []string {
	return strings.Split(string(sMap), "\n")
}

func DeleteAt(currLine []string, line, col int) []string {
	if line < 0 || line > len(currLine) {
		return currLine
	}
	runes := []rune(currLine[line])
	if col < 0 || col > len(runes) {
		return currLine
	}
	lastIndex := len(runes) - 1
	if runes[0] == '#' && runes[lastIndex] == '#' {
		runes[col] = '.'
	} else {
		runes = append(runes[:col], runes[col+1:]...)
	}
	currLine[line] = string(runes)
	return currLine
}

func ToText(lines []string) string {
	return strings.Join(lines, "\n")
}
