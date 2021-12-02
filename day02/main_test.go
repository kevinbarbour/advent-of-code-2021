package main

import (
	"testing"

	"github.com/kevinbarbour/advent-of-code-2021/internal/parser"
)

func Test_CalculatePosition(t *testing.T) {
	data := parser.ParseToStrings("input/test")
	want := 150
	got := CalculatePosition(data)

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}

func Test_CalculatePositionWithAim(t *testing.T) {
	data := parser.ParseToStrings("input/test")
	want := 900
	got := CalculatePositionWithAim(data)

	if want != got {
		t.Fatalf("want %d, got %d", want, got)
	}
}
