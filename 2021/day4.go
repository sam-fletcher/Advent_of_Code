package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	guesses, boards, err := readLines("day4.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Day 4 Part 1:")
	if err = part1(guesses, boards); err != nil {
		fmt.Println(err)
	}
	for _, b := range boards {
		b.resetFound()
	}
	fmt.Println("Day 4 Part 2:")
	if err = part2(guesses, boards); err != nil {
		fmt.Println(err)
	}
}

func part1(guesses []int, boards []*board) error {
	for _, g := range guesses {
		for k, b := range boards {
			for i := 0; i < boardSize; i++ {
				for j := 0; j < boardSize; j++ {
					if b.numbers[i][j] == g {
						b.found[i][j] = true
						if b.isWinner() {
							fmt.Printf("Winner found! Board %d:\n%v\n%v\n", k, b.numbers, b.found)
							fmt.Printf("Score = %d\n", b.calcScore()*g)
							return nil
						}
					}
				}
			}
		}
	}
	return errors.New(fmt.Sprintf("No winner found out of %d boards", len(boards)))
}

func part2(guesses []int, boards []*board) error {
	numWinners := 0
	for _, g := range guesses {
		for k, b := range boards {
			for i := 0; i < boardSize; i++ {
				for j := 0; j < boardSize; j++ {
					if b.numbers[i][j] == g {
						b.found[i][j] = true
						if b.isWinner() {
							numWinners++
							if numWinners == len(boards)-1 {
								fmt.Printf("Last winner found! Board %d:\n%v\n%v\n", k, b.numbers, b.found)
								fmt.Printf("Score = %d\n", b.calcScore()*g)
								return nil
							}
						}
					}
				}
			}
		}
	}
	return errors.New(fmt.Sprintf("only found %d winners out of %d boards", numWinners, len(boards)))
}

const boardSize = 5

type board struct {
	numbers    [][]int
	found      [][]bool
	alreadyWon bool
}

func (b *board) resetFound() {
	for i := range b.found {
		for j := range b.found[i] {
			b.found[i][j] = false
		}
	}
}

func (b *board) isWinner() bool {
	if b.alreadyWon {
		return false
	}
	// rows
	for i := 0; i < boardSize; i++ {
		hits := 0
		for j := 0; j < boardSize; j++ {
			if b.found[i][j] {
				hits++
				if hits == boardSize {
					b.alreadyWon = true
					return true
				}
			}
		}
	}
	// columns
	for j := 0; j < boardSize; j++ {
		hits := 0
		for i := 0; i < boardSize; i++ {
			if b.found[i][j] {
				hits++
				if hits == boardSize {
					b.alreadyWon = true
					return true
				}
			}
		}
	}
	return false
}

func (b *board) calcScore() int {
	score := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if !b.found[i][j] {
				score += b.numbers[i][j]
			}
		}
	}
	return score
}

func readLines(filename string) ([]int, []*board, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	splitLine := strings.Split(scanner.Text(), ",")
	guesses := make([]int, len(splitLine))
	for i, str := range splitLine {
		g, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			return nil, nil, err
		}
		guesses[i] = int(g)
	}

	boards := make([]*board, 0)
	grid := make([][]int, 0)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) == 0 {
			if len(grid) == 0 {
				continue
			}
			b := &board{numbers: grid, found: make([][]bool, boardSize)}
			for i := range b.found {
				b.found[i] = make([]bool, boardSize)
			}
			boards = append(boards, b)
			grid = make([][]int, 0)
		} else {
			nums := make([]int, boardSize)
			for i, xStr := range line {
				x, err := strconv.ParseInt(xStr, 10, 32)
				if err != nil {
					return nil, nil, err
				}
				nums[i] = int(x)
			}
			grid = append(grid, nums)
		}
	}
	return guesses, boards, scanner.Err()
}
