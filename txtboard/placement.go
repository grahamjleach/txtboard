package txtboard

type RoomOption func(*Room) error

func PlaceObstacle(object *Object, d direction) RoomOption {
	return func(room *Room) error {
		if _, found := room.obstacles[d]; found {
			return placementConflictError
		}

		room.obstacles[d] = object;

		return nil
	}
}

func PlaceObject(object *Object) RoomOption {
	return func(room *Room) error {
		// if _, found := room.obstacles[d]; found {
		// 	return placementConflictError
		// }
		//
		// room.objects = append(room.objects, object);

		return nil
	}
}
