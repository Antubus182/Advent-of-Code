package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Advent of code Day 3")
	input := getInput("input.txt") //note that input is 140 lines of 140 characters
	lines := strings.Split(input, "\n")
	charlist := strings.Split(input, "")

	for _, c := range charlist {
		fmt.Println(c)
	}

	for _, l := range lines {
		chars := strings.Split(l, "")
		for _, ch := range chars {
			fmt.Print([]byte(ch))
		}
		fmt.Println(len(chars))
	}
}

func getInput(file string) string {
	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(input)
}
