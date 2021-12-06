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
	rows, err := readLines("day6.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of element in input: %d\n", len(rows))
	fmt.Println("Day 6 Part 1:")
	if err = part(rows, 80); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Day 6 Part 2:")
	if err = part(rows, 256); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start)) // 0.15 ms
}

const (
	newFishDay = 8
	pregDays   = 6
)

func part(startingFish []int, numDays int) error {
	daysToBirth := make([]int, newFishDay+1)
	for _, fish := range startingFish {
		daysToBirth[fish]++
	}
	for i := 0; i < numDays; i++ {
		tmp := make([]int, len(daysToBirth))
		copy(tmp, daysToBirth)
		for d := len(daysToBirth) - 1; d >= 0; d-- {
			if d == 0 {
				tmp[pregDays] += daysToBirth[0]
				tmp[newFishDay] = daysToBirth[0]
			} else {
				tmp[d-1] = daysToBirth[d]
			}
		}
		daysToBirth = tmp
	}
	totalFish := 0
	for _, numFish := range daysToBirth {
		totalFish += numFish
	}
	fmt.Printf("Number of final fish: %d\n", totalFish)
	return nil
}

func readLines(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	nums := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		for _, t := range text {
			x, err := strconv.Atoi(t)
			if err != nil {
				return nil, err
			}
			nums = append(nums, x)
		}
	}
	return nums, scanner.Err()
}
