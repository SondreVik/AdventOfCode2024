package main

import (
	"AdventOfCode2024/day1"
	"AdventOfCode2024/day2"
	"AdventOfCode2024/day3"
	"AdventOfCode2024/day4"
	"fmt"
	"os"
	"strings"
)

func main() {
	arg := os.Args[1]
	parts := strings.Split(arg, "_")
	dayArg := parts[0]
	partArg := parts[1]
	if dayArg == "1" {
		if partArg == "1" {
			day1.SolvePart1()
		} else {
			day2.SolvePart2()
		}
		return
	}
	if dayArg == "2" {
		if partArg == "1" {
			day2.SolvePart1()
		} else {
			day2.SolvePart2()
		}
		return
	}
	if dayArg == "3" {
		if partArg == "1" {
			day3.SolvePart1()
		} else {
			day3.SolvePart2()
		}
		return
	}
	if dayArg == "4" {
		if partArg == "1" {
			day4.SolvePart1()
		} else {
			day4.SolvePart2()
		}
	}
	fmt.Println("Not solved yet...")
}
