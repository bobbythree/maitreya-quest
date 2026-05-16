package output

import "fmt"

func Print(text string) {
	fmt.Print(Render(text))
}

func Println(text string) {
	fmt.Println(Render(text))
}

func Printf(format string, args ...any) {
	Print(fmt.Sprintf(format, args...))
}

func Render(text string) string {
	return GlobalGameStyle.Render(text)
}
