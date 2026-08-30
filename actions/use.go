package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Use(gs *game.GameState, cmd parser.Command) string {
	directObject := cmd.DirectObject

	indirectObject := cmd.IndirectObject

	if directObject == "" {
		return "Use what?"
	}

	objA, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		return fmt.Sprintf("You don't see a %s.", directObject)
	}

	// standalone 'use' case
	if indirectObject == "" {
		if objA.UseAction != nil {
			return objA.UseAction(gs)
		}

		return "Use it with what?"
	}

	if cmd.Preposition != "with" && cmd.Preposition != "in" {
		return "You can't do that."
	}

	objB, ok := world.FindVisibleObject(gs, indirectObject)

	if !ok {
		return fmt.Sprintf("You don't see a %s.", indirectObject)
	}

	// thumbdrive -> computer interaction

	if objA.ID == "thumbdrive" && objB.ID == "home_computer" {
		return "This thumbdrive doesn't fit in the computer's port. You'll have to find a wayyy older computer."
	}

	return "Nothing happens."
}
