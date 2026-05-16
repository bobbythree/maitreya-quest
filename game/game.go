package game

func NewGame() *GameState {
	return &GameState{
		CurrentRoom: "apartment",
		Player: Player{
			Inventory: []string{},
		},
	}
}
