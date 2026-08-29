package game

type GameState struct {
	CurrentRoom      string
	Player           Player
	Dialogue         *DialogueState
	Flags            map[string]bool
	VisitedRooms     map[string]bool
	DialogueProgress map[string]*DialogueProgress
	WorkComputer     *WorkComputerState
}

type DialogueState struct {
	DialogueID string
	NodeID     string
}

type DialogueProgress struct {
	VisitedNodes map[string]bool
	Completed    bool
}

type WorkComputerState struct {
	Screen string
}
