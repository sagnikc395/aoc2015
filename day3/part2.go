package main

import (
	"fmt"

	"github.com/sagnikc395/aoc2015/utils"
)

type Direction2 struct {
	X int
	Y int
}

func (d *Direction2) UpdateDirection2(direction string) {
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
	f := utils.ReadFile("./input2.txt")

	data := string(f)

	//start posn at 0,0
	var santaDir Direction2

	var robotDirection Direction2

	visitedHouses := make(map[Direction2]bool)

	//marking the starting house as visited
	visitedHouses[santaDir] = true
	//mark the houses visited by robot
	visitedHouses[robotDirection] = true

	for index, val := range data {
		//since taking turns
		if index%2 == 0 {
			//this will be updated by santaDir
			switch val {
			case '<':
				santaDir.UpdateDirection2("<")
			case '>':
				santaDir.UpdateDirection2(">")
			case 'v':
				santaDir.UpdateDirection2("v")
			case '^':
				santaDir.UpdateDirection2("^")
			}
			visitedHouses[santaDir] = true
		} else {
			//this will be updated by robotDir
			switch val {
			case '<':
				santaDir.UpdateDirection2("<")
			case '>':
				santaDir.UpdateDirection2(">")
			case 'v':
				santaDir.UpdateDirection2("v")
			case '^':
				santaDir.UpdateDirection2("^")
			}
			visitedHouses[robotDirection] = true
		}
	}
	fmt.Printf("Number of houses that received at least one present: %d\n", len(visitedHouses))
	fmt.Printf("Final Position (Santa -> (%d ,%d) , Robot -> (%d, %d))", santaDir.X, santaDir.Y, robotDirection.X, robotDirection.Y)
}
