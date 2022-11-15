package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	pattern := string(content)
	currentFloor := 0

	for _, value := range pattern {
		if value == '(' {
			currentFloor += 1
		} else if value == ')' {
			currentFloor -= 1
		}
	}
	fmt.Println(currentFloor)

}
