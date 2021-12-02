package main

import (
	"fmt"
	"os"

	"github.com/kevinbarbour/advent-of-code-2021/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide an input file")
		os.Exit(1)
	}

	answer := CountIncreases(parser.ParseToInts(os.Args[1]))
	fmt.Printf("There are %d increases in the input.\n", answer)

	answer = CountWindowIncreases(parser.ParseToInts(os.Args[1]))
	fmt.Printf("There are %d window increases in the input.\n", answer)
}

func CountIncreases(depths []int) (increases int) {
	for i := range depths {
		if i > 0 && depths[i] > depths[i-1] {
			increases++
		}
	}

	return
}

func CountWindowIncreases(depths []int) (increases int) {
	prev := 0
	for i := 2; i < len(depths); i++ {
		curr := depths[i] + depths[i-1] + depths[i-2]

		if i > 2 && curr > prev {
			increases++
		}
		prev = curr
	}
	return
}
