package txtboard

import (
	"bufio"
	"fmt"
	"errors"
	"os"
	"os/exec"
	"strings"
)

type Board struct {
	rooms map[position]*Room
	currentPosition position
	reader *bufio.Reader
}

func NewBoard() *Board {
	return &Board{
		rooms: make(map[position]*Room),
	}
}

func (b *Board) AddRoom(description string, x, y int, options ...RoomOption) error {
	p := position{
		x: x,
		y: y,
	}

	if _, found := b.rooms[p]; found {
		return errors.New("blap")
	}

	room := &Room{
		Board: b,
	  description: description,
	 	obstacles: make(map[direction]*Object),
	}

	for _, option := range options {
		if err := option(room); err != nil {
			return err
		}
	}

	b.rooms[p] = room

	return nil
}

func (b *Board) narrate(text string) {
	fmt.Println(text)

	exec.Command("say", text).Run()
}

func (b *Board) Begin(x, y int) {
	// if active

	b.reader = bufio.NewReader(os.Stdin)

	b.currentPosition = position{
		x: x,
		y: y,
	}
	//
	// room, found := b.rooms[p]
	// if !found {
	// 	panic("no room")
	// }
	//
	// b.currentRoom = room

	b.play()
}

var errorBadCommand = errors.New("i don't understand")

func (b *Board) play() {
	Game:
	for {
		cp := b.currentPosition

		room, found := b.rooms[cp]
		if !found {
			panic("no room")
		}

		b.narrate(fmt.Sprintf("you enter %s\n", room.description))

		for cp == b.currentPosition {
			var i input

			PlayerPrompt:
			for {
				i = b.prompt()
				switch {
				case i.exit():
					break Game
				case i.valid():
					break PlayerPrompt
				default:
					b.narrate(errorBadCommand.Error())
				}
			}

			switch {
			case i.isMovement():
				if err := b.move(i.direction()); err != nil {
					b.narrate(err.Error())
				}
			default:
				b.narrate("you do dat")
			}
		}
	}
}

func (b *Board) move(d direction) error {
	var nextPosition position

	switch d {
		case DirectionNorth:
			nextPosition = position{
				x: b.currentPosition.x,
				y: b.currentPosition.y+1,
			}
		case DirectionSouth:
			nextPosition = position{
				x: b.currentPosition.x,
				y: b.currentPosition.y-1,
			}
		case DirectionEast:
			nextPosition = position{
				x: b.currentPosition.x+1,
				y: b.currentPosition.y,
			}
		case DirectionWest:
			nextPosition = position{
				x: b.currentPosition.x-1,
				y: b.currentPosition.y,
			}
		default:
			return errorBadCommand
	}

	if _, found := b.rooms[nextPosition]; !found {
			return errors.New("the way is blocked")
	}

	b.currentPosition = nextPosition

	return nil
}

func  (b *Board) prompt() input {
	fmt.Print("> ")
	c, _ := b.reader.ReadString('\n')

	// if noAlph {
	//
	// }

	p := clean(strings.Split(c, " "))

	i := input{}

	switch len(p) {
	case 0:
	case 1:
		i.action = p[0]
	case 2:
		i.action = p[0]
		i.object = p[1]
	default:
		i.action = p[0]
		i.object = p[1]
		i.parameters = p[2:]
	}

	return i

}
