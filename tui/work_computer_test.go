package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/workcomputer"
)

func startTestScan(t *testing.T, model Model) Model {
	t.Helper()

	updated, cmd := model.updateWorkComputer(tea.KeyPressMsg(tea.Key{Code: tea.KeyEnter}))
	if cmd == nil {
		t.Fatal("starting scan did not schedule a tick")
	}

	return updated.(Model)
}

func TestWorkComputerIgnoresTickFromPreviousScan(t *testing.T) {
	gs := game.NewGame()
	workcomputer.Start(gs)
	model := startTestScan(t, Model{gameState: gs})
	oldGeneration := model.workComputerUI.scanGeneration

	model = startTestScan(t, model)
	currentGeneration := model.workComputerUI.scanGeneration
	if currentGeneration == oldGeneration {
		t.Fatal("restarting scan did not advance its generation")
	}

	updated, cmd := model.updateWorkComputerScan(scanStepMsg{generation: oldGeneration})
	model = updated.(Model)
	computer, _ := gs.ActiveWorkComputer()

	if computer.ScanStep != 0 {
		t.Fatalf("stale tick advanced restarted scan to step %d", computer.ScanStep)
	}

	if cmd != nil {
		t.Fatal("stale tick scheduled another scan tick")
	}

	updated, cmd = model.updateWorkComputerScan(scanStepMsg{generation: currentGeneration})
	computer, _ = gs.ActiveWorkComputer()
	if computer.ScanStep != 1 {
		t.Fatalf("current tick advanced scan to step %d, want 1", computer.ScanStep)
	}

	if cmd == nil {
		t.Fatal("current tick did not schedule the next scan step")
	}
}

func TestWorkComputerIgnoresTickAfterExitAndReopen(t *testing.T) {
	gs := game.NewGame()
	workcomputer.Start(gs)
	model := startTestScan(t, Model{gameState: gs})
	oldGeneration := model.workComputerUI.scanGeneration

	gs.EndInteraction(game.InteractionWorkComputer)
	workcomputer.Start(gs)

	updated, cmd := model.updateWorkComputerScan(scanStepMsg{generation: oldGeneration})
	model = updated.(Model)
	computer, _ := gs.ActiveWorkComputer()

	if computer.Screen != "menu" || computer.ScanStep != 0 {
		t.Fatalf("stale tick changed reopened computer to screen %q, step %d", computer.Screen, computer.ScanStep)
	}

	if cmd != nil {
		t.Fatal("stale tick scheduled another scan tick after reopen")
	}

	model = startTestScan(t, model)
	newGeneration := model.workComputerUI.scanGeneration
	if newGeneration == oldGeneration {
		t.Fatal("new scan reused the previous scan generation")
	}

	updated, _ = model.updateWorkComputerScan(scanStepMsg{generation: oldGeneration})
	computer, _ = gs.ActiveWorkComputer()
	if computer.ScanStep != 0 {
		t.Fatalf("old tick advanced new scan to step %d", computer.ScanStep)
	}
}
