package world

import (
	"testing"

	"github.com/bobbythree/maitreya-quest/game"
)

func newTestGameState() *game.GameState {
	gs := game.NewGame()
	InitializeObjectStates(gs)
	return gs
}

func TestInitializeObjectStatesCreatesIndependentState(t *testing.T) {
	first := newTestGameState()
	second := newTestGameState()

	first.ObjectStates["drawer"].Open = true
	first.ObjectStates["drawer"].Contains[0] = "changed"

	if second.ObjectStates["drawer"].Open {
		t.Fatal("opening an object in one session changed another session")
	}

	if got := second.ObjectStates["drawer"].Contains[0]; got != "thumbdrive" {
		t.Fatalf("changing container contents in one session changed another: got %q", got)
	}

	drawerDefinition := Objects["drawer"]
	if drawerDefinition.Open {
		t.Fatal("opening an object in a session changed the object definition")
	}

	if got := drawerDefinition.Contains[0]; got != "thumbdrive" {
		t.Fatalf("changing session container contents changed the object definition: got %q", got)
	}
}

func TestFindVisibleObjectUsesSessionContainerState(t *testing.T) {
	first := newTestGameState()
	second := newTestGameState()

	if _, ok := FindVisibleObject(first, "thumbdrive"); ok {
		t.Fatal("found an object inside a closed nested container")
	}

	first.ObjectStates["drawer"].Open = true

	if obj, ok := FindVisibleObject(first, "thumbdrive"); !ok || obj.ID != "thumbdrive" {
		t.Fatal("did not find an object inside an open nested container")
	}

	if _, ok := FindVisibleObject(second, "thumbdrive"); ok {
		t.Fatal("container state leaked into another session")
	}
}
