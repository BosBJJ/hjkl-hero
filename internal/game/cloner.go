package game

import "slices"

func(gs *GameState) TakeSnapShot(pos Position, mapLines []string) {
	gs.SnapShot.Line = pos.Line
	gs.SnapShot.Column = pos.Column
	gs.MapInfo.MapSnapShot = slices.Clone(mapLines)
}
