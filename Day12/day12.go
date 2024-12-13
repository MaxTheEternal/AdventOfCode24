package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type side struct {
	x    int
	y    int
	xDir int
	yDir int
}

func Day12() {
	fmt.Println("Day 12")

	file := "./Day12/day12_input.txt"
	fmt.Println("Part 1: ", calcFencePrices(file))
	fmt.Println("Part 2: ", calcWithBulkDiscount(file))
}

func calcWithBulkDiscount(file string) int {
	garden := readFile(file)
	res := findClusters(garden)
	sum := 0
	for _, region := range res {
		sum += len(region) * calcSides(region)
	}
	return sum
}
func calcSides(region []coordinate) int {
	edges := []side{}
	for _, p := range region {
		if !slices.Contains(region, coordinate{p.x, p.y + 1}) {
			edges = append(edges, side{p.x, p.y, 0, 1})
		}
		if !slices.Contains(region, coordinate{p.x, p.y - 1}) {
			edges = append(edges, side{p.x, p.y, 0, -1})
		}
		if !slices.Contains(region, coordinate{p.x + 1, p.y}) {
			edges = append(edges, side{p.x, p.y, 1, 0})
		}
		if !slices.Contains(region, coordinate{p.x - 1, p.y}) {
			edges = append(edges, side{p.x, p.y, -1, 0})
		}
	}
	visitedEdges := []side{}

	sum := 0
	for _, e := range edges {
		sum += calcSideLength(e, &edges, &visitedEdges)
	}
	return sum
}

func calcSideLength(s side, edges, visited *[]side) int {
	if !slices.Contains(*edges, s) {
		return 0
	}
	if slices.Contains(*visited, s) {
		return 0
	}
	*visited = append(*visited, s)
	calcSideLength(side{s.x + 1, s.y, s.xDir, s.yDir}, edges, visited)
	calcSideLength(side{s.x - 1, s.y, s.xDir, s.yDir}, edges, visited)
	calcSideLength(side{s.x, s.y - 1, s.xDir, s.yDir}, edges, visited)
	calcSideLength(side{s.x, s.y + 1, s.xDir, s.yDir}, edges, visited)
	return 1
}

func calcFencePrices(file string) int {
	garden := readFile(file)
	res := findClusters(garden)
	sum := 0
	for _, region := range res {
		sum += len(region) * calcPerimiter(region)
	}
	return sum
}

func calcPerimiter(region []coordinate) int {
	perimiter := 0
	for _, reg := range region {
		if !slices.Contains(region, coordinate{reg.x + 1, reg.y}) {
			perimiter++
		}
		if !slices.Contains(region, coordinate{reg.x - 1, reg.y}) {
			perimiter++
		}
		if !slices.Contains(region, coordinate{reg.x, reg.y + 1}) {
			perimiter++
		}
		if !slices.Contains(region, coordinate{reg.x, reg.y - 1}) {
			perimiter++
		}
	}
	return perimiter
}

func findClusters(garden [][]string) [][]coordinate {
	xLen := len(garden[0])
	yLen := len(garden)
	result := [][]coordinate{}

	for y, line := range garden {
		for x, cell := range line {
			if cell == "" {
				continue
			}
			res := []coordinate{}
			findCluster(garden[y][x], x, y, xLen, yLen, &garden, &res)
			result = append(result, res)

		}
	}
	return result
}

func findCluster(cell string, x, y, xLen, yLen int, garden *[][]string, result *[]coordinate) {
	if x < 0 || x == xLen || y < 0 || y == yLen {
		return
	}
	curCell := (*garden)[y][x]
	if curCell == "" {
		return
	}
	if curCell != cell {
		return
	}
	*result = append(*result, coordinate{x, y})
	(*garden)[y][x] = ""
	findCluster(cell, x-1, y, xLen, yLen, garden, result)
	findCluster(cell, x+1, y, xLen, yLen, garden, result)
	findCluster(cell, x, y-1, xLen, yLen, garden, result)
	findCluster(cell, x, y+1, xLen, yLen, garden, result)

	return
}

func readFile(file string) [][]string {
	result := [][]string{}
	openfile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(openfile)

	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		result = append(result, chars)
	}
	return result
}
