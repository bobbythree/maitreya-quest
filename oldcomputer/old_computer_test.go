package oldcomputer

import (
	"testing"

	"github.com/bobbythree/maitreya-quest/game"
)

func TestUseDoesNotStartInteraction(t *testing.T) {
	gs := game.NewGame()

	if got := Use(gs); got != "The computer looks like it's 100 years old. It won't turn on." {
		t.Fatalf("using old computer returned %q", got)
	}

	if got := gs.InteractionKind(); got != game.InteractionNone {
		t.Fatalf("direct use started interaction kind %v", got)
	}
}

func TestStartUsesActiveInteraction(t *testing.T) {
	gs := game.NewGame()

	if got := Start(gs); got != awakeningNarration {
		t.Fatalf("starting old computer returned %q", got)
	}

	computer, ok := gs.ActiveOldComputer()
	if !ok || computer.Screen != "awakening" {
		t.Fatal("start did not activate the old-computer awakening narration")
	}
}

func TestStartRejectsBusyOrDestroyedComputer(t *testing.T) {
	t.Run("busy", func(t *testing.T) {
		gs := game.NewGame()
		gs.BeginDialogue(game.DialogueState{DialogueID: "test", NodeID: "start"})

		if got := Start(gs); got != "You're already busy." {
			t.Fatalf("starting while busy returned %q", got)
		}
	})

	t.Run("destroyed", func(t *testing.T) {
		gs := game.NewGame()
		gs.Flags[destroyedFlag] = true

		if got := Start(gs); got != "The old computer is dead. It won't turn on again." {
			t.Fatalf("starting destroyed computer returned %q", got)
		}

		if got := Use(gs); got != "The old computer is dead. It won't turn on again." {
			t.Fatalf("using destroyed computer returned %q", got)
		}
	})
}
