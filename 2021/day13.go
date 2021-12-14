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
	rows, folds, err := readLines("day13.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of elements in input: %d\n", len(rows))
	fmt.Println(folds)

	fmt.Println("Day 13 Part 1 & 2:")
	if err = part(rows, folds); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start)) // 9 ms
}

func part(rows [][]int, folds []int) error {
	x, y := max(rows)
	grid := make([][]int, y+1)
	for i := range grid {
		grid[i] = make([]int, x+1)
	}
	for _, r := range rows {
		grid[r[1]][r[0]] = 1
	}
	for _, fold := range folds {
		if fold > 0 {
			for i := fold + 1; i < len(grid); i++ {
				mirror := 2*fold - i
				for j := range grid[i] {
					if grid[i][j] == 1 {
						grid[mirror][j] = 1
					}
				}
			}
			grid = grid[:fold+1]
		} else {
			fold *= -1
			for j := fold + 1; j < len(grid[0]); j++ {
				mirror := 2*fold - j
				for i := range grid {
					if grid[i][j] == 1 {
						grid[i][mirror] = 1
					}
				}
			}
			for i := range grid {
				grid[i] = grid[i][:fold+1]
			}
		}
		fmt.Printf("Fold %d: Activated cells: %d\n", fold, countActivations(grid))
	}
	printGrid(grid)
	return nil
}

func printGrid(grid [][]int) {
	for _, line := range grid {
		for _, cell := range line {
			if cell == 1 {
				fmt.Printf("@")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}

func countActivations(grid [][]int) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				count++
			}
		}
	}
	return count
}

func max(counts [][]int) (int, int) {
	x, y := 0, 0
	for _, dims := range counts {
		if dims[0] > x {
			x = dims[0]
		}
		if dims[1] > y {
			y = dims[1]
		}
	}
	return x, y
}

func readLines(filename string) ([][]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	lines := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		lines = append(lines, line)
	}
	rows := make([][]int, 0)
	folds := make([]int, 0)
	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		if len(line) == 1 {
			if strings.Contains(line[0], "y") || strings.Contains(line[0], "x") {
				tmp := strings.Split(line[0], "=")
				num, err := strconv.Atoi(tmp[1])
				if err != nil {
					return nil, nil, err
				}
				if strings.Contains(line[0], "y") {
					folds = append(folds, num)
				}
				if strings.Contains(line[0], "x") {
					folds = append(folds, -num)
				}
			}
		} else {
			nums := make([]int, 0)
			for _, x := range line {
				num, err := strconv.Atoi(x)
				if err != nil {
					return nil, nil, err
				}
				nums = append(nums, num)
			}
			rows = append(rows, nums)
		}
	}
	return rows, folds, scanner.Err()
}
