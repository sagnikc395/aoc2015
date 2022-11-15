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
	result := string(content)
	currentFloor := 0
	for index, value := range result {
		if value == '(' {
			currentFloor += 1
		} else if value == ')' {
			currentFloor -= 1
		}
		if currentFloor == -1 {
			fmt.Println(index + 1)
			break
		}
	}
}
