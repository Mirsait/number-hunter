package main

import (
	"fmt"
	"math/rand/v2"
)

type GameState uint8

const (
	Start GameState = iota
	Game
	Win
	Exit
)

var hunting bool
var state = Start
var guessing_count int
var target_number int

func main() {
	hunting = true

	for {
		if !hunting {
			return
		}
		Clear()
		switch state {
		case Start:
			guessing_count = 0
			target_number = getRandomInt()
			start()
		case Game:
			fmt.Println("The secret number is hidden in the shadows...")
			fmt.Printf("Sharpen your senses, Hunter, and track it down!\n\n")
			guessing_count = game(target_number, guessing_count)
		case Win:
			win(guessing_count)
		case Exit:
			exit()
			hunting = false
		}
	}
}

func start() {
	fmt.Println("========================================")
	fmt.Println("      Welcome to NumberHunter!")
	fmt.Println("========================================")
	fmt.Println("You’ve entered the hunt for the secret number.")
	fmt.Println("Test your intuition and logic —")
	fmt.Println("can you uncover the mystery faster than anyone else?")
	fmt.Printf("\n\n")
	fmt.Println("NumberHunter — the hunt begins, the number awaits!")
	fmt.Println()
	fmt.Printf("\nPress Enter to continue ...")
	fmt.Scanln()
	state = Game
}

func game(target_number, guessing_count int) int {
	var guess_number int
	fmt.Println("Hunter, the shadows whisper...\nWhat number will you strike at?")
	fmt.Print("> [0 to Exit] ")
	_, err := fmt.Scanf("%d\n", &guess_number)

	if err != nil {
		return guessing_count
	}
	if guess_number == 0 {
		state = Exit
		return 0
	}

	guessing_count++
	if comparison(target_number, guess_number) {
		state = Win
	} else {
		fmt.Scanln()
	}

	return guessing_count
}

func comparison(target, guess int) bool {
	if guess < target {
		fmt.Println("Your arrow fell short... Aim higher, Hunter!")
		return false
	}
	if guess > target {
		fmt.Println("Your strike soared too far... Lower your aim!")
		return false
	}
	return true
}

func win(guessing_count int) {
	fmt.Println("Bullseye! You have slain the secret number.")
	fmt.Println("The hunt is complete — glory is yours!")
	fmt.Printf("\n\n\nHunter’s log: You loosed your arrows [%d]", guessing_count)
	fmt.Println(" times in pursuit of the secret number.")
	fmt.Println()
	replay()
}

func exit() {
	fmt.Println("The Hunter lays down his bow...")
	fmt.Println("The hunt ends, and the shadows remain unchallenged.")

	fmt.Printf("\nPress Enter to exit ...")
	fmt.Scanln()
}

func replay() {
	fmt.Print("Do you dare to begin another hunt? The numbers await in the darkness...\n")
	var answer string
	fmt.Print("> [Y] ")
	fmt.Scanln(&answer)
	if answer == "y" || answer == "Y" {
		state = Start
	} else {
		state = Exit
	}
}

func Clear() {
	fmt.Print("\033[H\033[2J")
}

func getRandomInt() int {
	return rand.IntN(10) + 1
}
