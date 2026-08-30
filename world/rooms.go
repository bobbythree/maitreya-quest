package world

type Room struct {
	ID                  string
	Name                string
	Description         string
	FirstVisitNarration string // optional
	Objects             []string
	Exits               map[string]string
}

var Rooms = map[string]Room{
	"apartment": {
		ID:          "apartment",
		Name:        "Your Apartment",
		Description: "You are in your small studio apartment which is dimly lit only by artificial light coming in though the window. In the room is your bed, your computer which sits on a desk, one window and a door to the east.",
		Objects: []string{
			"desk",
			"home_computer",
			"window",
			"door",
			"bed",
		},
		Exits: map[string]string{
			"east": "staircase",
		},
	},
	"staircase": {
		ID:                  "staircase",
		Name:                "Staircase",
		Description:         "You stand on the small landing at the top of the stairs. The door back into your apartment is to the [west]. A staircase leads [down] to the street",
		FirstVisitNarration: "this is a test of the first visit narration!",
		Exits: map[string]string{
			"west": "apartment",
			"down": "street",
		},
	},
	"street": {
		ID:          "street",
		Name:        "Street",
		Description: "The street outside your apartment is quiet. You see a man standing on the corner.",
		Objects: []string{
			"street_man",
		},
		Exits: map[string]string{
			"west": "staircase",
			"work": "work_main",
		},
	},
	"work_main": {
		ID:                  "work_main",
		Name:                "Security Office",
		Description:         "You are in your small security office that has one computer with a glowing screen. To the [west] is the breakroom. On the [north] wall is a security door to the rest of the building.",
		FirstVisitNarration: "You arrive to work at Learnex, an old software company. You work overnight security. You chose this job becuase it is one of the last places in the city that still uses older tech that requires an actual human to operate it.",
		Objects: []string{
			"work_computer",
		},
		Exits: map[string]string{
			"west":  "work_breakroom",
			"north": "hallway",
		},
	},
	"work_breakroom": {
		ID:          "work_breakroom",
		Name:        "Break Room",
		Description: "A small breakroom.",
		Objects: []string{
			"work_computer",
			"security_door",
		},
		Exits: map[string]string{
			"east": "work_main",
		},
	},
}
