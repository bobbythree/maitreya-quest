package game

func NewGame() *GameState {
	return &GameState{
		CurrentRoom: "apartment",
		Player: Player{
			Inventory: []string{},
		},
		Flags:                make(map[string]bool),
		VisitedDialogueNodes: make(map[string]map[string]bool),
	}
}
