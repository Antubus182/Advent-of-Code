package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	GameId int
	//Redmin   int
	Redmax int
	//Greenmin int
	Greenmax int
	//Bluemin  int
	Bluemax int
}

func main() {
	fmt.Println("Advent of Code day 2")
	input := getInput("input.txt")
	lines := strings.Split(input, "\n")
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	var Gameset []Game
	for _, g := range lines {
		sg := getGame(g)
		Gameset = append(Gameset, sg)
	}
	//fmt.Println(Gameset)
	var PossibleGames []Game
	for _, game := range Gameset {
		if game.Redmax <= 12 && game.Greenmax <= 13 && game.Bluemax <= 14 {
			PossibleGames = append(PossibleGames, game)
		}
	}
	//fmt.Println(PossibleGames)
	fmt.Print("Sum of id's of possible games part 1: ")
	fmt.Println(GameSum(PossibleGames))
	fmt.Print("Minimum Power of all games: ")
	fmt.Println(GamePower(Gameset))
}

func getInput(file string) string {
	input, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	return string(input)
}

func getGame(line string) Game {
	var game Game

	Idstring := strings.Split(line, ":")[0]
	Gameid := strings.Split(Idstring, "Game ")[1]

	GameSets := strings.Split(strings.Split(line, ":")[1], ";")
	game.GameId, _ = strconv.Atoi(Gameid)

	for _, set := range GameSets {
		//fmt.Println(set)
		cubes := strings.Split(set, ",")
		for _, color := range cubes {
			color = strings.TrimSpace(color)
			numCol := strings.Split(color, " ")
			//fmt.Println(numCol)
			//fmt.Print("color: ")
			//fmt.Println(numCol[1])
			//fmt.Print("Amount: ")
			//fmt.Println(numCol[0])
			amount, _ := strconv.Atoi(numCol[0])

			switch numCol[1] {
			case "red":
				if amount > game.Redmax {
					game.Redmax = amount
				}
			case "green":
				if amount > game.Greenmax {
					game.Greenmax = amount
				}
			case "blue":
				if amount > game.Bluemax {
					game.Bluemax = amount
				}
			default:
				fmt.Println("het is stuk")
			}

		}
	}
	return game
}
func GameSum(GameSet []Game) int {
	sum := 0
	for _, p := range GameSet {
		sum = sum + p.GameId
	}
	return sum
}

func GamePower(GameSet []Game) int {
	Tpower := 0
	for _, game := range GameSet {
		power := game.Redmax * game.Greenmax * game.Bluemax
		Tpower += power
	}
	return Tpower
}
