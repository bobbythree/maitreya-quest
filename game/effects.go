package game

func ApplyEffect(gs *GameState, effect string) string {
	switch effect {
	case "unlock_work":
		gs.Flags["unlock_work"] = true
		return "As you walk away from Silas, you glance at the time. Shit. You're going to be late for work. You'd better get over there."
	}

	return ""
}
