package main

import (
	"fmt"

	"github.com/scottyw/adventofcode2021/pkg/aoc"
)

func main() {
	score := score(aoc.FileLinesToStringSlice("input/04.txt"))
	fmt.Println(score)
}

func score(lines []string) int {

	numbers := aoc.ToIntSlice(lines[0], ",")
	boards := aoc.To3DIntSlice(lines[1:], " ")

	for _, number := range numbers {

		// Update boards
		for _, board := range boards {
			update(board, number)
		}

		// Check for a winner
		var nonwinners [][][]int
		for _, board := range boards {
			if winner(board) {
				// Is this the last remaining board?
				if len(boards) == 1 {
					return unmarked(board) * number
				}
			} else {
				nonwinners = append(nonwinners, board)
			}
		}
		boards = nonwinners

	}

	return 0
}

func update(board [][]int, number int) {
	for x, row := range board {
		for y, i := range row {
			if i == number {
				board[x][y] = -1
			}
		}
	}
}

func winner(board [][]int) bool {

	for y := 0; y < 5; y++ {
		var sum int
		for x := 0; x < 5; x++ {
			sum += board[x][y]
		}
		if sum == -5 {
			return true
		}
	}

	for x := 0; x < 5; x++ {
		var sum int
		for y := 0; y < 5; y++ {
			sum += board[x][y]
		}
		if sum == -5 {
			return true
		}
	}

	return false
}

func unmarked(board [][]int) int {
	var sum int
	for _, row := range board {
		for _, i := range row {
			if i != -1 {
				sum += i
			}
		}
	}
	return sum
}

func parse(lines []string) ([]int, [][][]int) {
	numbers := aoc.ToIntSlice(lines[0], ",")
	boards := aoc.To3DIntSlice(lines[1:], " ")
	return numbers, boards
}
