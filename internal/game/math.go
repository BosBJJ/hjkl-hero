package game

func getDiff(x, y int) int {
	diff := x - y
	if diff < 0 {
		return -diff
	}
	return diff
}
