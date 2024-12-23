package day3

import "testing"

const input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
const input2 = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestDay3Part1(t *testing.T) {
	got := Part1(input)
	want := 161
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDay3Part2(t *testing.T) {
	got := Part2(input2)
	want := 48
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
