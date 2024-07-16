package main

import (
	"fmt"

	"github.com/sagnikc395/aoc2015/utils"
)

type Direction struct {
	X int
	Y int
}

func (d *Direction) UpdateDirection(direction string) {
	switch direction {
	case ">":
		d.X += 1
	case "^":
		d.Y += 1
	case "<":
		d.X -= 1
	case "v":
		d.Y -= 1
	}

}

func main() {
	f := utils.ReadFile("./input.txt")

	data := string(f)

	//start posn at 0,0
	var dir Direction

	visitedHouses := make(map[Direction]bool)

	//marking the starting house as visited
	visitedHouses[dir] = true

	for _, val := range data {
		if val == '>' {
			dir.UpdateDirection(">")
		} else if val == '<' {
			dir.UpdateDirection("<")
		} else if val == 'v' {
			dir.UpdateDirection("v")
		} else if val == '^' {
			dir.UpdateDirection("^")
		}
		visitedHouses[dir] = true
	}
	fmt.Printf("Number of houses that received at least one present: %d\n", len(visitedHouses))
	fmt.Printf("Final Position (%d ,%d)", dir.X, dir.Y)
}
