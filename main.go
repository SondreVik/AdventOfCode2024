package main

import (
	"AdventOfCode2024/day1"
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]
	fmt.Println(os.Args)
	if arg == "1" {
		day1.Solve()
	} else {
		fmt.Println("Not solved yet...")
	}
}
