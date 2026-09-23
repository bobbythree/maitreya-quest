package actions

import (
	"strings"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
)

func Inventory(gs *game.GameState, cmd parser.Command) string {
	if len(gs.Player.Inventory) == 0 {
		return "You are carrying nothing."
	}

	result := strings.Builder{}

	result.WriteString("You are carrying:")

	for _, item := range gs.Player.Inventory {
		label := item
		if item == "adapter" && gs.ObjectStates["thumbdrive"].Parent == "adapter" {
			label = "adapter (with thumbdrive)"
		}
		result.WriteString("\n- " + label)
	}

	return result.String()
}
