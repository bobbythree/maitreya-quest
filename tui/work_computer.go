package tui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

// workComputerUIState holds work computer-specific TUI state.
type workComputerUIState struct {
	cursor int
}

// workComputerView renders the work computer interface.
func (m Model) workComputerView() string {
	computerStyle := lipgloss.NewStyle().
		Width(64).
		Border(lipgloss.DoubleBorder()).
		Padding(1, 2)

	titleStyle := lipgloss.NewStyle().
		Bold(true)

	selectedStyle := lipgloss.NewStyle().
		Bold(true)

	choices := []string{
		"Scan facility",
		"Exit",
	}

	var screen strings.Builder

	screen.WriteString(titleStyle.Render("LEARNEX FACILITY SECURITY SYSTEM"))
	screen.WriteString("\n")
	screen.WriteString("SYSTEM STATUS: ONLINE")
	screen.WriteString("\n\n")

	for i, choice := range choices {
		if i == m.workComputerUI.cursor {
			screen.WriteString(selectedStyle.Render("> " + choice))
		} else {
			screen.WriteString("  ")
			screen.WriteString(choice)
		}

		screen.WriteString("\n")
	}

	return computerStyle.Render(screen.String())
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
			// scan facility
			return m, nil
		case 1:
			// exit
			m.gameState.WorkComputer = nil
			m.workComputerUI.cursor = 0
		}
	}

	return m, nil
}
