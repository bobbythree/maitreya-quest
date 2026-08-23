package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/dialogue"
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Talk(gs *game.GameState, cmd parser.Command) string {
	if cmd.DirectObject == "" {
		return "Talk to whom?"
	}

	obj, ok := world.FindVisibleObject(gs, cmd.DirectObject)
	if !ok {
		return fmt.Sprintf("You don't see a %s to talk to.", cmd.DirectObject)
	}

	if !obj.Talkable {
		return "You'll just talk to anything huh? You get no response."
	}

	d, ok := dialogue.Dialogues[obj.DialogueID]
	if !ok {
		return "They have nothing to say."
	}

	gs.Dialogue = &game.DialogueState{
		DialogueID: d.ID,
		NodeID:     d.StartNode,
	}

	return ""
}
