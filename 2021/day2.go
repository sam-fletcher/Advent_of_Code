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
	fmt.Println("Day 2 Part 1:")
	err := part1()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Day 2 Part 2:")
	err = part2()
	if err != nil {
		fmt.Println(err)
	}
}

func part1() error {
	lines, err := readLines("day2.txt", 2)
	if err != nil {
		return err
	}
	horiz, depth := int64(0), int64(0)
	for _, line := range lines {
		direction := line[0]
		distance, err := strconv.ParseInt(line[1], 10, 32)
		if err != nil {
			return err
		}
		switch direction {
		case "forward":
			horiz += distance
		case "down":
			depth += distance
		case "up":
			depth -= distance
		default:
			return errors.New(fmt.Sprintf("invalid direction: %s", direction))
		}
	}
	fmt.Printf("Horizontal: %d, Depth: %d, Multiple: %d\n", horiz, depth, horiz*depth)
	return nil
}

func part2() error {
	lines, err := readLines("day2.txt", 2)
	if err != nil {
		return err
	}
	horiz, depth, aim := int64(0), int64(0), int64(0)
	for _, line := range lines {
		direction := line[0]
		distance, err := strconv.ParseInt(line[1], 10, 32)
		if err != nil {
			return err
		}
		switch direction {
		case "forward":
			horiz += distance
			depth += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		default:
			return errors.New(fmt.Sprintf("invalid direction: %s", direction))
		}
	}
	fmt.Printf("Horizontal: %d, Depth: %d, Multiple: %d\n", horiz, depth, horiz*depth)
	return nil
}

func readLines(filename string, itemsPerLine int) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) != itemsPerLine {
			return nil, errors.New(fmt.Sprintf("invalid line length: %s", line))
		}
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}
