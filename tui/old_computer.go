package tui

import (
	"strings"
	"time"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/bobbythree/maitreya-quest/game"
)

// oldComputerUIState holds old computer-specific TUI state.
type oldComputerUIState struct {
	passwordInput textinput.Model
	generation    uint64
}

// oldComputerStepMsg advances the old computer's timed sequence.
type oldComputerStepMsg struct {
	generation uint64
}

func newOldComputerUIState() oldComputerUIState {
	input := textinput.New()
	input.Prompt = ""

	return oldComputerUIState{passwordInput: input}
}

// nextOldComputerStep waits before advancing the old computer.
func nextOldComputerStep(generation uint64) tea.Cmd {
	return tea.Tick(1500*time.Millisecond, func(time.Time) tea.Msg {
		return oldComputerStepMsg{generation: generation}
	})
}

// beginOldComputerUI resets transient UI state and starts the loading timer.
func (m *Model) beginOldComputerUI() tea.Cmd {
	m.oldComputerUI.passwordInput.SetValue("")
	m.oldComputerUI.passwordInput.Blur()
	m.oldComputerUI.generation++

	return nextOldComputerStep(m.oldComputerUI.generation)
}

// updateOldComputerStep advances loading or completes shutdown.
func (m Model) updateOldComputerStep(msg oldComputerStepMsg) (tea.Model, tea.Cmd) {
	if msg.generation != m.oldComputerUI.generation {
		return m, nil
	}

	computer, ok := m.gameState.ActiveOldComputer()
	if !ok {
		return m, nil
	}

	switch computer.Screen {
	case "loading":
		computer.Screen = "password"
		return m, m.oldComputerUI.passwordInput.Focus()

	case "shutdown":
		m.gameState.EndInteraction(game.InteractionOldComputer)
		m.oldComputerUI.passwordInput.SetValue("")
		m.oldComputerUI.passwordInput.Blur()
		m.history = append(m.history, "The old computer sputters and dies. It won't turn on again.")
	}

	return m, nil
}

// updateOldComputer handles input while the old computer is active.
func (m Model) updateOldComputer(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	computer, ok := m.gameState.ActiveOldComputer()
	if !ok || computer.Screen != "password" {
		return m, nil
	}

	if msg.String() == "enter" {
		m.gameState.Flags["old_computer_destroyed"] = true
		computer.Screen = "shutdown"
		m.oldComputerUI.passwordInput.Blur()
		m.oldComputerUI.generation++

		return m, nextOldComputerStep(m.oldComputerUI.generation)
	}

	var cmd tea.Cmd
	m.oldComputerUI.passwordInput, cmd = m.oldComputerUI.passwordInput.Update(msg)

	return m, cmd
}

// oldComputerView renders the old computer interface.
func (m Model) oldComputerView() string {
	computer, ok := m.gameState.ActiveOldComputer()
	if !ok {
		return ""
	}

	computerStyle := lipgloss.NewStyle().
		Width(64).
		Border(lipgloss.DoubleBorder()).
		Padding(1, 2)

	var screen strings.Builder

	switch computer.Screen {
	case "loading":
		screen.WriteString("LOADING DATA FROM EXTERNAL DISK...")

	case "password":
		screen.WriteString("DISK ENCRYPTED")
		screen.WriteString("\n")
		screen.WriteString("ENTER PASSWORD:")
		screen.WriteString("\n")
		screen.WriteString(m.oldComputerUI.passwordInput.View())

	case "shutdown":
		screen.WriteString("ACCESS DENIED")
		screen.WriteString("\n")
		screen.WriteString("SYSTEM FAILURE")
		screen.WriteString("\n")
		screen.WriteString("SHUTTING DOWN...")
	}

	return computerStyle.Render(screen.String())
}
