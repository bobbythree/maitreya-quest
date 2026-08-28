package tui

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

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
		"Scan facility rooms",
		"View system status",
		"Log out",
	}

	var screen strings.Builder

	screen.WriteString(titleStyle.Render("LEARNEX FACILITY SECURITY SYSTEM"))
	screen.WriteString("\n")
	screen.WriteString("SYSTEM STATUS: ONLINE")
	screen.WriteString("\n\n")

	for i, choice := range choices {
		if i == m.gameState.WorkComputer.Cursor {
			screen.WriteString(selectedStyle.Render("> " + choice))
		} else {
			screen.WriteString("  ")
			screen.WriteString(choice)
		}

		screen.WriteString("\n")
	}

	return computerStyle.Render(screen.String())
}

func (m Model) updateWorkComputer(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.gameState.WorkComputer.Cursor > 0 {
			m.gameState.WorkComputer.Cursor--
		}

	case "down", "j":
		if m.gameState.WorkComputer.Cursor < 2 {
			m.gameState.WorkComputer.Cursor++
		}
	}

	return m, nil
}
