package world

// helper for seeing into parent containers

func findObjectRecursive(obj Object, target string) (Object, bool) {
	for _, childID := range obj.Contains {

		child := Objects[childID]

		if child.ID == target {
			return child, true
		}

		if child.Container {
			if !child.Openable || child.Open {

				found, ok := findObjectRecursive(child, target)

				if ok {
					return found, true
				}
			}
		}
	}

	return Object{}, false
}

// public func

func FindVisibleObject(room Room, target string) (Object, bool) {
	for _, objID := range room.Objects {

		obj := Objects[objID]

		if obj.ID == target {
			return obj, true
		}

		if obj.Container {
			if !obj.Openable || obj.Open {

				found, ok := findObjectRecursive(obj, target)

				if ok {
					return found, true
				}
			}
		}
	}

	return Object{}, false
}
