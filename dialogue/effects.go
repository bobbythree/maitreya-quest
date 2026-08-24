package dialogue

import "github.com/bobbythree/maitreya-quest/game"

func unlockWork(gs *game.GameState) string {
	gs.Flags["work_unlocked"] = true

	return "As you walk away, you glance at the time. Shit. You're going to be late for work. You'd better get over there."
}
