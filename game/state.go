package game

type GameState struct {
	CurrentRoom string
	Player      Player
	Dialogue    *DialogueState
}

type DialogueState struct {
	DialogueID string
	NodeID     string
}
