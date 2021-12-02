package parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseToInts(filename string) []int {
	split := ParseToStrings(filename)

	var numbers []int
	for _, s := range split {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}
	return numbers
}

func ParseToStrings(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Invalid file name %s", filename)
		os.Exit(1)
	}

	data := string(bytes)
	split := strings.Split(data, "\n")
	return split
}
