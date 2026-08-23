// Package tui
package tui

import (
	"log"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/bobbythree/maitreya-quest/actions"
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/lsferreira42/figlet-go/figlet"
)

// bubble tea

type Model struct {
	gameState *game.GameState
	input     textinput.Model
	history   []string
	width     int
	logo      string
	intro     string
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

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
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

			entry := "> " + input + "\n" + result
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
	content := m.logo + "\n" +
		m.intro + "\n\n" +
		history + "\n\n" +
		m.input.View()

	//return view
	return tea.NewView(style.Render(content))
}

// Run -  new run func
func Run(gs *game.GameState) {
	m := NewModel(gs)
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
