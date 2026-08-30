package game

type GameState struct {
	CurrentRoom       string
	Player            Player
	Flags             map[string]bool
	VisitedRooms      map[string]bool
	DialogueProgress  map[string]*DialogueProgress
	ObjectStates      map[string]*ObjectState
	activeInteraction *activeInteraction
}

type ObjectState struct {
	Open     bool
	Locked   bool
	Parent   string
	Contains []string
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
	Screen   string
	ScanStep int
}
