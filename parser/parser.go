package parser

import "strings"

type Command struct {
	Verb string
	Noun string
}

func Parse(input string) Command {
	words := strings.Fields(strings.ToLower(input))
	cmd := Command{}

	if len(words) > 0 {
		cmd.Verb = words[0]
	}

	if len(words) > 1 {
		cmd.Noun = words[1]
	}

	return cmd
}
