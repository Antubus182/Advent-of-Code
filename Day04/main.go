package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	var points int = 0
	fmt.Println("Day 4")
	input := getInput("input.txt")
	lines := strings.Split(input, "\n")
	checkPoints(lines[0])
	for _, l := range lines {
		points = points + checkPoints(l)
	}
	fmt.Printf("Pile is worth a total of %d points\n", points)
}

func checkPoints(line string) int {
	var matches int = 0
	var points float64 = 0
	f := strings.Split((strings.Split(line, ":")[1]), "|")
	f[0] = strings.TrimSpace(f[0])
	f[1] = strings.TrimSpace(f[1])
	inputs := strings.Split(f[1], " ")
	wins := strings.Split(f[0], " ")
	for _, s := range inputs {
		if slices.Contains(wins, s) && s != "" {
			fmt.Println("point on: ", s)
			matches++
		}
	}

	fmt.Printf("there are %d matches\n", matches)
	if matches > 0 {
		points = math.Pow(2, float64(matches-1))
		fmt.Println(points)
	}

	return int(points)
}

func getInput(file string) string {
	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(input)
}
