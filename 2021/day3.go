package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	rows, err := readLines("day3.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of lines in input: %d\n", len(rows))
	fmt.Println("Day 3 Part 1:")
	if err = part1(rows); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Day 3 Part 2:")
	if err = part2(rows); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start)) // 1ms
}

func part1(rows []string) error {
	zeroBit := "0"[0]
	sums := make([]int, len(rows[0]))
	for i := range rows[0] {
		for _, row := range rows {
			if row[i] != zeroBit {
				sums[i]++
			} else {
				sums[i]--
			}
		}
	}
	fmt.Println(sums)
	mostCommon := make([]byte, len(rows[0]))
	leastCommon := make([]byte, len(rows[0]))
	for i, sum := range sums {
		if sum > 0 {
			mostCommon[i] = 1
		} else {
			leastCommon[i] = 1
		}
	}
	var gamma, epsilon string
	for i := range mostCommon {
		gamma += fmt.Sprintf("%01b", mostCommon[i])
		epsilon += fmt.Sprintf("%01b", leastCommon[i])
	}
	g, err := strconv.ParseInt(gamma, 2, 64)
	if err != nil {
		return err
	}
	e, err := strconv.ParseInt(epsilon, 2, 64)
	if err != nil {
		return err
	}
	fmt.Printf("Power consumption: %d\n", g*e)
	return nil
}

func part2(rows []string) error {
	tmp := make([]string, len(rows))
	copy(tmp, rows) // took me ages to remember to remember I had to do this annoying Go slice copying
	mostCommon := findNumber(tmp, true)
	fmt.Println(mostCommon)
	copy(tmp, rows)
	leastCommon := findNumber(tmp, false)
	fmt.Println(leastCommon)

	ox, err := strconv.ParseInt(mostCommon, 2, 64)
	if err != nil {
		return err
	}
	co2, err := strconv.ParseInt(leastCommon, 2, 64)
	if err != nil {
		return err
	}
	fmt.Printf("Life support: %d\n", ox*co2)
	return nil
}

func findNumber(rows []string, findMost bool) string {
	zeroBit := "0"[0] // get the byte equivalent of the string
	for col := range rows[0] {
		fmt.Printf(".")
		sum := 0
		for _, row := range rows {
			if row[col] != zeroBit {
				sum++
			} else {
				sum--
			}
		}
		keep := "0"[0]
		if (findMost && sum >= 0) ||
			(!findMost && sum < 0) { // keep the less common number
			keep = "1"[0]
		}
		// fmt.Printf("Col %d keeping %d out of: %v\n", col, keep, rows)
		for i := len(rows) - 1; i >= 0; i-- {
			if rows[i][col] != keep {
				rows = append(rows[:i], rows[i+1:]...) // delete
				if len(rows) == 1 {
					return rows[0]
				}
			}
		}
	}
	return "oops"
}

func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
