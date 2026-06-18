// Package parser
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

	firstWord := words[0]

	// normalize direction aliases
	if alias, ok := DirectionAliases[firstWord]; ok {
		firstWord = alias
	}

	// handle bare directions without 'go' verb
	if Directions[firstWord] {
		cmd.Verb = "go"
		cmd.DirectObject = firstWord
		return cmd
	}

	verb, ok := VerbAliases[firstWord]
	if ok {
		cmd.Verb = verb
	} else {
		cmd.Verb = firstWord
	}

	if len(words) == 1 {
		return cmd
	}

	prepIndex := -1

	for i := 1; i < len(words); i++ {
		if Prepositions[words[i]] {
			prepIndex = i
			cmd.Preposition = words[i]
			break
		}
	}

	// no preposition:
	// "look door"
	if prepIndex == -1 {
		cmd.DirectObject = words[1]
		return cmd
	}

	// preposition immediately after verb:
	// "look at door"
	// "talk to man"
	if prepIndex == 1 {
		if prepIndex < len(words)-1 {
			cmd.DirectObject = words[prepIndex+1]
		}
		return cmd
	}

	// preposition after direct object:
	// "use key on door"
	cmd.DirectObject = words[1]

	if prepIndex < len(words)-1 {
		cmd.IndirectObject = words[prepIndex+1]
	}

	return cmd
}
