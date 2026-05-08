package game

func NewGame() *GameState {
	return &GameState{
		CurrentRoom: "bedroom",
		Player: Player{
			Inventory: []string{},
		},
	}
}
