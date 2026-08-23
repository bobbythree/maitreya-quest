package dialogue

var Dialogues = map[string]Dialogue{
	"street_man": {
		ID:        "street_man",
		StartNode: "prophecy",
		Nodes: map[string]Node{
			"prophecy": {
				Speaker: "Man",
				Text:    "The prophecy is true...",
				Choices: []Choice{
					{
						Text:     "What prophecy?",
						NextNode: "explanation",
					},
					{
						Text:     "You've got the wrong person.",
						NextNode: "denial",
					},
					{
						Text:     "Walk away.",
						NextNode: "",
					},
				},
			},

			"explanation": {
				Speaker: "Man",
				Text:    "The prophecy of Maitreya and the otherworldly being.",
				Choices: []Choice{
					{
						Text:     "Who is Maitreya?",
						NextNode: "maitreya",
					},
					{
						Text:     "I've heard enough.",
						NextNode: "",
					},
				},
			},

			"denial": {
				Speaker: "Man",
				Text:    "No. I don't think I do.",
				Choices: []Choice{
					{
						Text:     "Ok, what prophecy?",
						NextNode: "explanation",
					},
					{
						Text:     "Walk away.",
						NextNode: "",
					},
				},
			},

			"maitreya": {
				Speaker: "Man",
				Text:    "That is something you're going to have to figure out for yourself.",
				Choices: []Choice{
					{
						Text:     "Walk away.",
						NextNode: "",
					},
				},
			},
		},
	},
}
