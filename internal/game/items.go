package game

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/levels"
)

func (gs *GameState) DropItem(line, col, amount int, itemType ItemType) {
	newItem := Item{
		Type:   itemType,
		Line:   line,
		Col:    col,
		Amount: amount,
	}
	gs.Items = append(gs.Items, newItem)
}

func (gs *GameState) DropPotion(line, col int) {
	gs.DropItem(line, col, 1, HealthPotion)
}

func (gs *GameState) DropGold(line, col int) int {
	goldValue := rand.IntN(15)
	if goldValue == 0 {
		goldValue = 1
	}
	gs.DropItem(line, col, goldValue, Gold)
	return goldValue
}

func (gs *GameState) GrabItem() {
	for i, item := range gs.Items {
		if item.Col == gs.Player.Column && item.Line == gs.Player.Line {
			switch item.Type {
			case Gold:
				gs.Stats.Gold += item.Amount
			default:
				gs.Stats.Inventory = append(gs.Stats.Inventory, item)
			}
			gs.Items = append(gs.Items[:i], gs.Items[i+1:]...)
			return
		}
	}
}

func (gs *GameState) UsePotion() {
	for i, item := range gs.Stats.Inventory {
		if item.Type == HealthPotion {
			if gs.Stats.CurrentHealth >= gs.Stats.MaxHealth {
				gs.Stats.CurrentHealth += 3
			} else {
				gs.Stats.CurrentHealth += 6
			}
			gs.Stats.Inventory = append(gs.Stats.Inventory[:i], gs.Stats.Inventory[i+1:]...)
			return
		}
	}
}

func (gs GameState) ItemAt(line, col int) (Item, bool) {
	for _, item := range gs.Items {
		if item.Line == line && item.Col == col {
			return item, true
		}
	}
	return Item{}, false
}

func (gs *GameState) OpenChest(line, col int) string {
	obtained := ""
	mapLines := ToLines(*gs)
	levels.ReplaceTile(mapLines, line, col, '.')
	gs.MapInfo.LevelMap = levels.LevelMap(strings.Join(mapLines, "\n"))
	roll := rand.IntN(4)
	switch roll {
	case 4, 3:
		gs.DropPotion(line, col)
		gs.GrabItem()
		obtained = "Obtained health potion from chest"
	default:
		gold := gs.DropGold(line, col)
		gs.GrabItem()
		obtained = fmt.Sprintf("Obtained %v gold from chest", gold)
	}
	return obtained
}
