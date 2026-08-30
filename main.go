package main

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/tui"
	"github.com/bobbythree/maitreya-quest/world"
)

func main() {
	gs := game.NewGame()

	world.InitializeObjectStates(gs)

	tui.Run(gs)
}
