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
		screen.WriteString("SCANNING FACILITY...")
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
		case 1:
			// exit
			m.gameState.WorkComputer = nil
			m.workComputerUI.cursor = 0
		}
	}

	return m, nil
}
