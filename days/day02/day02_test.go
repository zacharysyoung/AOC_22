package day2

import (
	"fmt"
	"reflect"
	"testing"
)

const sampleInput = `
A Y
B X
C Z
`

func TestReadGuide(t *testing.T) {
	want := []strategy{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}
	got := readGuide(sampleInput)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("getMatches(%q) = %v; want %v", input, got, want)
	}
}

func TestThrowDown_1(t *testing.T) {
	for _, tc := range []struct {
		x    strategy
		want score
	}{
		{strategy{"A", "Y"}, 8},
		{strategy{"B", "X"}, 1},
		{strategy{"C", "Z"}, 6},
	} {
		if got := throwStrategy_1(tc.x); got != tc.want {
			t.Errorf("throwDown(%+v) = %d; want %d", tc.x, got, tc.want)
		}
	}
}

func TestPlayGuide_1(t *testing.T) {
	guide := readGuide(sampleInput)

	want := score(15)
	got := playGuide(guide, throwStrategy_1)

	if got != want {
		t.Errorf("playGuide_1 = %d; want %d", got, want)
	}

	// Get value for real input
	guide = readGuide(input)
	got = playGuide(guide, throwStrategy_1)
	fmt.Println("playGuide_1 real score:", got)
}

func TestPlayGuide_2(t *testing.T) {
	guide := readGuide(sampleInput)

	want := score(12)
	got := playGuide(guide, throwStrategy_2)

	if got != want {
		t.Errorf("playGuide_2 = %d; want %d", got, want)
	}

	// Get value for real input
	guide = readGuide(input)
	got = playGuide(guide, throwStrategy_2)
	fmt.Println("playGuide_2 real score:", got)
}
