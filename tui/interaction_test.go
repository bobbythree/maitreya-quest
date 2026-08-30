package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"

	"github.com/bobbythree/maitreya-quest/dialogue"
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/workcomputer"
)

func TestUpdateRoutesByActiveInteractionKind(t *testing.T) {
	t.Run("dialogue", func(t *testing.T) {
		gs := game.NewGame()
		if err := dialogue.Start(gs, "street_man"); err != nil {
			t.Fatalf("starting dialogue: %v", err)
		}

		model := Model{gameState: gs}
		updated, _ := model.Update(tea.KeyPressMsg(tea.Key{Code: tea.KeyDown}))
		got := updated.(Model)

		if got.dialogueUI.cursor != 1 {
			t.Fatalf("dialogue cursor is %d, want 1", got.dialogueUI.cursor)
		}

		if got.workComputerUI.cursor != 0 {
			t.Fatalf("work-computer cursor unexpectedly changed to %d", got.workComputerUI.cursor)
		}
	})

	t.Run("work computer", func(t *testing.T) {
		gs := game.NewGame()
		workcomputer.Start(gs)

		model := Model{gameState: gs}
		updated, _ := model.Update(tea.KeyPressMsg(tea.Key{Code: tea.KeyDown}))
		got := updated.(Model)

		if got.workComputerUI.cursor != 1 {
			t.Fatalf("work-computer cursor is %d, want 1", got.workComputerUI.cursor)
		}

		if got.dialogueUI.cursor != 0 {
			t.Fatalf("dialogue cursor unexpectedly changed to %d", got.dialogueUI.cursor)
		}
	})
}
