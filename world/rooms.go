package world

type Room struct {
	ID          string
	Name        string
	Description string
	Objects     []string
	Exits       map[string]string
}

var Rooms = map[string]Room{
	"apartment": {
		ID:          "apartment",
		Name:        "apartment",
		Description: "You are in your small studio apartment which is dimly lit only by artificial light coming in though the window. In the room is your bed, your computer which sits on a desk, one window and a door to the east.",
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
			"west": "apartment",
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
