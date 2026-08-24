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

func GetLevelsCount() int {
	return len(Maps) + 1
}

type LevelMap string

// When adding maps, go to game_helpers/DisplayTutorial and add description
var Maps = map[int]LevelMap{
	1: `
 Use H J K L to move around, your goal is to get to
 the fourth row and get to the symbol "*" in the middle and then
 use "r" which will replace under the cursor.
 Use "l" to get to the * and replace it with a "#"
 Type ":" to enter command mode and then "w", press enter to use it, this normally
 just saves in the editor but we use it to check if we're done`,
	2: `
 It was pretty annoying to just hold J to get to
 the bottom row so now
 we will be introducing the
 ability to use numbers to multiply
 how many times your next action is used
 This line is already 5 rows lower than the start
 To get to this line quickly you can press "6" and then "j", delete the "*" by pressing "x"`,
	3: `
 Now we're gonna introduce w/b/e, use "w" to jump to * and delete it
 then *come down to this line (you should be inside here) and use "b" to jump to come
 now we're at this line and need to delete the unnecessary* again, press "e" to jump to the end of words
 with vim motions using SHIFT before actions will cause them do act differently,
 capital "W" will skip to the complete new word and ignore punctuation so skip,this,phrase *with "W"
 Same *applies,to,using capital "B", it takes you to the beginning of a word ignoring punctuation
 and with "E" it,also,does,the,same,thing,but,goes,to,the,end*`,
	4: `
 There's a very easy way to jump to the beginning and end of a row
 to get to the very end of a row all you have to do is press "$"*
 *and to jump to the first character/rune just press "0"*
 use these to erase the different places you see a "*"`,
	5: `
 Now we're going to learn how to use delete
 The easiest way to delete a line is by hitting "d" to enter delete mode and then "d" to cut the line
 ********************************************************
 practice deleting that line, and then we can undo it by pressing "u"
 if we want to redo that action we can press "CTRL + r"
 You can also use "d" to enter delete mode and then select a direction
 When you're done playing around just undo everything and save the map with the asterisk line deleted`,
	6: `
 Yanking (copy) and pasting is something
 You'll be probably doing a lot of, theres a lot of different
 ways of using it, but the most common way I use it is by
 pressing "y" to enter yank mode and then "y" to yank the current
 line, to try that out, copy the line below by pressing "y" and "y" and then paste it with "p"
 ********************************************************
 Then try to copy ***** and put it inside the curly braces {}, to do this have the cursor
 on the first asterisk and then press "y" to enter yank mode and "w" to copy the word, go to the { and press "p"`,
	7: `
 Typing is very simple, we'll start by demonstrating what "i" and "a" do,
 you'll notice that when we do type the cursor will always be highlighting
 to the right of where we're typing, when we use "i" our cursor won't move
 and we'll begin typing to the left of our current location, when we use "a" our
 cursor moves one spot to the right and then we type to the left of that new position
 place a "*" in () and (), the first one stop on the "(" and press "a", the second one stop
 on the ")" and press "i" when you're done typing press "ESCAPE"
 If you need to create a new line and type in it you can press "o" practice that by adding
 the phrase "I love Vim" below.`,
	8: `
 Vim Motions has modifiers you can use, one of them is after you use a command you can press "i" to
 tell it do this command "IN" something, we'll practice yanking, deleting, and changing words.
 Our first task is to delete **** and we'll do that by going anywhere within the word and
 pressing "d" for delete mode and then either "w" to go from the cursor to end or "i"+"w" to
 target the entire word. Next we'll yank "y"+"i"+"w" and paste "p" the word "Rocks" by placing the
 cursor on the dash "Vim-" The last task is to use change mode, an easy way to edit an entire word is by 
 using "c"+"i"+"w", try changing "Vim is lame" to "Vim is awesome", remember when you're done typing
 to press "ESCAPE"`,
	9: `
 That's all for the current tutorial maps, more to be made in the future
 Now if you want to use the new keys you learned a lot in a game setting you can use ":wq" and
 the tutorials game mode will start, it is two levels on easy mode.
 The enemies take one action per one of your actions,
 if you are about to be hit you can use a number + direction and
 skip in that direction freely, use ":help" to see all of the commands available
 If you enjoy it, go to OPTIONS and switch GAME MODE to ROGUE, it is
 harder and has procedurally generated maps`,

	10: `###################################
#.................................#
#.................................#
###############...#############..##
#.............#........#..........#
#.............#........#..........#
#.............#...................#
#.............####..###...........#
#.........+.......................#
#.................................#
#.............+...................#
#....^............................#
#.................................#
###################################`,

	11: `###################################
###################################
###################################
#######################...........#
#.................................#
###############...#############..##
#.............#........#..........#
#.............#...$....#..........#
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
 Use H J K L to move around, your goal is to get to
 the fourth row and get to the symbol "*" in the middle and then
 use "r" which will replace under the cursor.
 Use "l" to get to the # and replace it with a "#"
 Type ":" to enter command mode and then "w", press enter to use it, this normally
 just saves in the editor but we use it to check if we're done`,
	2: `
 It was pretty annoying to just hold J to get to
 the bottom row so now
 we will be introducing the
 ability to use numbers to multiply
 how many times your next action is used
 This line is already 5 rows lower than the start
 To get to this line quickly you can press "6" and then "j", delete the "" by pressing "x"`,
	3: `
 Now we're gonna introduce w/b/e, use "w" to jump to  and delete it
 then come down to this line (you should be inside here) and use "b" to jump to come
 now we're at this line and need to delete the unnecessary again, press "e" to jump to the end of words
 with vim motions using SHIFT before actions will cause them do act differently,
 capital "W" will skip to the complete new word and ignore punctuation so skip,this,phrase with "W"
 Same applies,to,using capital "B", it takes you to the beginning of a word ignoring punctuation
 and with "E" it,also,does,the,same,thing,but,goes,to,the,end`,
	4: `
 There's a very easy way to jump to the beginning and end of a row
 to get to the very end of a row all you have to do is press "$"
 and to jump to the first character/rune just press "0"
 use these to erase the different places you see a ""`,
	5: `
 Now we're going to learn how to use delete
 The easiest way to delete a line is by hitting "d" to enter delete mode and then "d" to cut the line
 practice deleting that line, and then we can undo it by pressing "u"
 if we want to redo that action we can press "CTRL + r"
 You can also use "d" to enter delete mode and then select a direction
 When you're done playing around just undo everything and save the map with the asterisk line deleted`,
	6: `
 Yanking (copy) and pasting is something
 You'll be probably doing a lot of, theres a lot of different
 ways of using it, but the most common way I use it is by
 pressing "y" to enter yank mode and then "y" to yank the current
 line, to try that out, copy the line below by pressing "y" and "y" and then paste it with "p"
 ********************************************************
 ********************************************************
 Then try to copy ***** and put it inside the curly braces {***** }, to do this have the cursor
 on the first asterisk and then press "y" to enter yank mode and "w" to copy the word, go to the { and press "p"`,
	7: `
 Typing is very simple, we'll start by demonstrating what "i" and "a" do,
 you'll notice that when we do type the cursor will always be highlighting
 to the right of where we're typing, when we use "i" our cursor won't move
 and we'll begin typing to the left of our current location, when we use "a" our
 cursor moves one spot to the right and then we type to the left of that new position
 place a "*" in () and (), the first one stop on the "(" and press "a", the second one stop
 on the ")" and press "i" when you're done typing press "ESCAPE"
 If you need to create a new line and type in it you can press "o" practice that by adding
 the phrase "I love Vim" below.
 I love Vim`,
	8: `
 Vim Motions has modifiers you can use, one of them is after you use a command you can press "i" to
 tell it do this command "IN" something, we'll practice yanking, deleting, and changing words.
 Our first task is to delete  and we'll do that by going anywhere within the word and
 pressing "d" for delete mode and then either "w" to go from the cursor to end or "i"+"w" to
 target the entire word. Next we'll yank "y"+"i"+"w" and paste "p" the word "Rocks" by placing the
 cursor on the dash "Vim-Rocks" The last task is to use change mode, an easy way to edit an entire word is by 
 using "c"+"i"+"w", try changing "Vim is awesome" to "Vim is awesome", remember when you're done typing
 to press "ESCAPE"`,
	9: `
 That's all for the current tutorial maps, more to be made in the future
 Now if you want to use the new keys you learned a lot in a game setting you can use ":wq" and
 the tutorials game mode will start, it is two levels on easy mode.
 The enemies take one action per one of your actions,
 if you are about to be hit you can use a number + direction and
 skip in that direction freely, use ":help" to see all of the commands available
 If you enjoy it, go to OPTIONS and switch GAME MODE to ROGUE, it is
 harder and has procedurally generated maps`,
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
	maxSize := 18
	minSize := 8
	roomHeight := rand.Intn(maxSize-minSize+1) + minSize
	roomWidth := rand.Intn(maxSize-minSize+1) + minSize
	lines := len(tileMap)
	columns := len(tileMap[0])
	Y1 := rand.Intn(lines-roomHeight-1) + 1
	Y2 := Y1 + roomHeight
	X1 := rand.Intn(columns-roomWidth-1) + 1
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

// Marks borders of each room, ensures no collision
func MakeRooms(count int, tileMap [][]rune) []Room {
	var rooms []Room
	attempts := 0
	maxAttempts := count * 100
	for len(rooms) < count && attempts < maxAttempts {
		attempts++
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

// Uses "Drunkards walk", carves out one tile, then either continues towards target or stumbles into a random direction
func ConnectRooms(r1, r2 Room, tileMap [][]rune) {
	wY, wX := r1.GetCenter()
	targetY, targetX := r2.GetCenter()
	for wX != targetX || wY != targetY {
		roll := rand.Intn(101)
		switch {
		case roll >= 80:
			if wX < targetX {
				wX++
			} else if wX > targetX {
				wX--
			}
		case roll >= 40:
			if wY < targetY {
				wY++
			} else if wY > targetY {
				wY--
			}
		default:
			dir := rand.Intn(4)
			switch dir {
			case 0:
				if wX < len(tileMap[0])-1 {
					wX++
				}
			case 1:
				if wX > 0 {
					wX--
				}
			case 2:
				if wY < len(tileMap)-1 {
					wY++
				}
			case 3:
				if wY > 0 {
					wY--
				}
			}
		}
		tileMap[wY][wX] = '.'
	}
}

// numOfRooms = 15 + level, Every 3 levels == 18,21,24,27,30
func MakeMap(height, width, numOfRooms int) (LevelMap, bool) {
	tileMap := generateBlock(height, width)
	rooms := MakeRooms(numOfRooms, tileMap)
	for i, room := range rooms {
		CarveRoom(tileMap, room)
		if i < len(rooms)-1 {
			ConnectRooms(rooms[i], rooms[i+1], tileMap)
		}
	}
	MakeStairs(rooms, tileMap)
	specialEvent := MakeVendor(rooms, tileMap)
	MakeChest(rooms, tileMap)
	if numOfRooms > 20 {
		MakeChest(rooms, tileMap)
	}
	var newMap strings.Builder
	for _, row := range tileMap {
		for _, rune := range row {
			newMap.WriteString(string(rune))
		}
		newMap.WriteByte('\n')
	}

	return LevelMap(newMap.String()), specialEvent
}

func MakeStairs(rooms []Room, tileMap [][]rune) {
	room := rooms[rand.Intn(len(rooms))]
	x := rand.Intn(room.X2-room.X1) + room.X1
	y := rand.Intn(room.Y2-room.Y1) + room.Y1
	tileMap[y][x] = '^'
}

func MakeChest(rooms []Room, tileMap [][]rune) {
	roll := rand.Intn(100)
	if roll > 40 {
		room := rooms[rand.Intn(len(rooms))]
		x := rand.Intn(room.X2-room.X1) + room.X1
		y := rand.Intn(room.Y2-room.Y1) + room.Y1
		tileMap[y][x] = '+'
	}
}

func MakeVendor(rooms []Room, tileMap [][]rune) bool {
	roll := rand.Intn(100)
	if roll < 30 {
		for _, room := range rooms {
			height, width := room.GetRoomSize()
			if height > 10 && width > 10 {
				y, x := room.GetCenter()
				tileMap[y][x] = '$'
				return true
			}
		}
	}
	return false
}

func ReplaceTile(mapLines []string, line, col int, input rune) {
	runes := []rune(mapLines[line])
	runes[col] = input
	mapLines[line] = string(runes)
}
