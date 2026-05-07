package game

func NewGame() *GameState {
	return &GameState{
		CurrentRoom: "bedroom",
		Inventory:   []string{},
	}
}
