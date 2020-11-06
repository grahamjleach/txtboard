package main

import (
  "github.com/grahamjleach/txtboard/txtboard"
)

func main() {
  board := txtboard.NewBoard()
  _ = board.AddRoom("a dark room", 0, 0)
  _ = board.AddRoom("a light room", 0, 1)

  board.Begin(0, 0)
}
