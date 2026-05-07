package tui

import (
	"bufio"
	"fmt"
	"os"

	"github.com/bobbythree/maitreya-quest/actions"
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
)

func Run(gs *game.GameState) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Maitreya Quest.")
	fmt.Println()

	actions.Look(gs, parser.Command{})

	for {

		fmt.Print("\n> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		cmd := parser.Parse(input)

		action, ok := actions.ActionMap[cmd.Verb]

		if !ok {
			fmt.Println("I don't understand.")
			continue
		}

		action(gs, cmd)
	}
}
