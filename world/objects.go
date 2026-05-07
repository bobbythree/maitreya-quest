package world

type Object struct {
	ID          string
	Name        string
	Description string
	Portable    bool
	Container   bool
	Open        bool
	Contains    []string
}

var Objects = map[string]Object{
	"desk": {
		ID:          "desk",
		Name:        "desk",
		Description: "An old wooden desk with a drawer.",
		Container:   true,
		Open:        true,
		Contains: []string{
			"key",
		},
	},

	"computer": {
		ID:          "computer",
		Name:        "computer",
		Description: "A dusty old computer.",
	},

	"window": {
		ID:          "window",
		Name:        "window",
		Description: "Rain taps softly against the glass.",
	},

	"door": {
		ID:          "door",
		Name:        "door",
		Description: "A heavy metal door.",
	},

	"thumbdrive": {
		ID:          "thumbdrive",
		Name:        "thumbdrive",
		Description: "A small thumbdrive.",
		Portable:    true,
	},
}
