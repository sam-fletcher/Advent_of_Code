package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	rows, err := readLines("day11.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of elements in input: %d\n", len(rows))
	fmt.Println(rows)

	fmt.Println("Day 11:")
	if err = partBoth(rows); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start)) // 0.7 ms
}

type octopi struct {
	flashes   int
	grid      [][]int
	flashGrid [][]bool
}

func partBoth(grid [][]int) error {
	steps := 1000
	flashGrid := make([][]bool, len(grid))
	for i := range flashGrid {
		flashGrid[i] = make([]bool, len(grid[i]))
	}
	octo := octopi{grid: grid, flashGrid: flashGrid}
	for step := 1; step <= steps; step++ {
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				octo.update(i, j)
			}
		}
		numFlashes := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if octo.flashGrid[i][j] {
					octo.flashes++
					octo.flashGrid[i][j] = false
					octo.grid[i][j] = 0
					numFlashes++
				}
			}
		}
		if step == 100 {
			fmt.Printf("Flashes after %d steps: %d\n", steps, octo.flashes)
		}
		if numFlashes == len(grid)*len(grid[0]) {
			fmt.Printf("Step when all flash at once: %d\n", step)
			return nil
		}
	}
	return nil
}

func (o *octopi) update(i, j int) {
	o.grid[i][j]++
	if o.grid[i][j] <= 9 || o.flashGrid[i][j] {
		return
	}
	o.flashGrid[i][j] = true
	if i > 0 {
		o.update(i-1, j)
		// diagonals:
		if j > 0 {
			o.update(i-1, j-1)
		}
		if j < len(o.grid[i])-1 {
			o.update(i-1, j+1)
		}
	}
	if i < len(o.grid)-1 {
		o.update(i+1, j)
		// diagonals:
		if j < len(o.grid[i])-1 {
			o.update(i+1, j+1)
		}
		if j > 0 {
			o.update(i+1, j-1)
		}
	}
	if j > 0 {
		o.update(i, j-1)
	}
	if j < len(o.grid[i])-1 {
		o.update(i, j+1)
	}
	return
}

func readLines(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		nums := []int{}
		for _, x := range line {
			num, err := strconv.Atoi(x)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		lines = append(lines, nums)
	}
	return lines, scanner.Err()
}
