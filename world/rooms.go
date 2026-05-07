package world

type Room struct {
	ID          string
	Name        string
	Description string
	Objects     []string
}

var Rooms = map[string]Room{
	"bedroom": {
		ID:          "bedroom",
		Name:        "bedroom",
		Description: "This is your bedroom. You see a desk with a computer on it, a window and a door.",
		Objects: []string{
			"desk",
			"computer",
			"window",
			"door",
		},
	},
}
