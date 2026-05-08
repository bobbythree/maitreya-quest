package world

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
	Parent            string
	Contains          []string
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

	"computer": {
		ID:          "computer",
		Name:        "computer",
		Description: "Your computer sits on top of the desk.",
	},

	"window": {
		ID:                "window",
		Name:              "window",
		ClosedDescription: "The window is painted black and you cannot see out.",
		OpenDescription:   "You look out the window and see....",
		Openable:          true,
		Open:              false,
	},

	"door": {
		ID:                "door",
		Name:              "door",
		ClosedDescription: "Your bedroom door.",
		OpenDescription:   "Looking out the door you see a small landing that leads to a downward staircase.",
		Openable:          true,
		Open:              false,
	},

	"thumbdrive": {
		ID:          "thumbdrive",
		Name:        "thumbdrive",
		Description: "A small thumbdrive.",
		Parent:      "drawer",
		Portable:    true,
	},
}
