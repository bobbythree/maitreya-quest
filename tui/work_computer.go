package tui

import (
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
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
	if m.gameState.WorkComputer == nil {
		return m, nil
	}

	m.gameState.WorkComputer.ScanStep++

	// continue scan until the fault is reached
	if m.gameState.WorkComputer.ScanStep < 4 {
		return m, nextScanStep()
	}

	m.gameState.WorkComputer.Screen = "fault"

	return m, nil
}

// workComputerView renders the work computer interface.
func (m Model) workComputerView() string {
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
	switch m.gameState.WorkComputer.Screen {
	case "menu":
		screen.WriteString(m.workComputerMenu())

	case "scanning":
		switch m.gameState.WorkComputer.ScanStep {
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
			m.gameState.WorkComputer.Screen = "scanning"
			m.gameState.WorkComputer.ScanStep = 0

			return m, nextScanStep()
		case 1:
			// exit
			m.gameState.WorkComputer = nil
			m.workComputerUI.cursor = 0
		}
	}

	return m, nil
}
