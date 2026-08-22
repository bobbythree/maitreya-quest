// Package actions regisrty
package actions

import (
	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
)

type ActionFunc func(*game.GameState, parser.Command) string

var ActionMap = map[string]ActionFunc{
	"look":      Look,
	"get":       Get,
	"open":      Open,
	"close":     Close,
	"use":       Use,
	"go":        Go,
	"talk":      Talk,
	"inventory": Inventory,
}
