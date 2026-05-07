package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Get(gs *game.GameState, cmd parser.Command) {
	room := world.Rooms[gs.CurrentRoom]
	noun := cmd.Noun

	for _, objID := range room.Objects {

		obj := world.Objects[objID]

		if obj.ID == noun && obj.Portable {

			gs.Inventory = append(gs.Inventory, noun)

			fmt.Println("Taken.")

			return
		}

		if obj.Container {
			for i, insideID := range obj.Contains {
				if insideID == noun {

					item := world.Objects[insideID]

					if !item.Portable {
						fmt.Println("You can't take that.")
						return
					}

					gs.Inventory = append(gs.Inventory, noun)

					obj.Contains = append(
						obj.Contains[:i],
						obj.Contains[i+1:]...,
					)

					world.Objects[objID] = obj

					fmt.Println("Taken.")

					return
				}
			}
		}
	}

	fmt.Println("You don't see that.")
}
