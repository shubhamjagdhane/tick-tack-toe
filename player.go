package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	name           string
	selectedOption Options
}

func NewPlayer(name string, option Options) Player {
	return Player{
		name:           name,
		selectedOption: option,
	}
}

type validMoves uint

const (
	One validMoves = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

type Point struct {
	x, y uint
}

var matrixMovesByValidMoves = map[validMoves]Point{
	One:   {0, 0},
	Two:   {0, 1},
	Three: {0, 2},
	Four:  {1, 0},
	Five:  {1, 1},
	Six:   {1, 2},
	Seven: {2, 0},
	Eight: {2, 1},
	Nine:  {2, 2},
}

func (p Player) Move(game *TickTackToe, m validMoves) error {
	point, ok := matrixMovesByValidMoves[m]
	if !ok {
		return fmt.Errorf("Invalid move, please select the valid move")
	}

	if game[point.x][point.y] != "" {
		return fmt.Errorf("already inserted, please select the non-inserted one")
	}

	game[point.x][point.y] = p.selectedOption

	return nil
}

func (p Player) IsWon(game *TickTackToe) bool {
	return game.IsWinner(p.selectedOption)
}

func GetNewPlayer() Player {
	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)
	defer reader.Reset(os.Stdin)

	fmt.Printf("Please enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error while reading player name: ", err)
		os.Exit(1)
	}

	var opt Options

	for {
		fmt.Printf("Please select your option(O or X): ")
		option, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading player option: ", err)
			continue
		}

		option = strings.TrimRight(option, "\n")

		if option == string(Cross) {
			opt = Cross
			break
		} else if option == string(Circle) {
			opt = Circle
			break
		} else {
			fmt.Println("Selected invalid option, either X or O valid")
		}
	}

	return NewPlayer(strings.TrimRight(name, "\n"), opt)

}
