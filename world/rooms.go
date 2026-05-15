package world

type Room struct {
	ID          string
	Name        string
	Description string
	Objects     []string
	Exits       map[string]string
}

var Rooms = map[string]Room{
	"bedroom": {
		ID:          "bedroom",
		Name:        "bedroom",
		Description: "This is your bedroom. You see a [desk] with a [computer] on it, your [bed], a [window] and a [door].",
		Objects: []string{
			"desk",
			"computer",
			"window",
			"door",
			"bed",
		},
		Exits: map[string]string{
			"east": "hallway",
		},
	},
	"hallway": {
		ID:          "hallway",
		Name:        "hallway",
		Description: "You stand on the small landing at the top of the stairs. The door back into your apartment is to the [west]. A staircase leads [down] to the street",
		Exits: map[string]string{
			"west": "bedroom",
			"down": "street",
		},
	},
	"street": {
		ID:          "street",
		Name:        "street",
		Description: "The street outside your apartment is quiet. You see a man standing on the corner.",
		Exits: map[string]string{
			"west": "hallway",
		},
	},
}
