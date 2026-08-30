package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/bobbythree/maitreya-quest/dialogue"
)

// dialogueUIState holds dialogue-specific TUI state.
type dialogueUIState struct {
	cursor int
}

// dialogueView renders the active dialogue.
func (m Model) dialogueView() string {
	node, ok := dialogue.CurrentNode(m.gameState)
	if !ok {
		return ""
	}

	portrait := `
      _______
    /  ~   ~  \
    |  o   o  |
    |    >    |
    |  _____  |
     \_______/
`

	speakerStyle := lipgloss.NewStyle().
		Bold(true)

	selectedStyle := lipgloss.NewStyle().
		Bold(true)

	dialogueStyle := lipgloss.NewStyle().
		Width(64).
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2)

	var result strings.Builder

	result.WriteString(portrait)
	result.WriteString("\n")
	result.WriteString(speakerStyle.Render(node.Speaker))
	result.WriteString("\n\n")
	result.WriteString(node.Text)
	result.WriteString("\n")

	for i, choice := range node.Choices {
		choiceText := fmt.Sprintf("%d. %s", i+1, choice.Text)

		if i == m.dialogueUI.cursor {
			choiceText = selectedStyle.Render("> " + choiceText)
		} else {
			choiceText = "  " + choiceText
		}

		result.WriteString("\n")
		result.WriteString(choiceText)
	}

	return dialogueStyle.Render(result.String())
}

// updateDialogue handles input while dialogue is active.
func (m Model) updateDialogue(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	node, ok := dialogue.CurrentNode(m.gameState)
	if !ok {
		return m, nil
	}

	switch msg.String() {
	case "up", "k":
		if m.dialogueUI.cursor > 0 {
			m.dialogueUI.cursor--
		}

	case "down", "j":
		if m.dialogueUI.cursor < len(node.Choices)-1 {
			m.dialogueUI.cursor++
		}

	case "enter":
		outcome, err := dialogue.SelectChoice(
			m.gameState,
			m.dialogueUI.cursor,
		)
		if err != nil {
			return m, nil
		}

		if outcome.Narration != "" {
			m.history = append(m.history, outcome.Narration)
		}

		// reset selection for the next dialogue node
		m.dialogueUI.cursor = 0
	}

	return m, nil
}
