package day1

import (
	"testing"
)

func getLines() []string {
	return []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"}
}

func TestDay1Part1(t *testing.T) {
	lines := getLines()
	got := Part1(lines)
	want := 11
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
