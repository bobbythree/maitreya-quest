package game

// InteractionKind identifies the single special interaction currently active.
type InteractionKind uint8

const (
	InteractionNone InteractionKind = iota
	InteractionDialogue
	InteractionWorkComputer
)

type activeInteraction struct {
	kind         InteractionKind
	dialogue     *DialogueState
	workComputer *WorkComputerState
}

// InteractionKind returns the active interaction type, or InteractionNone.
func (gs *GameState) InteractionKind() InteractionKind {
	if gs.activeInteraction == nil {
		return InteractionNone
	}

	return gs.activeInteraction.kind
}

// BeginDialogue starts a dialogue if no other special interaction is active.
func (gs *GameState) BeginDialogue(state DialogueState) bool {
	if gs.activeInteraction != nil {
		return false
	}

	gs.activeInteraction = &activeInteraction{
		kind:     InteractionDialogue,
		dialogue: &state,
	}

	return true
}

// ActiveDialogue returns the active dialogue state, if dialogue is active.
func (gs *GameState) ActiveDialogue() (*DialogueState, bool) {
	if gs.activeInteraction == nil || gs.activeInteraction.kind != InteractionDialogue {
		return nil, false
	}

	return gs.activeInteraction.dialogue, true
}

// BeginWorkComputer starts the work computer if no other interaction is active.
func (gs *GameState) BeginWorkComputer(state WorkComputerState) bool {
	if gs.activeInteraction != nil {
		return false
	}

	gs.activeInteraction = &activeInteraction{
		kind:         InteractionWorkComputer,
		workComputer: &state,
	}

	return true
}

// ActiveWorkComputer returns the work-computer state, if it is active.
func (gs *GameState) ActiveWorkComputer() (*WorkComputerState, bool) {
	if gs.activeInteraction == nil || gs.activeInteraction.kind != InteractionWorkComputer {
		return nil, false
	}

	return gs.activeInteraction.workComputer, true
}

// EndInteraction ends the active interaction only when its kind matches.
func (gs *GameState) EndInteraction(kind InteractionKind) bool {
	if gs.activeInteraction == nil || gs.activeInteraction.kind != kind {
		return false
	}

	gs.activeInteraction = nil

	return true
}
