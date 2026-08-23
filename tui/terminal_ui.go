// Package tui
package tui

import (
	"fmt"
	"log"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/bobbythree/maitreya-quest/actions"
	"github.com/bobbythree/maitreya-quest/dialogue"
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/lsferreira42/figlet-go/figlet"
)

// bubble tea

type Model struct {
	gameState      *game.GameState
	input          textinput.Model
	history        []string
	width          int
	logo           string
	intro          string
	dialogueCursor int
}

func (m Model) Init() tea.Cmd {
	return nil
}

func NewModel(gs *game.GameState) Model {
	intro := "Our story takes place on Earth 100 years in the future and roughly 100 years since humankind achieved AGI (Artificial General Intelligence). As a result of handing nearly all creative and intellectual tasks over to AI long ago, the human mind has atrophied to a critical extent. The majority of humans are either almost too dumb to talk to, or animalistally violent. A prophecy tells of someone called 'Maitreya', who along with the help of an 'other wordly being', will over take the AI and restore humanity to it's former creative and intellectual glory."

	logo, err := figlet.Render(
		"MAITREYA'S QUEST",
		figlet.WithFont("smkeyboard"),
		figlet.WithColors(figlet.ColorCyan),
	)
	if err != nil {
		log.Fatal(err)
	}

	input := textinput.New()
	input.Focus()

	return Model{
		gameState: gs,
		input:     input,
		logo:      logo,
		intro:     intro,
	}
}

// dialogue helper

func (m Model) dialogueView() string {
	if m.gameState.Dialogue == nil {
		return ""
	}

	state := m.gameState.Dialogue

	d, ok := dialogue.Dialogues[state.DialogueID]
	if !ok {
		return ""
	}

	node, ok := d.Nodes[state.NodeID]
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

		if i == m.dialogueCursor {
			choiceText = selectedStyle.Render("> " + choiceText)
		} else {
			choiceText = "  " + choiceText
		}

		result.WriteString("\n")
		result.WriteString(choiceText)
	}

	return dialogueStyle.Render(result.String())
}

func (m Model) updateDialogue(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	state := m.gameState.Dialogue

	d := dialogue.Dialogues[state.DialogueID]
	node := d.Nodes[state.NodeID]

	switch msg.String() {
	case "up", "k":
		if m.dialogueCursor > 0 {
			m.dialogueCursor--
		}

	case "down", "j":
		if m.dialogueCursor < len(node.Choices)-1 {
			m.dialogueCursor++
		}

	case "enter":
		choice := node.Choices[m.dialogueCursor]

		if choice.NextNode == "" {
			m.gameState.Dialogue = nil
			m.dialogueCursor = 0
			return m, nil
		}

		m.gameState.Dialogue.NodeID = choice.NextNode
		m.dialogueCursor = 0
	}

	return m, nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if m.gameState.Dialogue != nil {
			return m.updateDialogue(msg)
		}

		// player executes command
		if msg.String() == "enter" {
			input := m.input.Value()
			// handle empty command
			if strings.TrimSpace(input) == "" {
				return m, nil
			}

			cmd := parser.Parse(input)
			action, ok := actions.ActionMap[cmd.Verb]

			if !ok {
				entry := "> " + input + "\nI don't get it."
				m.history = append(m.history, entry)
				m.input.SetValue("")

				return m, nil
			}

			result := action(m.gameState, cmd)
			entry := "> " + input
			if result != "" {
				entry += "\n" + result
			}

			m.history = append(m.history, entry)

			m.input.SetValue("")

			return m, nil
		}

	// term resize
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	return m, cmd
}

func (m Model) View() tea.View {
	// tui sizing
	contentWidth := 80

	if m.width > 0 && m.width-8 < contentWidth {
		contentWidth = m.width - 8
	}

	if contentWidth < 1 {
		contentWidth = 1
	}

	style := lipgloss.NewStyle().
		Width(contentWidth).
		Padding(0, 4)

	history := strings.Join(m.history, "\n\n")
	bottom := m.input.View()

	if m.gameState.Dialogue != nil {
		bottom = m.dialogueView()
	}

	content := m.logo + "\n" +
		m.intro + "\n\n" +
		history + "\n\n" +
		bottom

	//return view
	return tea.NewView(style.Render(content))
}

func Run(gs *game.GameState) {
	m := NewModel(gs)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
