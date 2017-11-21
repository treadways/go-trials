package main

import "fmt"
import "math/rand"

func main() {

	var repeat string = "yes"
	var newMatch = "yes"
	var playerMove, comMove string
	var outcome, playerPoints, comPoints float64
	var match int
	playerPoints = 0
	comPoints = 0
	match = 0


	/*
	var output float64
	output = checkWin("scissors", "paper")
	fmt.Println(output)
	*/


	for repeat == "yes" {
		fmt.Println(" ")
		fmt.Println("==============================================================================")
		fmt.Println("Welcome to Rock, Paper, Scissors by Shane Treadway.")
		fmt.Println("--")

		fmt.Println(" ")
		fmt.Println("You will play the computer for the best out of three games.")
		fmt.Println(" ")

		for newMatch == "yes" {
			match = match + 1
			fmt.Println("Match", match)
			fmt.Println("Type your choosen move (rock, paper, or scissors): ")
			fmt.Scanf("%s", &playerMove)

			comMove = RockPaperScissorsGenerator()
			fmt.Print("The computer played ", comMove, ".")
			fmt.Println(" ")

			outcome = checkWin(playerMove, comMove)

			if outcome == 1 {
				fmt.Println("Win")
			}
			if outcome == 0.5 {
				outcome = 0
				fmt.Println("Tie")
			}
			if outcome == 0 {
				fmt.Println("Loss")
			}

			fmt.Println(" ")
			playerPoints = playerPoints + outcome
			comPoints = comPoints + 1 - outcome
			fmt.Println("Your Wins: ", playerPoints)
			fmt.Println("Computer Wins: ", comPoints)
			fmt.Println(" ")

			if playerPoints == 2 {
				fmt.Println("Congradulations! You won in ", match, " matches.")
				newMatch = "no"
			}

			if comPoints == 2 {
				fmt.Println("Sorry, you lost in ", match, " matches.")
				newMatch = "no"
			}
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
		move = "rock"
	}
	if randomNumber == 1 {
		move = "paper"
	}
	if randomNumber == 2 {
		move = "scissors"
	}

	return move
}

func checkWin(move1 string, move2 string) float64 {
	var result float64
	if move1 == "rock" {
		if move2 == "rock" {
			result = 0.5
		}
		if move2 == "paper" {
			result = 0
		}
		if move2 == "scissors" {
			result = 1
		}
	}

	if move1 == "paper" {
		if move2 == "rock" {
			result = 1
		}
		if move2 == "paper" {
			result = 0.5
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
			result = 0.5
		}
	}

	return result
}