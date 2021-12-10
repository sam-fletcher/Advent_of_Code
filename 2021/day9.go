package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	rows, err := readLines("day9.txt")
	if err != nil {
		fmt.Println(err)
	}
	boundaryRow := make([]int, len(rows[0]))
	for j := range boundaryRow {
		boundaryRow[j] = maxNum
	}
	rows = append([][]int{boundaryRow}, rows...)
	rows = append(rows, boundaryRow)
	fmt.Printf("Number of elements in input: %d\n", len(rows))
	// fmt.Println(rows)

	fmt.Println("Day 9 Part 1:")
	if err = part1(rows); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Day 9 Part 2:")
	if err = part2(rows); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start)) // 1 ms
}

func part1(grid [][]int) error {
	total := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			a := grid[i][j]
			if a < grid[i-1][j] && a < grid[i+1][j] &&
				a < grid[i][j-1] && a < grid[i][j+1] {
				total += a + 1
			}
		}
	}
	fmt.Printf("Sum of risk levels: %d\n", total)
	return nil
}

func part2(grid [][]int) error {
	basinGrid := make([][]int, len(grid))
	for i := range basinGrid {
		basinGrid[i] = make([]int, len(grid[i]))
	}
	basinID := 1
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if basinGrid[i][j] != 0 || grid[i][j] == 9 {
				continue
			}
			basinGrid = growBasin(i, j, basinID, grid, basinGrid)
			basinID++
		}
	}
	basinCounts := make([]int, basinID)
	for i := 1; i < len(basinGrid)-1; i++ {
		for j := 1; j < len(basinGrid[i])-1; j++ {
			idx := basinGrid[i][j]
			basinCounts[idx]++
		}
	}
	basinCounts = basinCounts[1:]
	sort.Slice(basinCounts, func(i, j int) bool {
		return basinCounts[i] > basinCounts[j]
	})
	fmt.Printf("Basin counts: %v\n", basinCounts)
	fmt.Printf("Multiple of three largest basins: %d\n",
		basinCounts[0]*basinCounts[1]*basinCounts[2])
	return nil
}

func growBasin(i, j, basinID int, grid, basinGrid [][]int) [][]int {
	if basinGrid[i][j] == 0 && grid[i][j] != 9 {
		basinGrid[i][j] = basinID
		basinGrid = growBasin(i+1, j, basinID, grid, basinGrid)
		basinGrid = growBasin(i-1, j, basinID, grid, basinGrid)
		basinGrid = growBasin(i, j+1, basinID, grid, basinGrid)
		basinGrid = growBasin(i, j-1, basinID, grid, basinGrid)
	}
	return basinGrid
}

const maxNum = 9

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
		nums := []int{maxNum}
		for _, x := range line {
			num, err := strconv.Atoi(x)
			if err != nil {
				return nil, err
			}
			nums = append(nums, num)
		}
		nums = append(nums, maxNum)
		lines = append(lines, nums)
	}
	return lines, scanner.Err()
}
