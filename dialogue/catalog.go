package dialogue

var Dialogues = map[string]Dialogue{
	"street_man": {
		ID:        "street_man",
		StartNode: "prophecy",
		Completion: &Completion{
			RequiredNodes: []string{
				"explanation",
				"maitreya",
				"otherworldly",
			},
			OnComplete: unlockWork,
		},
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
						Text:     "Who are you?",
						NextNode: "introduce",
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
				Text:    "You haven't heard the prophecy? Maitreya will save us all. From AI and from ourselves. He lived 100 years ago. Some say before he died he embedded his mind in an adventure game. Those of us who still care about the human condition are searching for that game.",
				Choices: []Choice{
					{
						Text:     "And the otherworldly being?",
						NextNode: "otherworldly",
					},
					{
						Text:     "Who are you?",
						NextNode: "introduce",
					},
					{
						Text:     "Walk away.",
						NextNode: "",
					},
				},
			},
			"introduce": {
				Speaker: "Man",
				Text:    "My name is Silas Lundvort.",
				Choices: []Choice{
					{
						Text:     "What prophecy?",
						NextNode: "explanation",
					},
					{
						Text:     "Walk away.",
						NextNode: "",
					},
				},
			},
			"otherworldly": {
				Speaker: "Man",
				Text:    "That's the part nobody can figure out. The prophecy says that together they will defeat the AI and restore us to sanity",
				Choices: []Choice{
					{
						Text:     "What prophecy?",
						NextNode: "explanation",
					},
					{
						Text:     "Walk away.",
						NextNode: "",
					},
				},
			},
		},
	},
}
