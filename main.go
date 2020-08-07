package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Chess board setup
// X = "f3"
// Y = "h6"
// These queens cannot attack
//	 	 a b c d e f g h
//	1	 _ _ _ 7 _ _ _ _
//	2	 _ _ _ _ _ 4 _ _
//	3	 _ _ _ _ _ _ _ 2
//	4	 _ _ 5 _ _ _ _ _
//	5	 1 _ _ _ _ _ _ _
//	6	 _ _ _ _ _ _ 0 _
//	7	 _ _ _ _ 3 _ _ _
//	8	 _ 6 _ _ _ _ _ _

func main() {
	fmt.Println(time.Now())
	// create a two queen scenario
	queens := MakeQueens()

	// check if they can attack eachother
	attackPoss := CanQueenAttack(queens[0], queens[1])

	// if an attack is possible, recreate the scenario
	for attackPoss {
		queens = MakeQueens()
		attackPoss = CanQueenAttack(queens[0], queens[1])
	}

	// keep count (for interest sake)
	count := 0

	// loop until there are x queens on the board
	x := 7

	for len(queens) < x {
		// add another queen
		queens = append(queens, MakeCoord())

		// get length of new slice
		last := len(queens)

		// test the new queen against the existing queens
		for i := range queens {
			// dont test the queen against itself
			if i == last-1 {
				continue
			} else {
				attackPoss = CanQueenAttack(queens[i], queens[last-1])
			}

			// if an attack is possible, remove the recently added queen and exit the loop
			if attackPoss && len(queens) > 2 {
				queens = queens[:last-1]
				break
			}
		}
		count++
		// print new count on each loop
		fmt.Println(count)
	}

	// print the solution slice
	fmt.Println(queens)
	fmt.Println(time.Now())
}

func CanQueenAttack(x string, z string) bool {

	// calculate relative position of each point for diagonal attack
	// convert each string from ASCII to its relevant row/column
	dColumn := math.Abs(float64((int(x[0]) - 96) - (int(z[0]) - 96)))
	dRow := math.Abs(float64((int(x[1]) - 48) - (int(z[1]) - 48)))

	// test whether the queens are in the same column, row or diagonal
	// returning true means the queens CAN attack
	if x[0] == z[0] {
		return true
	} else if x[1] == z[1] {
		return true
	} else if dColumn == dRow {
		return true
	} else {
		return false
	}
}

// create a new coord from a1 through to h8 inclusive
func MakeCoord() string {
	row := rand.Intn(8) + 1
	col := rand.Intn(8) + 1
	var coord string
	switch col {
	case 1:
		coord = "a"
	case 2:
		coord = "b"
	case 3:
		coord = "c"
	case 4:
		coord = "d"
	case 5:
		coord = "e"
	case 6:
		coord = "f"
	case 7:
		coord = "g"
	case 8:
		coord = "h"
	}

	coord = coord + strconv.Itoa(row)

	return coord
}

// create initial pair
func MakeQueens() []string {
	queens := []string{"", ""}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(queens); i++ {
		queens[i] = MakeCoord()
	}

	return queens
}
