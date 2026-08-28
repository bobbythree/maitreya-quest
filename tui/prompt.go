package tui

import (
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/bobbythree/maitreya-quest/actions"
	"github.com/bobbythree/maitreya-quest/parser"
)

// handle normal text prompt input
func (m Model) updatePrompt(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	// execute entered command
	if msg.String() == "enter" {
		input := m.input.Value()

		// ignore empty input
		if strings.TrimSpace(input) == "" {
			return m, nil
		}

		cmd := parser.Parse(input)
		action, ok := actions.ActionMap[cmd.Verb]

		// handle unknown command
		if !ok {
			entry := "> " + input + "\nI don't get it."
			m.history = append(m.history, entry)
			m.input.SetValue("")

			return m, nil
		}

		// execute command
		previousRoom := m.gameState.CurrentRoom
		result := action(m.gameState, cmd)
		roomChanged := previousRoom != m.gameState.CurrentRoom

		// build history entry
		entry := "> " + input
		if result != "" {
			entry += "\n" + result
		}

		// reset history when entering a new room
		if roomChanged {
			m.history = []string{result}
			m.showIntro = false
		} else {
			m.history = append(m.history, entry)
		}

		m.input.SetValue("")

		return m, nil
	}

	// update text input
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}
