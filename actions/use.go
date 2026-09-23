package actions

import (
	"fmt"

	"github.com/bobbythree/maitreya-quest/game"
	"github.com/bobbythree/maitreya-quest/oldcomputer"
	"github.com/bobbythree/maitreya-quest/parser"
	"github.com/bobbythree/maitreya-quest/world"
)

func Use(gs *game.GameState, cmd parser.Command) string {
	directObject := cmd.DirectObject

	indirectObject := cmd.IndirectObject

	if directObject == "" {
		return "Use what?"
	}

	objA, ok := world.FindVisibleObject(gs, directObject)

	if !ok {
		return fmt.Sprintf("You don't see a %s.", directObject)
	}

	// standalone 'use' case
	if indirectObject == "" {
		if objA.UseAction != nil {
			return objA.UseAction(gs)
		}

		return "Use it with what?"
	}

	if cmd.Preposition != "with" && cmd.Preposition != "in" && cmd.Preposition != "on" {
		return "You can't do that."
	}

	objB, ok := world.FindVisibleObject(gs, indirectObject)

	if !ok {
		return fmt.Sprintf("You don't see a %s.", indirectObject)
	}

	isOldComputerPair := objA.ID == "thumbdrive" && objB.ID == "old_computer" ||
		objA.ID == "old_computer" && objB.ID == "thumbdrive"
	if isOldComputerPair {
		return oldcomputer.Start(gs)
	}

	isAdapterThumbdrivePair := objA.ID == "adapter" && objB.ID == "thumbdrive" ||
		objA.ID == "thumbdrive" && objB.ID == "adapter"
	if isAdapterThumbdrivePair {
		adapter := gs.ObjectStates["adapter"]
		thumbdrive := gs.ObjectStates["thumbdrive"]
		if thumbdrive.Parent == "adapter" {
			return "The thumbdrive is already connected to the adapter."
		}
		// Keep the combined item in the adapter's existing inventory slot.
		if adapter.Parent != "inventory" || thumbdrive.Parent != "inventory" {
			return "You need to take both items first."
		}
		for i, id := range gs.Player.Inventory {
			if id == "thumbdrive" {
				gs.Player.Inventory = append(gs.Player.Inventory[:i], gs.Player.Inventory[i+1:]...)
				break
			}
		}
		thumbdrive.Parent = "adapter"
		adapter.Contains = append(adapter.Contains, "thumbdrive")
		return "You connect the thumbdrive to the adapter."
	}

	isAdapterHomePair := objA.ID == "adapter" && objB.ID == "home_computer" ||
		objA.ID == "home_computer" && objB.ID == "adapter"
	if isAdapterHomePair {
		if gs.ObjectStates["thumbdrive"].Parent != "adapter" {
			return "that doesn't quite work....we're missing something"
		}
		if !gs.BeginOldComputer(game.OldComputerState{Screen: "password", Computer: "home"}) {
			return "You're already busy."
		}
		return "The computer asks for a password."
	}

	if objA.ID == "thumbdrive" && objB.ID == "work_computer" {
		return "this computer is pretty old, but it's not THAT old!"
	}

	if objA.ID == "thumbdrive" && objB.ID == "home_computer" {
		return "This thumbdrive doesn't fit in the computer's port. You'll have to find a wayyy older computer."
	}

	// Keep "on" limited to combinations that explicitly support it.
	if cmd.Preposition == "on" {
		return "You can't do that."
	}

	return "Nothing happens."
}
