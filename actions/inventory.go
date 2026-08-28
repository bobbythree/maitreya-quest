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
		result.WriteString("\n- " + item)
	}

	return result.String()
}
