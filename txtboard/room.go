package txtboard

type Room struct {
	*Board
	id string
  description string
  objects []*Object
  obstacles map[direction]*Object
}
