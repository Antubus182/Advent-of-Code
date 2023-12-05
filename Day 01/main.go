package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Advent of Code Day 1")
	Instructions := getInput("input.txt")
	lines := strings.Fields(Instructions)
	answer := 0
	for _, s := range lines {
		value := getNumbers(s)
		answer = answer + value
	}

	fmt.Printf("The answer is: %d", answer)
}

func getInput(file string) string {
	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(input)
}

func getNumbers(line string) int {
	value := 0
	regex := `\d`
	matched, _ := regexp.Compile(regex)
	numbers := matched.FindAllString(line, -1)
	tens, _ := strconv.Atoi(numbers[0])
	singles, _ := strconv.Atoi(numbers[len(numbers)-1])
	value = tens*10 + singles
	return value
}
