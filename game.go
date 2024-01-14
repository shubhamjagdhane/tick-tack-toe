package main

import (
	"bytes"
	"fmt"
)

type Options string

const (
	Circle Options = "O"
	Cross          = "X"
)

type TickTackToe [3][3]Options

func NewGame() *TickTackToe {
	return &TickTackToe{}
}

func (ttt TickTackToe) String() string {
	buf := bytes.Buffer{}
	defer buf.Reset()

	counter := 1
	for _, row := range ttt {
		buf.WriteString("-------------\n")
		buf.WriteString("|")
		for _, col := range row {
			if col == "" {
				buf.WriteString(fmt.Sprintf(" %d ", counter))
			} else {
				buf.WriteString(fmt.Sprintf(" %s ", col))
			}
			buf.WriteString("|")
			counter++
		}
		buf.WriteString("\n")
	}
	buf.WriteString("-------------\n")
	return buf.String()
}

func (ttt TickTackToe) IsEmpty() bool {
	counter := 0
	for row := range ttt {
		for col := range ttt[row] {
			if ttt[row][col] == "" {
				counter += 1
			}
		}
	}

	return counter == 9
}

func (ttt TickTackToe) IsWinner(option Options) bool {
	if ttt.IsEmpty() {
		return false
	}
	return ttt.CheckByDiagonals(option) || ttt.CheckByRows(option) || ttt.CheckByColumns(option)
}

func (ttt TickTackToe) CheckByRows(option Options) bool {
	// first row
	if ttt[0][0] == option &&
		ttt[0][0] == ttt[0][1] &&
		ttt[0][1] == ttt[0][2] {
		return true
	}

	// second row
	if ttt[1][0] == option &&
		ttt[1][0] == ttt[1][1] &&
		ttt[1][1] == ttt[1][2] {
		return true
	}

	// third row
	if ttt[2][0] == option &&
		ttt[2][0] == ttt[2][1] &&
		ttt[2][1] == ttt[2][2] {
		return true
	}

	return false
}

func (ttt TickTackToe) CheckByColumns(option Options) bool {
	// first column
	if ttt[0][0] == option &&
		ttt[0][0] == ttt[1][0] &&
		ttt[1][0] == ttt[2][0] {
		return true
	}

	// second column
	if ttt[0][1] == option &&
		ttt[0][1] == ttt[1][1] &&
		ttt[1][1] == ttt[2][1] {
		return true
	}
	// third column
	if ttt[0][2] == option &&
		ttt[0][2] == ttt[1][2] &&
		ttt[1][2] == ttt[2][2] {
		return true
	}

	return false
}

func (ttt TickTackToe) CheckByDiagonals(option Options) bool {
	// diagonal from left to right
	if ttt[0][0] == option &&
		ttt[0][0] == ttt[1][1] &&
		ttt[1][1] == ttt[2][2] {
		return true
	}

	// diagonal from right to left
	if ttt[0][2] == option &&
		ttt[0][2] == ttt[1][1] &&
		ttt[1][1] == ttt[2][0] {
		return true
	}

	return false
}

func (ttt TickTackToe) IsDraw() bool {
	if ttt.IsEmpty() {
		return false
	}

	counter := 0
	for row := range ttt {
		for col := range ttt[row] {
			if ttt[row][col] != "" {
				counter++
			}
		}
	}

	return counter == 9
}
