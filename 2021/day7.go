package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	rows, err := readLines("day7.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of elements in input: %d\n", len(rows))
	fmt.Println("Day 7 Part 1:")
	if err = part1(rows); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Day 7 Part 2:")
	if err = part2(rows); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start)) // 0.6 ms
}

func part1(positions []float64) error {
	midpoint := median(positions)
	fuel := 0
	for _, p := range positions {
		fuel += int(math.Abs(float64(midpoint - p)))
	}
	fmt.Printf("Total fuel cost: %d\n", fuel)
	return nil
}

func part2(positions []float64) error {
	midpoint := mean(positions)
	roundUp, roundDown := math.Ceil(midpoint), math.Floor(midpoint)
	fuelUp, fuelDown := 0.0, 0.0
	for _, p := range positions {
		n := math.Abs(roundUp - p)
		fuelUp += (n * (n + 1)) / 2
		n = math.Abs(roundDown - p)
		fuelDown += (n * (n + 1)) / 2
	}
	fmt.Printf("Total fuel cost: %0.0f or %0.0f\n", fuelUp, fuelDown) // not sure how to tell which one is optimal ahead of time
	return nil
}

func mean(x []float64) float64 {
	sum := 0.0
	for _, xi := range x {
		sum += xi
	}
	mean := sum / float64(len(x))
	fmt.Printf("Mean: %0.2f\n", mean)
	return mean
}

func median(x []float64) float64 {
	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})
	return x[(len(x) / 2)]
}

func readLines(filename string) ([]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	nums := make([]float64, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		for _, t := range text {
			x, err := strconv.Atoi(t)
			if err != nil {
				return nil, err
			}
			nums = append(nums, float64(x))
		}
	}
	return nums, scanner.Err()
}
