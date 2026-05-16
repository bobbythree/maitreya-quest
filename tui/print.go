package tui

import "fmt"

func Print(text string) {
	fmt.Println(GlobalGameStyle.Render(text))
}
