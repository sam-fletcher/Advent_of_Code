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
	rows, err := readLines("day5.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of lines in input: %d\n", len(rows))

	fmt.Println("Day 5 Part 1:")
	if err = part(rows, false); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Day 5 Part 2:")
	if err = part(rows, true); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start))  // 16ms
}

func part(vents []line, part2 bool) error {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}
	for _, line := range vents {
		if line.x1 == line.x2 || line.y1 == line.y2 {
			for x := line.minX(); x <= line.maxX(); x++ {
				for y := line.minY(); y <= line.maxY(); y++ {
					grid[x][y]++
				}
			}
		} else if part2 && line.isDiagonal() {
			if line.negSlope() {
				for x, y := line.minX(), line.minY(); x <= line.maxX() && y <= line.maxY(); {
					grid[x][y]++
					x++
					y++
				}
			} else if line.posSlope() {
				for x, y := line.minX(), line.maxY(); x <= line.maxX() && y >= line.minY(); {
					grid[x][y]++
					x++
					y--
				}
			}
		}
	}
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 1 {
				count++
			}
		}
	}
	fmt.Printf("Number of dangerous points: %d\n", count)
	return nil
}

type line struct {
	x1, y1, x2, y2 int
}

func (l *line) isDiagonal() bool {
	// is the height is the same as the width, aka a square
	return l.maxX()-l.minX() == l.maxY()-l.minY()
}

func (l *line) posSlope() bool {
	return (l.x1 < l.x2 && l.y1 > l.y2) || (l.x2 < l.x1 && l.y2 > l.y1)
}

func (l *line) negSlope() bool {
	return (l.x1 < l.x2 && l.y1 < l.y2) || (l.x2 < l.x1 && l.y2 < l.y1)
}

func (l *line) minX() int {
	if l.x1 < l.x2 {
		return l.x1
	}
	return l.x2
}

func (l *line) minY() int {
	if l.y1 < l.y2 {
		return l.y1
	}
	return l.y2
}

func (l *line) maxX() int {
	if l.x1 > l.x2 {
		return l.x1
	}
	return l.x2
}

func (l *line) maxY() int {
	if l.y1 > l.y2 {
		return l.y1
	}
	return l.y2
}

func readLines(filename string) ([]line, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]line, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		a := strings.FieldsFunc(scanner.Text(), Split)
		x1, err := strconv.Atoi(a[0])
		if err != nil {
			return nil, err
		}
		y1, err := strconv.Atoi(a[1])
		if err != nil {
			return nil, err
		}
		x2, err := strconv.Atoi(a[3])
		if err != nil {
			return nil, err
		}
		y2, err := strconv.Atoi(a[4])
		if err != nil {
			return nil, err
		}
		li := line{x1: x1, y1: y1, x2: x2, y2: y2}
		lines = append(lines, li)
	}
	return lines, scanner.Err()
}

func Split(r rune) bool {
	return r == ',' || r == ' '
}
