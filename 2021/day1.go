package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Part 1:")
	err := day1(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Part 2:")
	err := day1(3)
	if err != nil {
		fmt.Println(err)
	}
}

func day1(windowSize int) error {
	file, err := os.Open("day1.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	nums := make([]int64, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			return err
		}
		nums = append(nums, x)
	}
	count := 0
	fmt.Println(len(nums))
	for i := windowSize; i < len(nums); i++ {
		prev := nums[i-windowSize : i]
		next := nums[i+1-windowSize : i+1]
		if sum(next) > sum(prev) {
			count++
		}
	}
	fmt.Println(count)
	return scanner.Err()
}

func sum(nums []int64) int64 {
	sum := int64(0)
	for _, x := range nums {
		sum += x
	}
	return sum
}
