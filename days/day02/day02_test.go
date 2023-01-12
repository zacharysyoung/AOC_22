package day2

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {
	for _, tc := range []struct {
		them, you string
		score     score
		res       outcome
	}{
		{"A", "Y", 8, win},
		{"B", "X", 1, lose},
		{"C", "Z", 6, draw},
	} {
		score, res := throwDown(tc.them, tc.you)
		if score != tc.score || res != tc.res {
			t.Errorf("throwDown(%s, %s) = %d, %d; want %d, %d", tc.them, tc.you, score, res, tc.score, tc.res)
		}
	}

	var totalScore score
	for _, match := range getMatches(input) {
		myScore, _ := throwDown(match.them, match.you)
		totalScore += myScore
	}

	fmt.Println(totalScore)
}
