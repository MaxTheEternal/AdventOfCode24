package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func Day10() {
	file := "./Day10/day10_input.txt"
	fmt.Println("Day 10")
	part1, part2 := CalcPaths(file)
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}

func CalcPaths(file string) (int, int) {
	matrixMap, zeros := readFile(file)
	if len(zeros) == 0 {
		log.Fatal("WTF")
	}
	sumPeaks := 0
	sumPaths := 0
	for _, cord := range zeros {
		peaks := []coordinate{}
		sumPaths += calcRoutes(matrixMap, cord.x, cord.y, 0, &peaks)
		sumPeaks += len(peaks)
	}
	return sumPeaks, sumPaths

}

func calcRoutes(matrix [][]int, x, y, value int, peaks *[]coordinate) int {
	if x < 0 || x >= len(matrix[0]) || y < 0 || y >= len(matrix) {
		return 0
	}
	if matrix[y][x] != value {
		return 0
	}
	if value == 9 {
		peak := coordinate{x, y}
		if !slices.Contains(*peaks, peak) {
			*peaks = append(*peaks, peak)
		}
		return 1
	}
	sumPaths := 0
	sumPaths += calcRoutes(matrix, x-1, y, value+1, peaks)
	sumPaths += calcRoutes(matrix, x+1, y, value+1, peaks)
	sumPaths += calcRoutes(matrix, x, y-1, value+1, peaks)
	sumPaths += calcRoutes(matrix, x, y+1, value+1, peaks)
	return sumPaths

}

func readFile(file string) ([][]int, []coordinate) {

	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(openFile)
	yIndex := 0
	result := [][]int{}
	zeros := []coordinate{}

	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		row := make([]int, len(chars))
		for xIndex, char := range chars {
			num, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal(err)
			}
			if num == 0 {
				zeros = append(zeros, coordinate{xIndex, yIndex})
			}
			row[xIndex] = num

		}
		result = append(result, row)
		yIndex++
	}
	return result, zeros

}
