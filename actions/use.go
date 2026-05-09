package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Use(gs *game.GameState, cmd parser.Command) {
	directObject := cmd.DirectObject

	indirectObject := cmd.IndirectObject

	if directObject == "" {
		fmt.Println("Use what?")
		return
	}

	if indirectObject == "" {
		fmt.Println("Use it with what?")
		return
	}

	if cmd.Preposition != "with" && cmd.Preposition != "in" {
		fmt.Println("You can't do that.")
		return
	}

	objA, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		fmt.Printf("You don't see a %s.\n", directObject)
		return
	}

	objB, ok := world.FindVisibleObject(gs, indirectObject)

	if !ok {
		fmt.Printf("You don't see a %s.\n", indirectObject)
		return
	}

	// thumbdrive -> computer interaction

	if objA.ID == "thumbdrive" && objB.ID == "computer" {

		fmt.Println("This thumbdrive doesn't fit in the computer's port. You'll have to find a wayyy older computer.")

		return
	}

	fmt.Println("Nothing happens.")
}
