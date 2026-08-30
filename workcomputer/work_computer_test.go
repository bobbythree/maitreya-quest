package workcomputer

import (
	"testing"

	"github.com/bobbythree/maitreya-quest/game"
)

func TestStartUsesActiveInteraction(t *testing.T) {
	gs := game.NewGame()

	if got := Start(gs); got != "" {
		t.Fatalf("starting work computer returned %q", got)
	}

	computer, ok := gs.ActiveWorkComputer()
	if !ok || computer.Screen != "menu" {
		t.Fatal("start did not activate the work-computer menu")
	}
}

func TestStartDoesNotReplaceDialogue(t *testing.T) {
	gs := game.NewGame()
	gs.BeginDialogue(game.DialogueState{DialogueID: "test", NodeID: "start"})

	if got := Start(gs); got != "You're already busy." {
		t.Fatalf("starting work computer during dialogue returned %q", got)
	}

	if got := gs.InteractionKind(); got != game.InteractionDialogue {
		t.Fatalf("work computer replaced dialogue with interaction kind %v", got)
	}

	if _, ok := gs.ActiveWorkComputer(); ok {
		t.Fatal("work computer became active alongside dialogue")
	}
}
