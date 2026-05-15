package parser

var VerbAliases = map[string]string{
	"look":    "look",
	"examine": "look",
	"inspect": "look",

	"get":  "get",
	"take": "get",
	"grab": "get",

	"open": "open",

	"close": "close",
	"shut":  "close",

	"use": "use",

	"inventory": "inventory",
	"inv":       "inventory",
	"i":         "inventory",
}

var DirectionAliases = map[string]string{
	"n": "north",
	"s": "south",
	"e": "east",
	"w": "west",
	"u": "up",
	"d": "down",
}

var Directions = map[string]bool{
	"north": true,
	"south": true,
	"east":  true,
	"west":  true,
	"up":    true,
	"down":  true,
}
