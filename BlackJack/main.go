package main

import (
	"fmt"
	"module/game"
)

func main() {
	var Player1Name, Player2Name string
	isWin := false
	fmt.Println("Welcome to BlackJack Game")
	fmt.Println("Enter player 1 Name: ")
	fmt.Scanln(&Player1Name)
	fmt.Println("Enter player 2 Name: ")
	fmt.Scanln(&Player2Name)
	NewGame := game.CreateGame(Player1Name, Player2Name)
	for i := 0; true; i++ {
		isWin = NewGame.Play()
		if isWin == true {
			break
		}

	}

}
