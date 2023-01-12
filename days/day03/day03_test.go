package day03

import (
	"fmt"
	"reflect"
	"testing"
)

const sampleInput = `
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

var sampleRucksacks = []rucksack{
	{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
	{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
	{"PmmdzqPrV", "vPwwTWBwg"},
}

func TestReadInput(t *testing.T) {
	rucksacks := readInput(sampleInput)[:3]
	if !reflect.DeepEqual(rucksacks, sampleRucksacks) {
		t.Errorf("got %v; want %v", rucksacks, sampleRucksacks)
	}

}

var sampleDuplicates = []item{
	{"p", 16},
	{"L", 38},
	{"P", 42},
	{"v", 22},
	{"t", 20},
	{"s", 19},
}

func TestMakeItem(t *testing.T) {
	testItem := func(want item) {
		l := want.letter
		r := []rune(l)[0]
		got := makeItem(r)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("makeItem(%q) = %v; want %v", r, got, want)
		}
	}

	for _, boundaryItem := range []item{
		{"a", 1},
		{"z", 26},
		{"A", 27},
		{"Z", 52},
	} {
		testItem(boundaryItem)
	}

	for _, sampleItem := range sampleDuplicates {
		testItem(sampleItem)
	}

}

const samplePrioritySum = 157

func TestSampleDuplicates(t *testing.T) {
	rucksacks := readInput(sampleInput)
	dupeItems := make([]item, len(rucksacks))
	for i, rucksack := range rucksacks {
		dupeItem := findDupeItem(rucksack)
		dupeItems[i] = dupeItem
	}

	if !reflect.DeepEqual(dupeItems, sampleDuplicates) {
		t.Errorf("got duplicate items %v; want %v", dupeItems, sampleDuplicates)
	}

	got := getPrioritySum(dupeItems)

	if got != samplePrioritySum {
		t.Errorf("got priority sum %d; want %d", got, samplePrioritySum)
	}
}

func TestGetRealAnswer_1(t *testing.T) {
	rucksacks := readInput(input)
	dupeItems := make([]item, len(rucksacks))
	for i, rucksack := range rucksacks {
		dupeItem := findDupeItem(rucksack)
		dupeItems[i] = dupeItem
	}

	fmt.Println("Real answer 1:", getPrioritySum(dupeItems))
}

func TestMakeGroups(t *testing.T) {
	rucksacks := readInput(sampleInput)
	for _, elfGroup := range makeGroups(rucksacks) {
		if len(elfGroup) != 3 {
			t.Errorf("expected elf group to only be 3 big; instead it's %d big", len(elfGroup))
		}
	}
}

var sampleBadges = []item{
	{"r", 18},
	{"Z", 52},
}

const sampleBadgePrioritySum = 70

func TestFindGroupBadges(t *testing.T) {
	rucksacks := readInput(sampleInput)
	groups := makeGroups(rucksacks)

	badges := getGroupBadges(groups)

	if !reflect.DeepEqual(badges, sampleBadges) {
		t.Errorf("got sample group badges %v; want %v", badges, sampleBadges)
	}

	sum := getPrioritySum(badges)

	if sum != sampleBadgePrioritySum {
		t.Errorf("got sample priority sum %d; want %d", sum, sampleBadgePrioritySum)
	}
}

func TestGetRealAnswer_2(t *testing.T) {
	rucksacks := readInput(input)
	groups := makeGroups(rucksacks)
	badges := getGroupBadges(groups)

	fmt.Println("Real answer 1:", getPrioritySum(badges))
}
