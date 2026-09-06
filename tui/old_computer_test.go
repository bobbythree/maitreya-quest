package tui

import (
	"strings"
	"testing"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/oldcomputer"
	"github.com/bobbythree/maitreya-quest/world"
)

func oldComputerTestModel(gs *game.GameState) Model {
	return Model{
		gameState:     gs,
		oldComputerUI: newOldComputerUIState(),
	}
}

func TestPromptStartsOldComputerTimer(t *testing.T) {
	gs := game.NewGame()
	world.InitializeObjectStates(gs)
	gs.CurrentRoom = "room_108"
	gs.Player.Inventory = append(gs.Player.Inventory, "thumbdrive")
	gs.ObjectStates["thumbdrive"].Parent = "inventory"

	promptInput := textinput.New()
	promptInput.SetValue("use thumbdrive with computer")
	model := oldComputerTestModel(gs)
	model.input = promptInput

	updated, cmd := model.updatePrompt(tea.KeyPressMsg(tea.Key{Code: tea.KeyEnter}))
	model = updated.(Model)

	if cmd == nil {
		t.Fatal("starting old computer did not schedule the loading tick")
	}

	if model.oldComputerUI.generation == 0 {
		t.Fatal("starting old computer did not initialize a timer generation")
	}

	if got := gs.InteractionKind(); got != game.InteractionOldComputer {
		t.Fatalf("prompt started interaction kind %v", got)
	}
	if got := model.history[len(model.history)-1]; !strings.Contains(got, "Amazingly, and beyone all reason") {
		t.Fatalf("prompt history did not include awakening narration: %q", got)
	}
}

func TestOldComputerLoadingPasswordAndShutdownFlow(t *testing.T) {
	gs := game.NewGame()
	oldcomputer.Start(gs)
	model := oldComputerTestModel(gs)
	if cmd := model.beginOldComputerUI(); cmd == nil {
		t.Fatal("beginning old computer UI did not schedule a loading tick")
	}
	generation := model.oldComputerUI.generation

	if view := model.oldComputerView(); view != "" {
		t.Fatalf("awakening rendered computer UI %q", view)
	}

	updated, cmd := model.updateOldComputerStep(oldComputerStepMsg{generation: generation - 1})
	model = updated.(Model)
	computer, _ := gs.ActiveOldComputer()
	if computer.Screen != "awakening" || cmd != nil {
		t.Fatal("stale tick advanced the awakening narration")
	}

	updated, cmd = model.updateOldComputerStep(oldComputerStepMsg{generation: generation})
	model = updated.(Model)
	computer, _ = gs.ActiveOldComputer()
	if computer.Screen != "loading" {
		t.Fatalf("awakening tick advanced to %q, want loading", computer.Screen)
	}
	if cmd == nil {
		t.Fatal("loading screen did not schedule its timer")
	}
	if view := model.oldComputerView(); !strings.Contains(view, "LOADING DATA FROM EXTERNAL DISK...") {
		t.Fatalf("loading view was %q", view)
	}

	updated, cmd = model.updateOldComputerStep(oldComputerStepMsg{generation: generation})
	model = updated.(Model)
	computer, _ = gs.ActiveOldComputer()
	if computer.Screen != "password" {
		t.Fatalf("loading tick advanced to %q, want password", computer.Screen)
	}
	if cmd == nil {
		t.Fatal("password screen did not focus its input")
	}

	updated, _ = model.Update(tea.KeyPressMsg(tea.Key{Code: 'x', Text: "x"}))
	model = updated.(Model)
	if got := model.oldComputerUI.passwordInput.Value(); got != "x" {
		t.Fatalf("password input is %q, want x", got)
	}
	passwordView := model.oldComputerView()
	if !strings.Contains(passwordView, "*") || strings.Contains(passwordView, "x") {
		t.Fatalf("password view did not mask input: %q", passwordView)
	}

	updated, cmd = model.Update(tea.KeyPressMsg(tea.Key{Code: tea.KeyEnter}))
	model = updated.(Model)
	computer, _ = gs.ActiveOldComputer()
	if computer.Screen != "shutdown" {
		t.Fatalf("password submission advanced to %q, want shutdown", computer.Screen)
	}
	if !gs.Flags["old_computer_destroyed"] {
		t.Fatal("password submission did not persist destruction")
	}
	if cmd == nil {
		t.Fatal("password submission did not schedule shutdown")
	}

	shutdownGeneration := model.oldComputerUI.generation
	if view := model.oldComputerView(); !strings.Contains(view, "SHUTTING DOWN...") {
		t.Fatalf("shutdown view was %q", view)
	}

	updated, cmd = model.updateOldComputerStep(oldComputerStepMsg{generation: generation})
	model = updated.(Model)
	if gs.InteractionKind() != game.InteractionOldComputer || cmd != nil {
		t.Fatal("stale loading tick ended the shutdown screen")
	}

	updated, cmd = model.updateOldComputerStep(oldComputerStepMsg{generation: shutdownGeneration})
	model = updated.(Model)
	if got := gs.InteractionKind(); got != game.InteractionNone {
		t.Fatalf("shutdown ended with interaction kind %v", got)
	}
	if cmd != nil {
		t.Fatal("completed shutdown scheduled another command")
	}
	if got := model.oldComputerUI.passwordInput.Value(); got != "" {
		t.Fatalf("shutdown retained password value %q", got)
	}
	if len(model.history) != 1 || model.history[0] != "The old computer sputters and dies. It won't turn on again." {
		t.Fatalf("shutdown history is %#v", model.history)
	}
}

func TestOldComputerAcceptsEmptyPassword(t *testing.T) {
	gs := game.NewGame()
	oldcomputer.Start(gs)
	model := oldComputerTestModel(gs)
	model.beginOldComputerUI()

	updated, _ := model.updateOldComputerStep(oldComputerStepMsg{generation: model.oldComputerUI.generation})
	model = updated.(Model)
	updated, _ = model.updateOldComputerStep(oldComputerStepMsg{generation: model.oldComputerUI.generation})
	model = updated.(Model)
	updated, cmd := model.updateOldComputer(tea.KeyPressMsg(tea.Key{Code: tea.KeyEnter}))
	model = updated.(Model)

	computer, _ := gs.ActiveOldComputer()
	if computer.Screen != "shutdown" || !gs.Flags["old_computer_destroyed"] || cmd == nil {
		t.Fatal("empty password did not start persistent shutdown")
	}
}
