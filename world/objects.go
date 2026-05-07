package world

type Object struct {
	ID          string
	Name        string
	Description string
	Portable    bool
	Container   bool
	Openable    bool
	Open        bool
	Parent      string
	Contains    []string
}

var Objects = map[string]Object{
	"desk": {
		ID:          "desk",
		Name:        "desk",
		Description: "An old wooden desk with a drawer.",
	},

	"drawer": {
		ID:          "drawer",
		Name:        "drawer",
		Description: "A wooden drawer built into the desk.",

		Container: true,
		Openable:  true,
		Open:      false,

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
		ID:          "window",
		Name:        "window",
		Description: "The window is painted black and you cannot see out.",
		Openable:    true,
		Open:        false,
	},

	"door": {
		ID:          "door",
		Name:        "door",
		Description: "Your bedroom door.",
		Openable:    true,
		Open:        false,
	},

	"thumbdrive": {
		ID:          "thumbdrive",
		Name:        "thumbdrive",
		Description: "A small thumbdrive.",
		Portable:    true,
	},
}
