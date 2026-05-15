package tui

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bobbythree/maitreya-quest/actions"
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/lsferreira42/figlet-go/figlet"
)

func Run(gs *game.GameState) {
	fmt.Print("\033[H\033[2J")
	scanner := bufio.NewScanner(os.Stdin)

	openingNarration := "Our story takes place on Earth 100 years in the future and roughly 100 years since humankind achieved AGI (Artificial General Intelligence). As a result of handing nearly all creative and intellectual tasks over to AI long ago, the human mind has atrophied to a critical extent. The majority of humans are either almost too dumb to talk to, or animalistally violent. A prophecy tells of someone called 'Maitreya', who along with the help of an 'other wordly being', will over take the AI and restore humanity to it's former creative and intellectual glory."

	// logo
	logo, err := figlet.Render(
		"MAITREYA'S QUEST",
		figlet.WithFont("smkeyboard"),
		figlet.WithColors(figlet.ColorCyan),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(logo)

	time.Sleep(1 * time.Second)
	typewriterEffect(openingNarration)
	fmt.Println()

	// game loop
	for {

		fmt.Print("\n> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		// check for empty command (user hits enter with no command)
		if strings.TrimSpace(input) == "" {
			continue
		}

		cmd := parser.Parse(input)

		action, ok := actions.ActionMap[cmd.Verb]

		if !ok {
			fmt.Println("I don't understand.")
			continue
		}

		action(gs, cmd)
	}
}

// other funcs
func typewriterEffect(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(10 * time.Millisecond)
	}
}
