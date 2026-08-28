package main

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/tui"
)

func main() {
	gs := game.NewGame()

	tui.Run(gs)
}
