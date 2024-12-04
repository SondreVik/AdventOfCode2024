package day1

import (
	"AdventOfCode2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolvePart1() {
	fmt.Println(Part1(utils.ReadFile("day1/data.txt")))
}

func SolvePart2() {
	fmt.Println(Part2(utils.ReadFile("day1/data.txt")))
}

func Part1(input []string) int {
	colA, colB := getColumns(input)
	total := 0
	for key, numA := range colA {
		numB := colB[key]
		if numA > numB {
			total = total + (numA - numB)
		} else {
			total = total + (numB - numA)
		}
	}
	return total
}

func Part2(input []string) int {
	colA, colB := getColumns(input)
	total := 0
	for _, numA := range colA {
		multiplier := 0
		for _, numB := range colB {
			if numB == numA {
				multiplier = multiplier + 1
			}
		}
		total = total + numA*multiplier
	}
	return total
}

func getColumns(input []string) (colA []int, colB []int) {
	for _, line := range input {
		splitIndex := strings.IndexByte(line, ' ')
		partA := line[:splitIndex]
		partB := strings.Trim(line[splitIndex:], " ")
		numA, errA := strconv.Atoi(partA)
		if errA != nil {
			panic(errA)
		}
		numB, errB := strconv.Atoi(partB)
		if errB != nil {
			panic(errB)
		}
		colA = append(colA, numA)
		colB = append(colB, numB)
	}
	slices.Sort(colA)
	slices.Sort(colB)
	return
}
