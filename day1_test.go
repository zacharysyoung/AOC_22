package day1

import (
	"reflect"
	"testing"
)

var getAllCalsTestCases = []struct {
	input string
	want  calsList
}{
	{`
1`, calsList{1}},
	{`
1
2`, calsList{3}},
	{`
1
2

3
4`, calsList{3, 7}},
	{`
1
2


3
4

5
6
7


`, calsList{3, 7, 18}},
}

func TestGetAllCals(t *testing.T) {
	for _, tc := range getAllCalsTestCases {
		if got := getAllCals(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("getAllCals(%+q) = %v; want %v\n", tc.input, got, tc.want)
		}
	}
}

var inputCals = getAllCals(input)

const part1_want = 69289

func TestPart1(t *testing.T) {
	for _, tc := range []struct {
		allCals calsList
		want    int
	}{
		{calsList{0, 1, 2}, 2},
		{calsList{2, 1, 0}, 2},
		{calsList{1, 1, 2}, 2},
	} {
		if got := part1(tc.allCals); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("part1(%v) = %d; want %d\n", tc.allCals, got, tc.want)
		}

	}

	if got := part1(inputCals); got != part1_want {
		t.Errorf("part1(%v) = %d; want %d\n", inputCals, got, part1_want)
	}
}

const part2_want = 205615

func TestPart2_1(t *testing.T) {
	if got := part2_1(inputCals); got != part2_want {
		t.Errorf("part2_1() = %d; want %d", got, part2_want)
	}
}
func TestPart2_2(t *testing.T) {
	if got := part2_2(inputCals); got != part2_want {
		t.Errorf("part2_2() = %d; want %d", got, part2_want)
	}
}

func TestPart2_3(t *testing.T) {
	for _, tc := range []struct {
		allCals calsList
		want    int
	}{
		{calsList{1, 2, 3}, 6},
		{calsList{2, 3, 4}, 9},
	} {
		if got := part2_3(tc.allCals); got != tc.want {
			t.Errorf("part2_3(%v) = %d; want %d", tc.allCals, got, tc.want)
		}
	}

	if got := part2_3(inputCals); got != part2_want {
		t.Errorf("part2_3(inputCals) = %d; want %d", got, part2_want)
	}
}

func BenchmarkPart2(b *testing.B) {
	b.Run("Part2_1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part2_1(inputCals)
		}
	})
	b.Run("Part2_2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part2_2(inputCals)
		}
	})
	b.Run("Part2_3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			part2_3(inputCals)
		}
	})
}

func TestShift(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		idx  int
		want []int
	}{
		{[]int{0, 1, 2, 3}, 0, []int{0, 0, 1, 2}},
		{[]int{0, 1, 2, 3}, 1, []int{0, 1, 1, 2}},
		{[]int{0, 1, 2, 3}, 2, []int{0, 1, 2, 2}},
	} {
		cpy := make([]int, len(tc.s))
		copy(cpy, tc.s)
		shift(cpy, tc.idx)
		if !reflect.DeepEqual(cpy, tc.want) {
			t.Errorf("shift(%v, %d) = %v; want %v\n", tc.s, tc.idx, cpy, tc.want)
		}
	}
}
