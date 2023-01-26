package day03

import (
	"fmt"
	"reflect"
	"testing"
)

const sampleInput = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

var sampleStacks = stacks{
	1: {"Z", "N"},
	2: {"M", "C", "D"},
	3: {"P"},
}

var sampleMoves = moves{
	{1, 2, 1},
	{3, 1, 3},
	{2, 2, 1},
	{1, 1, 2},
}

func TestReadInput(t *testing.T) {
	stacks, moves := readInput(sampleInput)
	fmt.Println(stacks)
	if !reflect.DeepEqual(stacks, sampleStacks) {
		t.Errorf("got stacks %v; want %v", stacks, sampleStacks)
	}
	if !reflect.DeepEqual(moves, sampleMoves) {
		t.Errorf("got moves %v; want %v", moves, sampleMoves)
	}
}

/*
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
*/
