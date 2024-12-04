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
		if isSafe(report) {
			safeReports = safeReports + 1
		}
	}
	return 0
}

func Part2(input []string) int {
	return 0
}

func isSafe(report []int) bool {
	delta := report[1] - report[0]
	last := report[1]
	for _, el := range report[2:] {
		newDelta := el - last
		if newDelta == 0 {
			return false
		}
		if newDelta > 0 && delta < 0 {
			return false
		}
		if newDelta < 0 && delta > 0 {
			return false
		}

		if el-last < 1 || el-last > 3 {
			return false
		}
		delta = newDelta
		last = el
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
