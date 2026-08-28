package world

import "github.com/bobbythree/maitreya-quest/game"

type Object struct {
	ID                string
	Name              string
	Description       string
	OpenDescription   string
	ClosedDescription string
	Portable          bool
	Container         bool
	Openable          bool
	Open              bool
	Lockable          bool
	Locked            bool
	Talkable          bool
	DialogueID        string
	Parent            string
	Contains          []string
	UseAction         func(*game.GameState) string
}

var Objects = map[string]Object{
	"desk": {
		ID:          "desk",
		Name:        "desk",
		Description: "An old wooden desk with a drawer.",
		Container:   true,
		Contains: []string{
			"drawer",
		},
	},

	"drawer": {
		ID:                "drawer",
		Name:              "drawer",
		ClosedDescription: "A wooden drawer built into the desk.",
		OpenDescription:   "The drawer is open.",
		Container:         true,
		Openable:          true,
		Open:              false,

		Parent: "desk",

		Contains: []string{
			"thumbdrive",
		},
	},

	"bed": {
		ID:          "bed",
		Name:        "bed",
		Description: "a small mattress on the floor. Typical.",
	},

	"computer": {
		ID:          "computer",
		Name:        "computer",
		Description: "Your computer sits on top of the desk.",
	},

	"window": {
		ID:                "window",
		Name:              "window",
		ClosedDescription: "The window is very dirty and difficult to see through. All you see is the indication of artificial light outside.",
		OpenDescription:   "Looking out the open window you can see many video screens up on tall poles that are broadcasting ads for various corporate interests such as: Pear, MacroFirm and Moser",
		Openable:          true,
		Open:              false,
	},

	"door": {
		ID:                "door",
		Name:              "door",
		ClosedDescription: "Your apartment door.",
		OpenDescription:   "Looking out the door you see a small landing that leads to a downward staircase.",
		Openable:          true,
		Open:              false,
	},

	"thumbdrive": {
		ID:          "thumbdrive",
		Name:        "thumbdrive",
		Description: "A really old thumbdrive. Not a model that you recognize.",
		Parent:      "drawer",
		Portable:    true,
	},
	"street_man": {
		ID:          "street_man",
		Name:        "man",
		Description: "Wow, this dude looks like a total bad ass.",
		Talkable:    true,
		DialogueID:  "street_man",
	},
	"work_computer": {
		ID:          "work_computer",
		Name:        "computer",
		Description: "Your work machine.",
	},
}
