package actions

import (
	"testing"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func newObjectStateTestGame() *game.GameState {
	gs := game.NewGame()
	world.InitializeObjectStates(gs)
	return gs
}

func TestOpenAndCloseUseSessionState(t *testing.T) {
	first := newObjectStateTestGame()
	second := newObjectStateTestGame()
	command := parser.Command{DirectObject: "drawer"}

	first.ObjectStates["drawer"].Locked = true
	if got := Open(first, command); got != "It's locked" {
		t.Fatalf("opening a session-locked object returned %q", got)
	}

	if got := Open(second, command); got != "You open the drawer." {
		t.Fatalf("opening an unlocked object returned %q", got)
	}

	if !second.ObjectStates["drawer"].Open {
		t.Fatal("open did not update session object state")
	}

	if world.Objects["drawer"].Open {
		t.Fatal("open changed the static object definition")
	}

	if got := Close(second, command); got != "You close the drawer." {
		t.Fatalf("closing an open object returned %q", got)
	}

	if second.ObjectStates["drawer"].Open {
		t.Fatal("close did not update session object state")
	}
}

func TestGetMovesObjectFromContainerToInventoryInSessionState(t *testing.T) {
	gs := newObjectStateTestGame()
	gs.ObjectStates["drawer"].Open = true

	if got := Get(gs, parser.Command{DirectObject: "thumbdrive"}); got != "Taken." {
		t.Fatalf("getting the thumbdrive returned %q", got)
	}

	if len(gs.Player.Inventory) != 1 || gs.Player.Inventory[0] != "thumbdrive" {
		t.Fatalf("unexpected inventory after get: %#v", gs.Player.Inventory)
	}

	if got := gs.ObjectStates["thumbdrive"].Parent; got != "inventory" {
		t.Fatalf("thumbdrive parent is %q, want inventory", got)
	}

	if got := len(gs.ObjectStates["drawer"].Contains); got != 0 {
		t.Fatalf("drawer still has %d object(s) after get", got)
	}

	if got := world.Objects["thumbdrive"].Parent; got != "drawer" {
		t.Fatalf("get changed the thumbdrive definition parent to %q", got)
	}

	if got := len(world.Objects["drawer"].Contains); got != 1 {
		t.Fatalf("get changed the drawer definition contents: length %d", got)
	}

	if got := Look(gs, parser.Command{DirectObject: "drawer"}); got != world.Objects["drawer"].OpenDescription+"\nIt's empty" {
		t.Fatalf("look did not use session container contents: got %q", got)
	}

	gs.ObjectStates["drawer"].Open = false
	if obj, ok := world.FindVisibleObject(gs, "thumbdrive"); !ok || obj.ID != "thumbdrive" {
		t.Fatal("inventory object was not visible after its former container closed")
	}

	if got := Look(gs, parser.Command{DirectObject: "drawer"}); got != world.Objects["drawer"].ClosedDescription {
		t.Fatalf("look did not use session open state: got %q", got)
	}
}
