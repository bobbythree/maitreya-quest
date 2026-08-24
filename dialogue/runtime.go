package dialogue

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
)

func Start(gs *game.GameState, dialogueID string) error {
	d, ok := Dialogues[dialogueID]
	if !ok {
		return fmt.Errorf("dialogue %q not found", dialogueID)
	}

	if _, ok := d.Nodes[d.StartNode]; !ok {
		return fmt.Errorf("dialogue start node %q not found", d.StartNode)
	}

	if gs.DialogueProgress == nil {
		gs.DialogueProgress = make(map[string]*game.DialogueProgress)
	}

	if gs.DialogueProgress[dialogueID] == nil {
		gs.DialogueProgress[dialogueID] = &game.DialogueProgress{
			VisitedNodes: make(map[string]bool),
		}
	}

	if gs.DialogueProgress[dialogueID].VisitedNodes == nil {
		gs.DialogueProgress[dialogueID].VisitedNodes = make(map[string]bool)
	}

	gs.Dialogue = &game.DialogueState{
		DialogueID: dialogueID,
		NodeID:     d.StartNode,
	}

	gs.DialogueProgress[dialogueID].VisitedNodes[d.StartNode] = true

	return nil
}

func CurrentNode(gs *game.GameState) (Node, bool) {
	if gs.Dialogue == nil {
		return Node{}, false
	}

	d, ok := Dialogues[gs.Dialogue.DialogueID]
	if !ok {
		return Node{}, false
	}

	node, ok := d.Nodes[gs.Dialogue.NodeID]
	if !ok {
		return Node{}, false
	}

	return node, true
}

func SelectChoice(gs *game.GameState, choiceIndex int) (Outcome, error) {
	if gs.Dialogue == nil {
		return Outcome{}, fmt.Errorf("no active dialogue")
	}

	dialogueID := gs.Dialogue.DialogueID

	d, ok := Dialogues[dialogueID]
	if !ok {
		return Outcome{}, fmt.Errorf("dialogue %q not found", dialogueID)
	}

	node, ok := d.Nodes[gs.Dialogue.NodeID]
	if !ok {
		return Outcome{}, fmt.Errorf("dialogue node %q not found", gs.Dialogue.NodeID)
	}

	if choiceIndex < 0 || choiceIndex >= len(node.Choices) {
		return Outcome{}, fmt.Errorf("invalid choice index %d", choiceIndex)
	}

	choice := node.Choices[choiceIndex]

	if choice.NextNode != "" {
		if _, ok := d.Nodes[choice.NextNode]; !ok {
			return Outcome{}, fmt.Errorf("dialogue node %q not found", choice.NextNode)
		}

		gs.Dialogue.NodeID = choice.NextNode
		gs.DialogueProgress[dialogueID].VisitedNodes[choice.NextNode] = true

		return Outcome{}, nil
	}

	outcome := Outcome{
		Ended: true,
	}

	progress := gs.DialogueProgress[dialogueID]

	if d.Completion != nil && !progress.Completed {
		complete := true

		for _, requiredNode := range d.Completion.RequiredNodes {
			if !progress.VisitedNodes[requiredNode] {
				complete = false
				break
			}
		}

		if complete {
			progress.Completed = true

			if d.Completion.OnComplete != nil {
				outcome.Narration = d.Completion.OnComplete(gs)
			}
		}
	}

	gs.Dialogue = nil

	return outcome, nil
}
