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

func Day12() {
	fmt.Println("Day 12")

	file := "./Day12/day12_input.txt"
	fmt.Println("Part 1: ", calcFencePrices(file))
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

	return 0
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
