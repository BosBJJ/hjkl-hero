package game

import (
	"fmt"
	"math/rand/v2"
	"strings"

	"github.com/BosBJJ/hjkl-hero/internal/levels"
)

func (t ItemType) String() string {
	switch t {
	case HealthPotion:
		return "Health Potion"
	case Gold:
		return "Gold"
	case LuckyFeather:
		return "Lucky Feather"
	default:
		return "UnknownItem"
	}
}

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

func (gs *GameState) GrabItem() string {
	grabbedItem := ""
	for i, item := range gs.Items {
		if item.Col == gs.Player.Column && item.Line == gs.Player.Line {
			switch item.Type {
			case Gold:
				gs.Stats.Gold += item.Amount
				grabbedItem = fmt.Sprintf("Collected %v gold", item.Amount)
			default: //Only other item is potions at the moment,add more switches later
				gs.Stats.Inventory = append(gs.Stats.Inventory, item)
				grabbedItem = "Collected 1 potion"
			}
			gs.Items = append(gs.Items[:i], gs.Items[i+1:]...)
			return grabbedItem
		}
	}
	return grabbedItem
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

func (gs *GameState) UseFeather() {
	for i, item := range gs.Stats.Inventory {
		if item.Type == LuckyFeather {
			gs.Stats.Inventory = append(gs.Stats.Inventory[:i], gs.Stats.Inventory[i+1:]...)
		}
	}
}

func (gs *GameState) HaveFeather() bool {
	for _, item := range gs.Stats.Inventory {
		if item.Type == LuckyFeather {
			return true
		}
	}
	return false
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
		obtained = gs.GrabItem()
	default:
		gs.DropGold(line, col)
		obtained = gs.GrabItem()
	}
	return obtained
}

// Add more options in the future, items instead of potion + xp
func (gs *GameState) UseMerchant(option int) (string, bool) {
	switch option {
	case 1:
		if gs.Stats.Gold < 15 {
			return "Not Enough Gold", false
		}
		gs.Stats.Gold -= 15
		gs.Stats.Inventory = append(gs.Stats.Inventory, Item{Type: HealthPotion})
		return "Bought a health potion", false
	case 2:
		if gs.Stats.Gold < 40 {
			return "Not Enough Gold", false
		}
		gs.Stats.Gold -= 40
		gs.Stats.TotalXP += 10
		gs.Stats.XPGained += 10
		return "10 XP purchased", false
	case 3:
		if gs.Stats.Gold < 150 {
			return "Not Enough Gold", false
		}
		gs.Stats.Gold -= 150
		gs.Stats.Inventory = append(gs.Stats.Inventory, Item{Type: LuckyFeather})
		return "Purchased a lucky feather!", false
	default:
		return "Good luck!", true
	}
}
