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
	rows, err := readLines("day8.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of elements in input: %d\n", len(rows))

	fmt.Println("Day 8 Part 1:")
	if err = part1(rows); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Day 8 Part 2:")
	if err = part2(rows); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start)) // 6 ms
}

func part1(rows [][]string) error {
	count := 0
	for _, row := range rows {
		afterPipe := false
		for _, code := range row {
			if code == "|" {
				afterPipe = true
			} else if afterPipe {
				if len(code) == 2 || len(code) == 3 || len(code) == 4 || len(code) == 7 {
					count++
				}
			}
		}
	}
	fmt.Printf("Number of unique numbers: %d\n", count)
	return nil
}

func part2(rows [][]string) error {
	total := 0
	for _, row := range rows {
		codes := row[:10]
		output := row[11:]
		digits := make(map[int]string)
		for _, code := range codes {
			if len(code) == 2 {
				digits[1] = code
			} else if len(code) == 4 {
				digits[4] = code
			}
		}
		for _, code := range codes {
			digit := determineNumber(code, digits[1], digits[4])
			digits[digit] = code
		}
		number := ""
		for _, code := range output {
			codeSet := makeSet(code)
			for k, v := range digits {
				if len(code) != len(v) {
					continue
				}
				digitSet := makeSet(v)
				if len(intersection(codeSet, digitSet)) == len(codeSet) {
					number += strconv.Itoa(k)
					break
				}
			}
		}
		num, err := strconv.Atoi(number)
		if err != nil {
			return err
		}
		total += int(num)
	}
	fmt.Printf("Total: %d\n", total)
	return nil
}

func determineNumber(code, one, four string) int {
	if len(code) == 2 {
		return 1
	}
	if len(code) == 4 {
		return 4
	}
	if len(code) == 3 {
		return 7
	}
	if len(code) == 7 {
		return 8
	}
	codeSet := makeSet(code)
	fourSet := makeSet(four)
	oneSet := makeSet(one)
	if len(code) == 5 { // 2, 3, 5
		if len(intersection(codeSet, fourSet)) == 3 {
			if len(intersection(codeSet, oneSet)) == 2 {
				return 3
			}
			return 5
		}
		return 2
	}
	if len(code) == 6 { // 0, 6, 9
		if len(intersection(codeSet, fourSet)) == 4 {
			return 9
		}
		if len(intersection(codeSet, oneSet)) == 2 {
			return 0
		}
		return 6
	}
	return -1
}

func makeSet(code string) map[string]struct{} {
	set := make(map[string]struct{})
	for _, c := range strings.Split(code, "") {
		set[c] = struct{}{}
	}
	return set
}

func intersection(a, b map[string]struct{}) map[string]struct{} {
	intersection := make(map[string]struct{})
	for k := range a {
		if _, exists := b[k]; exists {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func setDiff(a, b map[string]struct{}) map[string]struct{} {
	diff := make(map[string]struct{})
	for k := range a {
		if _, exists := b[k]; !exists {
			diff[k] = struct{}{}
		}
	}
	return diff
}

func readLines(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}
