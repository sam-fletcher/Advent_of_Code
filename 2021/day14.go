package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	list, rules, err := readLines("day14.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of elements in input: %d\n", len(rules))
	fmt.Println(list)

	fmt.Println("Day 13 Part 1:")
	if err = part1(list, rules); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Day 13 Part 2:")
	if err = part2(list, rules); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start)) // 14 ms
}

func part1(start []string, rules []*rule) error {
	steps := 10
	ll := linkedList{}
	for _, data := range start {
		ll.add(data)
	}
	for step := 1; step <= steps; step++ {
		currNode := ll.head
		for currNode.next != nil {
			for _, r := range rules {
				if currNode.data == r.prev && currNode.next.data == r.next {
					ll.insert(currNode, currNode.next, r.new)
					currNode = currNode.next
					break
				}
			}
			currNode = currNode.next
		}
	}
	ll.printCountDiff()
	return nil
}

// Part 2 uses a completely different approach (no linked list required)
func part2(start []string, rules []*rule) error {
	counts := make(map[string]int)
	for i := 0; i < len(start)-1; i++ {
		a := start[i]
		b := start[i+1]
		counts[a+b]++
	}
	steps := 40
	for step := 1; step <= steps; step++ {
		tmp := make(map[string]int)
		for k, v := range counts {
			tmp[k] = v
		}
		for pair, count := range tmp {
			for _, r := range rules {
				if pair == r.prev+r.next {
					new1 := r.prev + r.new
					new2 := r.new + r.next
					counts[new1] += count
					counts[new2] += count
					counts[pair] -= count
					if counts[pair] < 1 {
						delete(counts, pair)
					}
					break
				}
			}
		}
	}
	singleCounts := make(map[string]int)
	for k, v := range counts {
		chars := strings.Split(k, "")
		singleCounts[chars[0]] += v
	}
	printDiff(singleCounts)
	return nil
}

type node struct {
	data string
	next *node
}

type linkedList struct {
	head *node
	tail *node
}

func (l *linkedList) add(data string) {
	n := &node{data: data}
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	l.tail.next = n
	l.tail = n
}

func (l *linkedList) insert(prev, next *node, newData string) {
	n := &node{data: newData}
	n.next = next
	prev.next = n
}

func (l linkedList) printCountDiff() {
	counts := make(map[string]int)
	currNode := l.head
	for currNode != nil {
		counts[currNode.data]++
		currNode = currNode.next
	}
	printDiff(counts)
}

func printDiff(counts map[string]int) {
	minCount, maxCount := math.MaxInt, 0
	for _, count := range counts {
		if count < minCount {
			minCount = count
		}
		if count > maxCount {
			maxCount = count
		}
	}
	fmt.Printf("Count difference: %d-%d = %d\n", maxCount, minCount, maxCount-minCount)
}

type rule struct {
	prev, next, new string
}

func readLines(filename string) (init []string, rules []*rule, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	chars := strings.Split(scanner.Text(), "")
	for _, c := range chars {
		init = append(init, c)
	}

	lines := make([][]string, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " -> ")
		lines = append(lines, line)
	}
	for _, line := range lines {
		if len(line) < 2 {
			continue
		}
		r := rule{}
		prevNext := strings.Split(line[0], "")
		r.prev = prevNext[0]
		r.next = prevNext[1]
		r.new = line[1]
		rules = append(rules, &r)
	}
	return init, rules, scanner.Err()
}
