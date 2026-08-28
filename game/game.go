package game

func NewGame() *GameState {
	return &GameState{
		CurrentRoom: "apartment",
		Player: Player{
			Inventory: []string{},
		},
		Flags:            make(map[string]bool),
		VisitedRooms:     make(map[string]bool),
		DialogueProgress: make(map[string]*DialogueProgress),
	}
}
