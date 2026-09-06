package actions

import (
	"testing"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func oldComputerUseTestGame() *game.GameState {
	gs := game.NewGame()
	world.InitializeObjectStates(gs)
	gs.CurrentRoom = "room_108"
	gs.Player.Inventory = append(gs.Player.Inventory, "thumbdrive")
	gs.ObjectStates["thumbdrive"].Parent = "inventory"

	return gs
}

func TestUseOldComputerDirectlyDoesNotStartInteraction(t *testing.T) {
	gs := oldComputerUseTestGame()

	if got := Use(gs, parser.Parse("use computer")); got != "The computer looks like it's 100 years old. It won't turn on." {
		t.Fatalf("direct use returned %q", got)
	}

	if got := gs.InteractionKind(); got != game.InteractionNone {
		t.Fatalf("direct use started interaction kind %v", got)
	}
}

func TestUseThumbdriveWithOldComputerStartsInteraction(t *testing.T) {
	const narration = "Amazingly, and beyone all reason, the computer suddenly comes to life!!!"
	tests := []string{
		"use thumbdrive with computer",
		"use thumbdrive on computer",
		"use thumbdrive in computer",
		"use computer with thumbdrive",
		"use computer on thumbdrive",
		"use computer in thumbdrive",
	}

	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			gs := oldComputerUseTestGame()

			if got := Use(gs, parser.Parse(input)); got != narration {
				t.Fatalf("use returned %q", got)
			}

			computer, ok := gs.ActiveOldComputer()
			if !ok || computer.Screen != "awakening" {
				t.Fatal("combination did not start the old computer")
			}
		})
	}
}

func TestUseThumbdriveCannotRestartDestroyedOldComputer(t *testing.T) {
	gs := oldComputerUseTestGame()
	gs.Flags["old_computer_destroyed"] = true

	got := Use(gs, parser.Parse("use thumbdrive with computer"))
	if got != "The old computer is dead. It won't turn on again." {
		t.Fatalf("using thumbdrive with destroyed computer returned %q", got)
	}

	if got := gs.InteractionKind(); got != game.InteractionNone {
		t.Fatalf("destroyed computer restarted interaction kind %v", got)
	}
}

func TestUseThumbdriveWithHomeComputerKeepsExistingResponse(t *testing.T) {
	want := "This thumbdrive doesn't fit in the computer's port. You'll have to find a wayyy older computer."
	for _, preposition := range []string{"with", "in", "on"} {
		t.Run(preposition, func(t *testing.T) {
			gs := game.NewGame()
			world.InitializeObjectStates(gs)
			gs.Player.Inventory = append(gs.Player.Inventory, "thumbdrive")
			gs.ObjectStates["thumbdrive"].Parent = "inventory"

			got := Use(gs, parser.Parse("use thumbdrive "+preposition+" computer"))
			if got != want {
				t.Fatalf("home-computer use returned %q, want %q", got, want)
			}
		})
	}
}

func TestUseThumbdriveWithWorkComputerExplainsMismatch(t *testing.T) {
	want := "this computer is pretty old, but it's not THAT old!"
	for _, preposition := range []string{"with", "in", "on"} {
		t.Run(preposition, func(t *testing.T) {
			gs := game.NewGame()
			world.InitializeObjectStates(gs)
			gs.CurrentRoom = "work_breakroom"
			gs.Player.Inventory = append(gs.Player.Inventory, "thumbdrive")
			gs.ObjectStates["thumbdrive"].Parent = "inventory"

			got := Use(gs, parser.Parse("use thumbdrive "+preposition+" computer"))
			if got != want {
				t.Fatalf("work-computer use returned %q, want %q", got, want)
			}
		})
	}
}
