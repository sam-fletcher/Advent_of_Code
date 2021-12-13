package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	rows, err := readLines("day12.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of elements in input: %d\n", len(rows))
	fmt.Println(rows)

	fmt.Println("Day 12 Part 1:")
	if err = part(rows, 1); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Day 12 Part 2:")
	if err = part(rows, 2); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time taken: %s\n", time.Since(start)) // 850 ms
}

func part(rows [][]string, maxVisits int) error {
	g, err := makeGraph(rows)
	if err != nil {
		return err
	}
	fmt.Printf("Number of possible paths: %d\n", g.numPaths(maxVisits))
	return nil
}

func makeGraph(rows [][]string) (graph, error) {
	g := graph{vertices: make(map[string]*vertex), visitOrders: make([][]string, 0)}
	for _, row := range rows {
		for _, ele := range row {
			g.addVertex(ele)
		}
	}
	for _, row := range rows {
		if err := g.addEdge(row[0], row[1]); err != nil {
			return graph{}, err
		}
	}
	return g, nil
}

type vertex struct {
	big      bool
	adjacent map[string]*vertex
}

type graph struct {
	vertices    map[string]*vertex
	visitOrders [][]string
	maxVisits   int
}

func (g *graph) numPaths(visits int) int {
	g.maxVisits = visits
	g.DFS("start", []string{})
	return len(g.visitOrders)
}

func (g *graph) DFS(startKey string, visitOrder []string) {
	visitOrder = append(visitOrder, startKey)
	if startKey == "end" {
		g.visitOrders = append(g.visitOrders, visitOrder)
		return
	}
	for key := range g.vertices[startKey].adjacent {
		if g.cannotRevist(key, visitOrder) {
			continue
		}
		g.DFS(key, visitOrder)
	}
}

func (g *graph) cannotRevist(key string, visitOrder []string) bool {
	if g.vertices[key].big {
		return false
	}
	keyCount := 0
	visitCounts := make(map[string]int)
	for _, k := range visitOrder {
		if !g.vertices[k].big {
			visitCounts[k]++
		}
		if k == key {
			keyCount++
		}
	}
	if keyCount < 1 {
		return false
	}
	if keyCount >= 1 && (key == "start" || key == "end") {
		return true
	}
	maxCount := max(visitCounts)
	return maxCount >= g.maxVisits
}

func (g *graph) addVertex(key string) {
	if _, exists := g.vertices[key]; exists {
		return
	}
	v := &vertex{}
	if isUpper(key) {
		v.big = true
	}
	g.vertices[key] = v
	g.vertices[key].adjacent = make(map[string]*vertex)
}

func (g *graph) addEdge(from, to string) error {
	if _, exists := g.vertices[from].adjacent[to]; exists {
		return errors.New(fmt.Sprintf("error, origin vertex %s doesn't exist", from))
	}
	if _, exists := g.vertices[to]; !exists {
		return errors.New(fmt.Sprintf("error, target vertex %s doesn't exist", to))
	}
	g.vertices[from].adjacent[to] = g.vertices[to]
	g.vertices[to].adjacent[from] = g.vertices[from]
	return nil
}

// UTILS:

func isUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func max(counts map[string]int) int {
	maxCount := 0
	for _, v := range counts {
		if v > maxCount {
			maxCount = v
		}
	}
	return maxCount
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
		line := strings.Split(scanner.Text(), "-")
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}
