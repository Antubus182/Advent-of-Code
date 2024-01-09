package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Day 4")
	input := getInput("test.txt") //note that input is 140 lines of 140 characters
	lines := strings.Split(input, "\n")
	fmt.Println(lines[0])
	fmt.Println(lines[1])
	fmt.Println(lines[2])
	fmt.Println(lines[3])
}

func getInput(file string) string {
	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(input)
}
