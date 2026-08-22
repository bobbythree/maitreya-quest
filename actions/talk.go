package actions

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/output"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Talk(gs *game.GameState, cmd parser.Command) {
	if cmd.DirectObject == "" {
		output.Println("Talk to whom?")
		return
	}

	obj, ok := world.FindVisibleObject(gs, cmd.DirectObject)
	if !ok {
		output.Printf("You don't see a %v to talk to", obj)
	}

	if !obj.Talkable {
		output.Println("You'll just talk to anything huh? You get no response.")
	}

	output.Println(obj.Dialog)
}
