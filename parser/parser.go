package parser

import "strings"

type Command struct {
	Verb           string
	DirectObject   string
	Preposition    string
	IndirectObject string
}

func Parse(input string) Command {
	words := strings.Fields(strings.ToLower(input))
	cmd := Command{}

	if len(words) == 0 {
		return cmd
	}

	verb, ok := VerbAliases[words[0]]
	if ok {
		cmd.Verb = verb
	} else {
		cmd.Verb = words[0]
	}

	if len(words) == 1 {
		return cmd
	}

	prepIndex := -1                   // no prep found yet
	for i := 1; i < len(words); i++ { // start loop at word 1, not 0 (skip verb)
		if Prepositions[words[i]] {
			prepIndex = i
			cmd.Preposition = words[i]
			break
		}
	}

	// if no prep
	if prepIndex == -1 {
		cmd.DirectObject = words[1]
		return cmd
	}

	// if prep is after word 1, then word 1 is dirObj
	if prepIndex > 1 {
		cmd.DirectObject = words[1]
	}

	// if there's a word after the prep, that word is indObj
	if prepIndex < len(words)-1 {
		cmd.IndirectObject = words[prepIndex+1]
	}

	return cmd
}
