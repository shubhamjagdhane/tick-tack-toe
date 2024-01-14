package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	game := NewGame()

	player1 := GetNewPlayer()
	player2 := GetNewPlayer()

	reader := bufio.NewReader(os.Stdin)
	defer reader.Reset(os.Stdin)

	moveCounter := 0

	for !game.IsDraw() {
		if player1.IsWon(game) {
			fmt.Printf("****** %s has won the game ******", player1.name)
			break
		}

		if player2.IsWon(game) {
			fmt.Printf("****** %s has won the game ******", player2.name)
			break
		}

		var player Player
		fmt.Println(game)

		if moveCounter%2 == 0 {
			player = player1
		} else {
			player = player2
		}

		moveCounter++

		fmt.Printf("Please enter your move %s: ", player.name)
		move, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading move:", err)
			continue
		}

		move = strings.TrimRight(move, "\n")

		m, err := strconv.ParseInt(move, 10, 64)
		if err != nil {
			fmt.Println("Error while reading move:", err)
			continue
		}

		if err = player.Move(game, validMoves(m)); err != nil {
			fmt.Println(err)
		}

	}

	fmt.Println()

}
