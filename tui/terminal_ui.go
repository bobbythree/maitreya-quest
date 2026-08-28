// Package tui
package tui

import (
	"fmt"
	"log"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/world"
	"github.com/lsferreira42/figlet-go/figlet"
)

// Model holds the main TUI state.
type Model struct {
	gameState      *game.GameState
	input          textinput.Model
	history        []string
	width          int
	logo           string
	intro          string
	showIntro      bool
	dialogueCursor int
}

func (m Model) Init() tea.Cmd {
	return nil
}

// NewModel initializes the TUI.
func NewModel(gs *game.GameState) Model {
	intro := "Our story takes place on Earth 100 years in the future and roughly 100 years since humankind achieved AGI (Artificial General Intelligence). As a result of handing nearly all creative and intellectual tasks over to AI long ago, the human mind has atrophied to a critical extent. The majority of humans are either almost too dumb to talk to, or animalistally violent. A prophecy tells of someone called 'Maitreya', who along with the help of an 'other wordly being', will over take the AI and restore humanity to it's former creative and intellectual glory."

	// render game logo
	logo, err := figlet.Render(
		"MAITREYA'S QUEST",
		figlet.WithFont("smkeyboard"),
		figlet.WithColors(figlet.ColorCyan),
	)
	if err != nil {
		log.Fatal(err)
	}

	// initialize command prompt
	input := textinput.New()
	input.Focus()

	// build initial room text
	room := world.Rooms[gs.CurrentRoom]

	initialText := room.Description
	if room.FirstVisitNarration != "" {
		initialText = room.FirstVisitNarration + "\n\n" + room.Description
	}

	gs.VisitedRooms[gs.CurrentRoom] = true

	return Model{
		gameState: gs,
		history:   []string{initialText},
		input:     input,
		logo:      logo,
		intro:     intro,
		showIntro: true,
	}
}

// Update routes incoming TUI events.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		// quit the game
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		// route dialogue input
		if m.gameState.Dialogue != nil {
			return m.updateDialogue(msg)
		}

		// route work computer input
		if m.gameState.WorkComputer != nil {
			return m.updateWorkComputer(msg)
		}

		// route normal prompt input
		return m.updatePrompt(msg)

	case tea.WindowSizeMsg:
		m.width = msg.Width
	}

	return m, nil
}

// View builds the current terminal view.
func (m Model) View() tea.View {
	// size content to the terminal
	contentWidth := 80

	if m.width > 0 && m.width-8 < contentWidth {
		contentWidth = m.width - 8
	}

	if contentWidth < 1 {
		contentWidth = 1
	}

	style := lipgloss.NewStyle().
		Width(contentWidth).
		Padding(2, 4)

	// choose the active interaction
	history := strings.Join(m.history, "\n\n")
	bottom := m.input.View()

	if m.gameState.Dialogue != nil {
		bottom = m.dialogueView()
	}

	if m.gameState.WorkComputer != nil {
		bottom = m.workComputerView()
	}

	// build room display
	room := world.Rooms[m.gameState.CurrentRoom]
	roomName := strings.ToUpper(room.Name)

	content := ""

	if m.showIntro {
		content += m.logo + "\n" +
			m.intro + "\n\n"
	}

	content += roomName + "\n\n"

	if history != "" {
		content += history + "\n\n"
	}

	content += bottom

	return tea.NewView(style.Render(content))
}

// Run starts the Bubble Tea program.
func Run(gs *game.GameState) {
	// clear the terminal before starting
	fmt.Print("\033[2J\033[H")

	m := NewModel(gs)
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
