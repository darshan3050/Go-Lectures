package game

import (
	"fmt"
	"module/deck"
	"module/player"
)

type Game struct {
	player1 player.Player
	player2 player.Player
	turn    int
	deck    deck.Deck
}

var newGame *Game

func CreateGame(Player1Name, Player2Name string) *Game {
	if newGame != nil {
		return newGame
	}
	newGame = CreateNewGame(Player1Name, Player2Name)
	return newGame

}
func CreateNewGame(Player1Name, Player2Name string) *Game {
	var tempGame Game
	tempGame.player1 = *player.CreateNewPlayer(Player1Name)
	tempGame.player2 = *player.CreateNewPlayer(Player2Name)
	tempGame.turn = 0
	tempGame.deck = *deck.CreateNewDeck()
	return &tempGame
}
func (g *Game) Play() bool {
	fmt.Println("\n\n\n\n\n\n\n\nNew Turn\n")
	var IsSkipByPlayer1 string
	var IsSkipByPlayer2 string
	var MoveByPlayer1 bool
	var MoveByPlayer2 bool
	//Player 1

	fmt.Println("\n\nPlayer 1")
	//player draw card
	fmt.Println("Press S to Skip  or Any other key to Pick Card")
	fmt.Scanln(&IsSkipByPlayer1)
	//fmt.Println("test line", IsSkipByPlayer1)
	if IsSkipByPlayer1 == "S" || IsSkipByPlayer1 == "s" {
		MoveByPlayer1 = true
	}
	if MoveByPlayer1 == false {
		//	fmt.Println(MoveByPlayer1)
		Score_Drawn := g.deck.DrawCard()
		fmt.Println("\n\n You drawn : ", Score_Drawn)
		g.player1.Score += Score_Drawn
		fmt.Println("\n Player 1 Total: ", g.player1.Score)
		if g.player1.Score == 21 {
			fmt.Println("\n\n\n\n\n\n\t\t\t", g.player1.Name, "   Wins   the game ")
			return true
		} else if g.player1.Score > 21 {
			fmt.Println("\n\n\n\n\n\n\t\t\t", g.player2.Name, "   Wins   the game ")
			return true
		}

	}

	//Player 2
	fmt.Println("\n\nPlayer 2")
	fmt.Println("Press S to Skip or Any other key to Pick Card")
	fmt.Scanln(&IsSkipByPlayer2)
	if IsSkipByPlayer2 == "S" || IsSkipByPlayer2 == "s" {
		MoveByPlayer2 = true
	}
	//player draw card
	if MoveByPlayer2 == false {
		Score_Drawn := g.deck.DrawCard()
		g.player2.Score += Score_Drawn
		fmt.Println("You drawn : ", Score_Drawn)
		fmt.Println("Player 2 Total: ", g.player2.Score)

		if g.player2.Score == 21 {
			fmt.Println("\n\n\n\n\n\n\t\t\t", g.player2.Name, "   Wins   the game ")
			return true
		} else if g.player2.Score > 21 {
			fmt.Println("\n\n\n\n\n\n\t\t\t", g.player1.Name, "   Wins   the game ")
			return true
		}

	}

	if MoveByPlayer1 == true && MoveByPlayer2 == true {
		Iswin := IsSkipWin(g.player1, g.player2)
		return Iswin
	} else {
		return false
	}

}

// func IsPlayerWinner(player1, player2 player.Player) bool {
// 	if player1.Score == player2.Score {
// 		fmt.Println(" Match is Draw", player1.Score, "       ", player2.Score)

// 		return true
// 	} else if (player1.Score > player2.Score && player1.Score < 22) || player1.Score == 21 {
// 		fmt.Println("Player 1 Wins ")
// 		return true
// 	} else if (player1.Score < player2.Score && player2.Score < 22) || player2.Score == 21 {
// 		fmt.Println("Player 2 Wins ")
// 		return true
// 	} else {
// 		return false
// 	}

// }
func IsSkipWin(player1, player2 player.Player) bool {
	if player1.Score == player2.Score {
		fmt.Println(" Match is Draw\n \t\t\t,", player1.Name, " Score :", player1.Score, "\n\t\t\t", player2.Name, "  Score:", player2.Score)
		return true
	} else if player1.Score > player2.Score && player1.Score < 22 {
		fmt.Println("\n\n\n\n\n\n\t\t\t", player1.Name, "   Wins   the game ")
		return true
	} else if player1.Score < player2.Score && player2.Score < 22 {
		fmt.Println("\n\n\n\n\n\n\t\t\t", player2.Name, "   Wins   the game ")
		return true
	} else {
		return false
	}
}
