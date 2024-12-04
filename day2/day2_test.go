package day2

import (
	"strings"
	"testing"
)

func getLines() []string {
	text := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`
	return strings.Split(text, "\n")
}

func TestDay2Part1(t *testing.T) {
	lines := getLines()
	got := Part1(lines)
	want := 2
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDay1Part2(t *testing.T) {
	lines := getLines()
	got := Part2(lines)
	want := 31
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
