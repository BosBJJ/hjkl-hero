package game

import "slices"

func (gs *GameState) TakeSnapShot(pos Position, mapLines []string) {
	mapClone := slices.Clone(mapLines)
	newSnap := SnapShot{
		PlayerSnapShot: pos,
		MapSnapShot:    mapClone,
	}
	gs.undoSnap = append(gs.undoSnap, newSnap)
	gs.redoSnap = nil
}
