package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sagnikc395/aoc2015/utils"
)

func presentVolume(l, b, h int) int {
	return l * b * h
}

func smallestPerimeter(l, b, h int) int {
	peri1 := 2 * (l + b)
	peri2 := 2 * (b + h)
	peri3 := 2 * (l + h)

	return min(peri1, peri2, peri3)
}

func main() {
	f := utils.ReadFile("./input.txt")

	data := strings.Split(string(f), "\n")

	total := 0
	for _, val := range data {
		items := strings.Split(val, "x")
		l, _ := strconv.Atoi(items[0])
		b, _ := strconv.Atoi(items[1])
		h, _ := strconv.Atoi(items[2])
		total += smallestPerimeter(l, b, h) + presentVolume(l, b, h)
	}

	fmt.Println(total)
}
