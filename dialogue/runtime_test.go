package dialogue

import (
	"testing"

	"github.com/bobbythree/maitreya-quest/game"
)

func TestDialogueLifecycleUsesActiveInteraction(t *testing.T) {
	gs := game.NewGame()

	if err := Start(gs, "street_man"); err != nil {
		t.Fatalf("starting dialogue: %v", err)
	}

	if got := gs.InteractionKind(); got != game.InteractionDialogue {
		t.Fatalf("interaction kind is %v, want dialogue", got)
	}

	node, ok := CurrentNode(gs)
	if !ok || node.Text == "" {
		t.Fatal("active dialogue node was unavailable")
	}

	outcome, err := SelectChoice(gs, 3)
	if err != nil {
		t.Fatalf("ending dialogue: %v", err)
	}

	if !outcome.Ended {
		t.Fatal("terminal dialogue choice did not report an ended dialogue")
	}

	if got := gs.InteractionKind(); got != game.InteractionNone {
		t.Fatalf("interaction kind after dialogue exit is %v, want none", got)
	}
}

func TestDialogueCannotReplaceWorkComputer(t *testing.T) {
	gs := game.NewGame()
	gs.BeginWorkComputer(game.WorkComputerState{Screen: "menu"})

	if err := Start(gs, "street_man"); err == nil {
		t.Fatal("started dialogue while work computer was active")
	}

	computer, ok := gs.ActiveWorkComputer()
	if !ok || computer.Screen != "menu" {
		t.Fatal("failed dialogue entry replaced the work computer")
	}

	if progress := gs.DialogueProgress["street_man"]; progress != nil {
		t.Fatal("failed dialogue entry changed dialogue progress")
	}
}
