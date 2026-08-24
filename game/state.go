package game

type GameState struct {
	CurrentRoom          string
	Player               Player
	Dialogue             *DialogueState
	Flags                map[string]bool
	VisitedDialogueNodes map[string]map[string]bool
}

type DialogueState struct {
	DialogueID string
	NodeID     string
}
