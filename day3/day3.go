package day3

import (
	"AdventOfCode2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const mulPattern = "mul+\\(\\d{1,3},\\d{1,3}\\)"
const mulPattern2 = "(mul+\\(\\d{1,3},\\d{1,3}\\))|do\\(\\)|don't\\(\\)"

func SolvePart1() {
	fmt.Println(Part1(utils.ReadFileString("day3/data.txt")))
}

func SolvePart2() {
	fmt.Println(Part2(utils.ReadFileString("day3/data.txt")))
}

func Part1(input string) int {
	r, _ := regexp.Compile(mulPattern)
	matches := r.FindAllString(input, -1)
	result := 0
	for _, match := range matches {
		result = result + mulString(match)
	}
	return result
}

func Part2(input string) int {
	r, _ := regexp.Compile(mulPattern2)
	matches := r.FindAllString(input, -1)
	result := 0
	enabled := true
	for _, match := range matches {
		if match == "do()" {
			enabled = true
			continue
		}
		if match == "don't()" {
			enabled = false
			continue
		}
		if enabled {
			result = result + mulString(match)
		}

	}
	return result
}

func mulString(input string) int {
	result := strings.Split(input, ",")
	arg1 := strings.Replace(result[0], "mul(", "", 1)
	arg2 := strings.Replace(result[1], ")", "", 1)
	arg1Int, _ := strconv.Atoi(arg1)
	arg2Int, _ := strconv.Atoi(arg2)
	return mul(arg1Int, arg2Int)
}

func mul(a int, b int) int {
	return a * b
}
