package game

func (gs *GameState) DropPotion(line, col int) {
	newPotion := Item{
		Type:   HealthPotion,
		Line:   line,
		Col:    col,
		Amount: 1,
	}
	gs.Items = append(gs.Items, newPotion)
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

func (gs *GameState) UsePotion() { //Currently can overheal, maybe let it overheal but half strength or turn into shield
	for i, item := range gs.Stats.Inventory {
		if item.Type == HealthPotion {
			gs.Stats.CurrentHealth += 5
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
