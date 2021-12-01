package main

import (
	"reflect"
	"testing"
)

func Test_ReadInput(t *testing.T) {
	filename := "input/test"
	got := ReadInput(filename)
	want := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v got %v", want, got)
	}
}

func Test_CountIncreases(t *testing.T) {
	data := ReadInput("input/test")
	got := CountIncreases(data)
	want := 7

	if want != got {
		t.Fatalf("Want %d, got %d", want, got)
	}
}

func Test_CountWindowIncreases(t *testing.T) {
	data := ReadInput("input/test")
	got := CountWindowIncreases(data)
	want := 5

	if want != got {
		t.Fatalf("Want %d, got %d", want, got)
	}
}
