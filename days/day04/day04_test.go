package day03

import (
	"fmt"
	"reflect"
	"testing"
)

const sampleInput = `
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

var sampleAssignmentPairs = []assignmentPair{
	{assignment{2, 4}, assignment{6, 8}},
	{assignment{2, 3}, assignment{4, 5}},
	{assignment{5, 7}, assignment{7, 9}},
	{assignment{2, 8}, assignment{3, 7}},
	{assignment{6, 6}, assignment{4, 6}},
	{assignment{2, 6}, assignment{4, 8}},
}

func TestReadInput(t *testing.T) {
	pairs := readInput(sampleInput)
	if !reflect.DeepEqual(pairs, sampleAssignmentPairs) {
		t.Errorf("got %v; want %v", pairs, sampleAssignmentPairs)
	}
}

var sampleTotalOverlaps = []assignmentPair{
	{assignment{2, 8}, assignment{3, 7}},
	{assignment{6, 6}, assignment{4, 6}},
}

func TestSampleTotalOverlaps(t *testing.T) {
	pairs := readInput(sampleInput)
	overlaps := findTotalOverlaps(pairs)
	if !reflect.DeepEqual(overlaps, sampleTotalOverlaps) {
		t.Errorf("got overlapping pairs %v; want %v", overlaps, sampleTotalOverlaps)
	}
}

const sampleTotalOverlapCount = 2

func TestGetSampleTotalCount(t *testing.T) {
	pairs := readInput(sampleInput)
	overlaps := findTotalOverlaps(pairs)
	got := getOverlapCount(overlaps)
	if got != sampleTotalOverlapCount {
		t.Errorf("got overlap count of %d; want %d", got, sampleTotalOverlapCount)
	}
}

func TestGetRealAnswer_1(t *testing.T) {
	pairs := readInput(input)
	overlaps := findTotalOverlaps(pairs)
	fmt.Println("Real answer 1:", getOverlapCount(overlaps))
}

var samplePartialOverlaps = []assignmentPair{
	{assignment{5, 7}, assignment{7, 9}},
	{assignment{2, 8}, assignment{3, 7}},
	{assignment{6, 6}, assignment{4, 6}},
	{assignment{2, 6}, assignment{4, 8}},
}

func TestSamplePartialOverlaps(t *testing.T) {
	pairs := readInput(sampleInput)
	overlaps := findPartialOverlaps(pairs)
	if !reflect.DeepEqual(overlaps, samplePartialOverlaps) {
		t.Errorf("got overlapping pairs %v; want %v", overlaps, samplePartialOverlaps)
	}
}

const samplePartialOverlapCount = 4

func TestGetSamplePartialCount(t *testing.T) {
	pairs := readInput(sampleInput)
	overlaps := findPartialOverlaps(pairs)
	got := getOverlapCount(overlaps)
	if got != samplePartialOverlapCount {
		t.Errorf("got overlap count of %d; want %d", got, samplePartialOverlapCount)
	}
}
func TestGetRealAnswer_2(t *testing.T) {
	pairs := readInput(input)
	overlaps := findPartialOverlaps(pairs)
	fmt.Println("Real answer 2:", getOverlapCount(overlaps))
}

/*
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
*/
