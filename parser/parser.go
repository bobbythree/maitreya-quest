package parser

import "strings"

type Command struct {
	Verb           string
	DirectObject   string
	Prepositions   string
	IndirectObject string
}

func Parse(input string) Command {
	words := strings.Fields(strings.ToLower(input))
	cmd := Command{}
	verb, ok := VerbAliases[words[0]]
	if ok {
		cmd.Verb = verb
	} else {
		cmd.Verb = words[0]
	}

	if len(words) > 0 {
		cmd.Verb = words[0]
	}

	if len(words) > 1 {
		cmd.DirectObject = words[1]
	}

	return cmd
}
