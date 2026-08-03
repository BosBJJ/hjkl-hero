package levels

// MATH - ensures at least 1 tile between rooms
func (r Room) RoomOverlap(new Room) bool {
	return r.X1-1 < new.X2 && r.X2+1 > new.X1 && r.Y1-1 < new.Y2 && r.Y2+1 > new.Y1
}

func (r Room) GetRoomSize() (height, width int) {
	height = r.Y2 - r.Y1
	width = r.X2 - r.X1
	return height, width
}

func (r Room) GetCenter() (y, x int) {
	y = (r.Y1 + r.Y2) / 2
	x = (r.X1 + r.X2) / 2
	return y, x
}
