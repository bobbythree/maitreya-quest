package game_test

import (
	"testing"

	"github.com/bobbythree/maitreya-quest/game"
)

func TestInteractionsAreMutuallyExclusive(t *testing.T) {
	gs := game.NewGame()

	if got := gs.InteractionKind(); got != game.InteractionNone {
		t.Fatalf("new game interaction kind is %v, want none", got)
	}

	if !gs.BeginDialogue(game.DialogueState{DialogueID: "test", NodeID: "start"}) {
		t.Fatal("could not begin dialogue with no active interaction")
	}

	if gs.BeginWorkComputer(game.WorkComputerState{Screen: "menu"}) {
		t.Fatal("began work computer while dialogue was active")
	}

	dialogueState, ok := gs.ActiveDialogue()
	if !ok || dialogueState.DialogueID != "test" {
		t.Fatal("failed work-computer entry replaced the active dialogue")
	}

	if gs.EndInteraction(game.InteractionWorkComputer) {
		t.Fatal("ended dialogue using the wrong interaction kind")
	}

	if got := gs.InteractionKind(); got != game.InteractionDialogue {
		t.Fatalf("wrong-kind exit changed interaction to %v", got)
	}

	if !gs.EndInteraction(game.InteractionDialogue) {
		t.Fatal("could not end the active dialogue")
	}

	if !gs.BeginWorkComputer(game.WorkComputerState{Screen: "menu"}) {
		t.Fatal("could not begin work computer after dialogue ended")
	}

	computerState, ok := gs.ActiveWorkComputer()
	if !ok || computerState.Screen != "menu" {
		t.Fatal("work computer did not become the active interaction")
	}

	if _, ok := gs.ActiveDialogue(); ok {
		t.Fatal("dialogue state remained active with the work computer")
	}

	if gs.BeginOldComputer(game.OldComputerState{Screen: "loading"}) {
		t.Fatal("began old computer while work computer was active")
	}

	if !gs.EndInteraction(game.InteractionWorkComputer) {
		t.Fatal("could not end the active work computer")
	}

	if !gs.BeginOldComputer(game.OldComputerState{Screen: "loading"}) {
		t.Fatal("could not begin old computer after work computer ended")
	}

	oldComputerState, ok := gs.ActiveOldComputer()
	if !ok || oldComputerState.Screen != "loading" {
		t.Fatal("old computer did not become the active interaction")
	}

	if gs.EndInteraction(game.InteractionWorkComputer) {
		t.Fatal("ended old computer using the work-computer interaction kind")
	}
}
