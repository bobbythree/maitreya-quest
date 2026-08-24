package game

func MarkDialogueNodeVisited(gs *GameState, dialogueID, nodeID string) {
	if gs.VisitedDialogueNodes[dialogueID] == nil {
		gs.VisitedDialogueNodes[dialogueID] = make(map[string]bool)
	}

	gs.VisitedDialogueNodes[dialogueID][nodeID] = true
}

func HasVisitedDialogueNode(gs *GameState, dialogueID, nodeID string) bool {
	return gs.VisitedDialogueNodes[dialogueID][nodeID]
}

func HasCompletedSilasIntroduction(gs *GameState) bool {
	return HasVisitedDialogueNode(gs, "street_man", "explanation") &&
		HasVisitedDialogueNode(gs, "street_man", "maitreya") &&
		HasVisitedDialogueNode(gs, "street_man", "otherworldly")
}
