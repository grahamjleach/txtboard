package txtboard

import "errors"

var placementConflictError = errors.New("placement conflict")

type direction int

const (
	directionNone direction = iota
	DirectionNorth
  DirectionSouth
  DirectionEast
  DirectionWest
)

var stringToDirection = map[string]direction{
	"north": DirectionNorth,
	"south": DirectionSouth,
	"east": DirectionEast,
	"west": DirectionWest,
}

func directionFromString(s string) direction {
	return stringToDirection[s]
}

// type Object struct {
// 	description string
// 	isInventoryable bool
// 	isAddressable bool
// 	state map[string]bool
// }

type position struct {
	x int
	y int
}

type Character interface {
  Describe() string
}

type Object interface {
  Describe() string
}

type Player interface {

}

type gameError struct {
  message string
}

func (e *gameError) Error() string {
  return e.message
}
