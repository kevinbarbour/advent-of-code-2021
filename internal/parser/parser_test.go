package parser

import (
	"reflect"
	"testing"
)

func Test_ParseToInts(t *testing.T) {
	filename := "input/test"
	got := ParseToInts(filename)
	want := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("want %v got %v", want, got)
	}
}
