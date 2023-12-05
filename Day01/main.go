package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var WrittenNumbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}
var LowCombo = map[string]int{
	"one":       1,
	"two":       2,
	"three":     3,
	"four":      4,
	"five":      5,
	"six":       6,
	"seven":     7,
	"eight":     8,
	"nine":      9,
	"nineight":  9,
	"fiveight":  5,
	"threeight": 3,
	"oneight":   1,
	"sevenine":  7,
	"twone":     2,
	"eightwo":   8,
}

var HighCombo = map[string]int{
	"one":       1,
	"two":       2,
	"three":     3,
	"four":      4,
	"five":      5,
	"six":       6,
	"seven":     7,
	"eight":     8,
	"nine":      9,
	"nineight":  8,
	"fiveight":  8,
	"threeight": 8,
	"oneight":   8,
	"sevenine":  9,
	"twone":     1,
	"eightwo":   2,
}

func main() {
	fmt.Println("Advent of Code Day 1 part 2")
	Instructions := getInput("input.txt")
	lines := strings.Fields(Instructions)
	answer := 0
	for _, s := range lines {
		//value := getNumbers(s)
		value := getWrittenNumbers(s)
		answer = answer + value
	}

	fmt.Printf("The answer is: %d", answer)
	time.Sleep(10 * time.Second)
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

func getWrittenNumbers(line string) int {
	regex := `(\d|eightwo|twone|nineight|fiveight|threeight|oneight|sevenine|nine|eight|seven|six|five|four|three|two|one)`
	matched, _ := regexp.Compile(regex)
	numbers := matched.FindAllString(line, -1)
	tens, terr := strconv.Atoi(numbers[0])
	if terr != nil {
		tens = LowCombo[numbers[0]] * 10
	} else {
		tens = tens * 10
	}
	single, serr := strconv.Atoi(numbers[len(numbers)-1])
	if serr != nil {
		single = HighCombo[numbers[len(numbers)-1]]
	} else {
		single = single * 1
	}
	value := tens + single
	return value
}
