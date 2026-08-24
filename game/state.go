package game

type GameState struct {
	CurrentRoom      string
	Player           Player
	Dialogue         *DialogueState
	Flags            map[string]bool
	DialogueProgress map[string]*DialogueProgress
}

type DialogueState struct {
	DialogueID string
	NodeID     string
}

type DialogueProgress struct {
	VisitedNodes map[string]bool
	Completed    bool
}
