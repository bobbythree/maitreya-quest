package tui

import (
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/bobbythree/maitreya-quest/game"
)

// workComputerUIState holds work computer-specific TUI state.
type workComputerUIState struct {
	cursor int
}

// scanStepMsg advances the facility scan.
type scanStepMsg struct{}

// nextScanStep waits before advancing the scan.
func nextScanStep() tea.Cmd {
	return tea.Tick(1500*time.Millisecond, func(time.Time) tea.Msg {
		return scanStepMsg{}
	})
}

// updateWorkComputerScan advances the facility scan.
func (m Model) updateWorkComputerScan() (tea.Model, tea.Cmd) {
	// ignore scan messages if the computer is no longer active
	computer, ok := m.gameState.ActiveWorkComputer()
	if !ok {
		return m, nil
	}

	computer.ScanStep++

	// continue scan until the fault is reached
	if computer.ScanStep < 4 {
		return m, nextScanStep()
	}

	computer.Screen = "fault"

	return m, nil
}

// workComputerView renders the work computer interface.
func (m Model) workComputerView() string {
	computer, ok := m.gameState.ActiveWorkComputer()
	if !ok {
		return ""
	}

	computerStyle := lipgloss.NewStyle().
		Width(64).
		Border(lipgloss.DoubleBorder()).
		Padding(1, 2)

	titleStyle := lipgloss.NewStyle().
		Bold(true)

	var screen strings.Builder

	screen.WriteString(titleStyle.Render("LEARNEX FACILITY SECURITY SYSTEM"))
	screen.WriteString("\n")
	screen.WriteString("SYSTEM STATUS: ONLINE")
	screen.WriteString("\n\n")

	// render current computer screen
	switch computer.Screen {
	case "menu":
		screen.WriteString(m.workComputerMenu())

	case "scanning":
		switch computer.ScanStep {
		case 0:
			screen.WriteString("SCANNING FACILITY...")
		case 1:
			screen.WriteString("Checking secrity system status...")
		case 2:
			screen.WriteString("Checking access controls...")
		case 3:
			screen.WriteString("Checking door sensors...")
		}
	case "fault":
		screen.WriteString("SENSOR FAULT")
		screen.WriteString("\n\n\n")
		screen.WriteString("[1] View fault details")

	case "fault_details":
		screen.WriteString("SENSOR FAULT")
		screen.WriteString("\n\n")
		screen.WriteString("Room: 108")
		screen.WriteString("\n")
		screen.WriteString("Sensor: Motion")
		screen.WriteString("\n")
		screen.WriteString("Status: No Response")
		screen.WriteString("\n\n")
		screen.WriteString("[1] Unlock security door")
		screen.WriteString("\n")
		screen.WriteString("[2] Exit")

	case "unlocked":
		screen.WriteString("SECURITY DOOR UNLOCKED")
		screen.WriteString("\n\n")
		screen.WriteString("Access authorized for Room 108 investigation.")
		screen.WriteString("\n\n")
		screen.WriteString("[Enter] Exit")
	}

	return computerStyle.Render(screen.String())
}

// workComputerMenu renders the main computer menu.
func (m Model) workComputerMenu() string {
	selectedStyle := lipgloss.NewStyle().
		Bold(true)

	choices := []string{
		"Scan facility",
		"Exit",
	}

	var menu strings.Builder

	for i, choice := range choices {
		if i == m.workComputerUI.cursor {
			menu.WriteString(selectedStyle.Render("> " + choice))
		} else {
			menu.WriteString("  ")
			menu.WriteString(choice)
		}

		menu.WriteString("\n")
	}

	return menu.String()
}

// updateWorkComputer handles input while the work computer is active.
func (m Model) updateWorkComputer(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	computer, ok := m.gameState.ActiveWorkComputer()
	if !ok {
		return m, nil
	}

	// exit after unlocking the security door
	if computer.Screen == "unlocked" {
		if msg.String() == "enter" {
			m.gameState.EndInteraction(game.InteractionWorkComputer)
			m.workComputerUI.cursor = 0
			m.history = append(m.history, "The security door to the [north] is now open.")
		}

		return m, nil
	}

	// show fault details
	if computer.Screen == "fault" {
		if msg.String() == "1" {
			computer.Screen = "fault_details"
		}

		return m, nil
	}

	// handle fault detail choices
	if computer.Screen == "fault_details" {
		if msg.String() == "1" {
			m.gameState.Flags["security_door_unlocked"] = true
			computer.Screen = "unlocked"
		}
		if msg.String() == "2" {
			m.gameState.EndInteraction(game.InteractionWorkComputer)
			m.workComputerUI.cursor = 0
		}

		return m, nil
	}

	switch msg.String() {
	case "up", "k":
		if m.workComputerUI.cursor > 0 {
			m.workComputerUI.cursor--
		}

	case "down", "j":
		if m.workComputerUI.cursor < 1 {
			m.workComputerUI.cursor++
		}
	case "enter":
		switch m.workComputerUI.cursor {
		case 0:
			// start facility scan
			computer.Screen = "scanning"
			computer.ScanStep = 0

			return m, nextScanStep()
		case 1:
			// exit
			m.gameState.EndInteraction(game.InteractionWorkComputer)
			m.workComputerUI.cursor = 0
		}
	}

	return m, nil
}
