package dialogue

import "github.com/bobbythree/maitreya-quest/game"

type Effect func(*game.GameState) string

type Dialogue struct {
	ID         string
	StartNode  string
	Nodes      map[string]Node
	Completion *Completion
}

type Completion struct {
	RequiredNodes []string
	OnComplete    Effect
}

type Node struct {
	Speaker string
	Text    string
	Choices []Choice
}

type Choice struct {
	Text     string
	NextNode string
}

type Outcome struct {
	Narration string
	Ended     bool
}
