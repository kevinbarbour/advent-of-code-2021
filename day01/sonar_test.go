package main

import (
	"testing"

	"github.com/kevinbarbour/advent-of-code-2021/internal/parser"
)

func Test_CountIncreases(t *testing.T) {
	data := parser.ParseToInts("input/test")
	got := CountIncreases(data)
	want := 7

	if want != got {
		t.Fatalf("Want %d, got %d", want, got)
	}
}

func Test_CountWindowIncreases(t *testing.T) {
	data := parser.ParseToInts("input/test")
	got := CountWindowIncreases(data)
	want := 5

	if want != got {
		t.Fatalf("Want %d, got %d", want, got)
	}
}
