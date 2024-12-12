package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func Day12() {
	fmt.Println("Day 12")

}

func calcFencePrices(file string) int {
	garden := readFile(file)

	res := findClusters(garden)
	fmt.Println(res)
	return len(res)
}

func findClusters(garden [][]string) [][]coordinate {
	xLen := len(garden[0])
	yLen := len(garden)
	result := [][]coordinate{}

	result = append(result, findCluster(garden[0][0], 0, 0, xLen, yLen, &garden))
	// for y, line := range garden {
	// 	for y, cell := range line {
	// 		if cell == "" {
	// 			continue
	// 		}
	//
	// 	}
	// }
	return result
}

func findCluster(cell string, x, y, xLen, yLen int, garden *[][]string) []coordinate {
	if x < 0 || x == xLen || y < 0 || y == yLen {
		return []coordinate{}
	}
	curCell := (*garden)[y][x]
	if curCell == "" {
		return []coordinate{}
	}
	if curCell != cell {
		return []coordinate{}
	}
	result := []coordinate{coordinate{x, y}}
	(*garden)[y][x] = ""
	result = append(result, findCluster(cell, x-1, y, xLen, yLen, garden)...)
	result = append(result, findCluster(cell, x+1, y, xLen, yLen, garden)...)
	result = append(result, findCluster(cell, x, y-1, xLen, yLen, garden)...)
	result = append(result, findCluster(cell, x, y-1, xLen, yLen, garden)...)

	return result
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
