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
	input.EchoMode = textinput.EchoPassword
	input.EchoCharacter = '*'

	return oldComputerUIState{passwordInput: input}
}

// nextOldComputerStep waits before advancing the old computer.
func nextOldComputerStep(generation uint64) tea.Cmd {
	return tea.Tick(4*time.Second, func(time.Time) tea.Msg {
		return oldComputerStepMsg{generation: generation}
	})
}

// beginOldComputerUI resets transient UI state for either computer.
func (m *Model) beginOldComputerUI() tea.Cmd {
	m.oldComputerUI.passwordInput.SetValue("")
	m.oldComputerUI.passwordInput.Blur()
	m.oldComputerUI.generation++

	computer, _ := m.gameState.ActiveOldComputer()
	if computer.Screen == "password" {
		return m.oldComputerUI.passwordInput.Focus()
	}
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
	case "awakening":
		computer.Screen = "loading"
		return m, nextOldComputerStep(m.oldComputerUI.generation)

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
		if computer.Computer == "home" {
			// Introduce the correct password and successful login here when known.
			computer.Error = "password incorrect, try again"
			m.oldComputerUI.passwordInput.SetValue("")
			return m, nil
		}
		m.gameState.Flags["old_computer_destroyed"] = true
		computer.Screen = "shutdown"
		m.oldComputerUI.passwordInput.Blur()
		m.oldComputerUI.generation++

		return m, nextOldComputerStep(m.oldComputerUI.generation)
	}
	if msg.String() == "esc" && computer.Computer == "home" {
		m.gameState.EndInteraction(game.InteractionOldComputer)
		m.oldComputerUI.passwordInput.SetValue("")
		m.oldComputerUI.passwordInput.Blur()
		return m, nil
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
	if computer.Screen == "awakening" {
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
		if computer.Computer == "home" {
			screen.WriteString("COMPUTER LOCKED")
		} else {
			screen.WriteString("DISK ENCRYPTED")
		}
		screen.WriteString("\n")
		screen.WriteString("ENTER PASSWORD:")
		screen.WriteString("\n")
		screen.WriteString(m.oldComputerUI.passwordInput.View())
		if computer.Error != "" {
			screen.WriteString("\n")
			screen.WriteString(computer.Error)
		}
		if computer.Computer == "home" {
			screen.WriteString("\nPress Esc to leave")
		}

	case "shutdown":
		screen.WriteString("ACCESS DENIED")
		screen.WriteString("\n")
		screen.WriteString("SYSTEM FAILURE")
		screen.WriteString("\n")
		screen.WriteString("SHUTTING DOWN...")
	}

	return computerStyle.Render(screen.String())
}
