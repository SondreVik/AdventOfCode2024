package day2

import (
	"AdventOfCode2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func SolvePart1() {
	fmt.Println(Part1(utils.ReadFile("day2/data.txt")))
}

func SolvePart2() {
	fmt.Println(Part2(utils.ReadFile("day2/data.txt")))
}

func Part1(input []string) int {
	safeReports := 0
	for _, line := range input {
		report := getReport(line)
		if isSafe(report, 0) {
			safeReports = safeReports + 1
		}
	}
	return safeReports
}

func Part2(input []string) int {
	safeReports := 0
	for _, line := range input {
		report := getReport(line)
		if isSafe(report, 1) {
			safeReports = safeReports + 1
		}
	}
	return safeReports
}

func isSafe(report []int, tolerance int) bool {
	if tolerance == -1 {
		return false
	}
	ascending := false
	last := 0
	for key, el := range report {
		if key == 0 {
			last = el
			continue
		}
		newDelta := el - last
		newAscending := isAscending(newDelta)
		if !isDeltaValid(newDelta) || (key > 1 && ascending != newAscending) {
			part1 := utils.RemoveIndex(report, key-1)
			part2 := utils.RemoveIndex(report, key)
			return isSafe(part1, tolerance-1) || isSafe(part2, tolerance-1)
		}
		ascending = newAscending
		last = el
	}
	return true
}

func isAscending(diff int) bool {
	return diff > 0
}

func isDeltaValid(delta int) bool {
	var diff int
	if delta < 0 {
		diff = -delta
	} else {
		diff = delta
	}
	if diff < 1 || diff > 3 {
		return false
	}
	return true
}

func getReport(input string) (report []int) {
	for _, el := range strings.Split(input, " ") {
		val, err := strconv.Atoi(el)
		if err != nil {
			panic(err)
		}
		report = append(report, val)
	}
	return
}
