package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/kevinbarbour/advent-of-code-2021/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input file")
		os.Exit(1)
	}

	answer := CalculatePosition(parser.ParseToStrings(os.Args[1]))
	fmt.Printf("The answer is %d\n", answer)

	answer = CalculatePositionWithAim(parser.ParseToStrings(os.Args[1]))
	fmt.Printf("The answer with aim is %d\n", answer)
}

func CalculatePositionWithAim(movements []string) int {
	x, y, aim := 0, 0, 0

	for _, m := range movements {
		direction, value := parseMovement(m)
		switch direction {
		case "forward":
			x += value
			y += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	return x * y
}

func CalculatePosition(movements []string) int {
	x, y := 0, 0

	for _, m := range movements {
		direction, value := parseMovement(m)
		switch direction {
		case "forward":
			x += value
		case "down":
			y += value
		case "up":
			y -= value
		}
	}

	return x * y
}

func parseMovement(movement string) (direction string, value int) {
	split := strings.Split(movement, " ")
	direction = split[0]
	value, _ = strconv.Atoi(split[1])

	return
}
