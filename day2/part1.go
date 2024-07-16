package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sagnikc395/aoc2015.git/utils"
)

func smallestSideArea(l, b, h int) int {
	return min(l*b, b*h, h*l)
}

func surfaceArea(l, b, h int) int {
	return 2*l*b + 2*b*h + 2*h*l
}

func main() {
	f := utils.ReadFile("./input.txt")

	data := strings.Split(string(f), "\n")

	totalCount := 0
	for _, item := range data {
		dims := strings.Split(item, "x")
		l, _ := strconv.Atoi(dims[0])
		b, _ := strconv.Atoi(dims[1])
		h, _ := strconv.Atoi(dims[2])
		surface_area := surfaceArea(l, b, h)
		smallestSide_area := smallestSideArea(l, b, h)
		totalCount += surface_area + smallestSide_area
	}
	fmt.Println(totalCount)
}
