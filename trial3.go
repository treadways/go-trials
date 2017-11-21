package main

import "fmt"
import "math/rand"

func main() {
	var repeat string = "yes"
	var playerMove, comMove string
	var outcome, playerPoints, comPoints int
	var match int
	playerPoints = 0
	comPoints = 0
	match = 0

	for repeat == "yes" {
		fmt.Println(" ")
		fmt.Println("==============================================================================")
		fmt.Println("Welcome to Rock, Paper, Scissors by Shane Treadway.")
		fmt.Println("--")

		fmt.Println(" ")
		fmt.Println("You will play the computer for the best out of three games.")

		for playerPoints < 2 || comPoints < 2 {
			match = match + 1
			fmt.Println("Match", match)
			fmt.Println("Type your choosen move (rock, paper, or scissors): ")
			fmt.Scanf("%s", &playerMove)

			comMove = RockPaperScissorsGenerator()
			fmt.Print("The computer played ", comMove, ".")
			fmt.Println(" ")

			outcome = checkWin(playerMove, comMove)
			playerPoints = playerPoints + outcome
			comPoints = comPoints + 1 - outcome
		}

		fmt.Println("Would you like to play again? (yes/no)")
		fmt.Scanf("%s", &repeat)
		fmt.Println("==============================================================================")
		fmt.Println(" ")
	}
}

func RockPaperScissorsGenerator() string {
	var randomNumber int = rand.Intn(3)
	var move string

	if randomNumber == 0 {
		move = "Rock"
	}
	if randomNumber == 1 {
		move = "Paper"
	}
	if randomNumber == 2 {
		move = "Scissors"
	}

	return move
}

func checkWin(move1 string, move2 string) int {
	var result int
	if move1 == "rock" {
		if move2 == "rock" {
			result = 0
		}
		if move2 == "paper" {
			result = 0
		}
		if move2 == "scissor" {
			result = 1
		}
	}

	if move1 == "paper" {
		if move2 == "rock" {
			result = 1
		}
		if move2 == "paper" {
			result = 0
		}
		if move2 == "scissors" {
			result = 0
		}
	}
	if move1 == "scissors" {
		if move2 == "rock" {
			result = 0
		}
		if move2 == "paper" {
			result = 1
		}
		if move2 == "scissors" {
			result = 0
		}
	}

	return result
}