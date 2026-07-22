package levels

import (
	"math/rand"
	"strings"
)

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
#....^............................#
#.................................#
###################################`,
	3: `###################################
###################################
###################################
#######################...........#
#.................................#
###############...#############..##
#.............#........#..........#
#.............#........#..........#
#.............#...................#
#.............####..###...........#
#.................................#
#.................................#
#....^............................#
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

type Room struct {
	Y1 int
	Y2 int
	X1 int
	X2 int
}

// Makes a map filled in with wall tiles
func generateBlock(height, width int) [][]rune {
	blockMap := make([][]rune, height)
	for h := range blockMap {
		row := make([]rune, width)
		for w := range row {
			row[w] = '#'
		}
		blockMap[h] = row
	}
	return blockMap
}

// Marks Top/Bottom Left/Right of a room within a block
func MakeRoom(tileMap [][]rune) Room {
	maxSize := 20
	minSize := 8
	roomHeight := rand.Intn(maxSize-minSize+1) + minSize
	roomWidth := rand.Intn(maxSize-minSize+1) + minSize
	lines := len(tileMap)
	columns := len(tileMap[0])
	Y1 := rand.Intn(lines-roomHeight) + 1
	Y2 := Y1 + roomHeight
	X1 := rand.Intn(columns-roomWidth) + 1
	X2 := X1 + roomWidth
	return Room{
		Y1: Y1,
		Y2: Y2,
		X1: X1,
		X2: X2,
	}
}

// Replaces # with . for selected room
func CarveRoom(tileMap [][]rune, room Room) {
	for y := room.Y1; y < room.Y2; y++ {
		for x := room.X1; x < room.X2; x++ {
			tileMap[y][x] = '.'
		}
	}
}

// MATH - ensures at least 1 tile between rooms
func (r Room) RoomOverlap(new Room) bool {
	return r.X1-1 < new.X2 && r.X2+1 > new.X1 && r.Y1-1 < new.Y2 && r.Y2+1 > new.Y1
}

// Marks borders of each room, ensures no collision
func MakeRooms(count int, tileMap [][]rune) []Room {
	var rooms []Room
	for len(rooms) < count {
		newRoom := MakeRoom(tileMap)
		overlap := false

		for _, room := range rooms {
			if room.RoomOverlap(newRoom) {
				overlap = true
				break
			}
		}
		if !overlap {
			rooms = append(rooms, newRoom)
		}
	}
	return rooms
}

func ConnectRooms(r1, r2 Room, tileMap [][]rune) {
	wX := (r1.X1 + r1.X2) / 2
	wY := (r1.Y1 + r1.Y2) / 2
	targetX := (r2.X1 + r2.X2) / 2
	targetY := (r2.Y1 + r2.Y2) / 2
	for wX != targetX || wY != targetY {
		roll := rand.Intn(101)
		switch {
		case roll >= 70:
			if wX < targetX {
				wX++
			} else if wX > targetX {
				wX--
			}
		case roll >= 30:
			if wY < targetY {
				wY++
			} else if wY > targetY {
				wY--
			}
		default:
			dir := rand.Intn(4)
			switch dir {
			case 0:
				wX++
			case 1:
				wX--
			case 2:
				wY++
			case 3:
				wY--
			}
		}
		tileMap[wY][wX] = '.'
	}
}

func MakeMap(height, width, numOfRooms int) LevelMap {
	tileMap := generateBlock(height, width)
	rooms := MakeRooms(numOfRooms, tileMap)
	for i, room := range rooms {
		CarveRoom(tileMap, room)
		if i < len(rooms)-1 {
			ConnectRooms(rooms[i], rooms[i+1], tileMap)
		}
	}
	MakeStairs(rooms, tileMap)
	var newMap strings.Builder
	for _, row := range tileMap {
		for _, rune := range row {
			newMap.WriteString(string(rune))
		}
		newMap.WriteByte('\n')
	}

	return LevelMap(newMap.String())
}

func MakeStairs(rooms []Room, tileMap [][]rune) {
	room := rooms[rand.Intn(len(rooms))]
	x := rand.Intn(room.X2-room.X1) + room.X1
	y := rand.Intn(room.Y2-room.Y1) + room.Y1
	tileMap[y][x] = '^'
}
