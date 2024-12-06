package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func Day6() {
	file := "./Day6/day6_input.txt"
	fmt.Println("Day 6")
	fmt.Printf("Part 1: %v\n", partOne(file))

}

func partOne(path string) int {
	matrix, x, y := readFile(path)
	running := true

	xDir := 0
	yDir := -1

	yLen := len(matrix)
	if yLen == 0 {
		log.Fatal("Matirx shouldnt be empty")
	}
	xlen := len(matrix[0])

	matrix[y][x] = "X"

	fmt.Printf("xLen: %v yLen: %v\n", xlen, yLen)
	for running {
		nextX := x + xDir
		nextY := y + yDir

		if nextX == xlen || nextX < 0 {
			running = false
			continue
		}
		if nextY == yLen || nextY < 0 {
			running = false
			continue
		}

		// fmt.Printf("X: %v Y: %v\n", x, y)

		if matrix[nextY][nextX] == "#" {
			// Turning 90 Degrees
			switch {
			case xDir == 0 && yDir == -1:
				xDir = 1
				yDir = 0
			case xDir == 1 && yDir == 0:
				xDir = 0
				yDir = 1
			case xDir == 0 && yDir == 1:
				xDir = -1
				yDir = 0
			case xDir == -1 && yDir == 0:
				xDir = 0
				yDir = -1
			}
		} else {
			x = nextX
			y = nextY
			matrix[y][x] = "X"
		}
	}

	numberOfX := 0

	for _, row := range matrix {
		for _, char := range row {
			if char == "X" {
				numberOfX += 1
			}
		}
	}
	return numberOfX
}

func readFile(path string) ([][]string, int, int) {
	openFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer openFile.Close()

	result := [][]string{}
	lineIndex := 0
	x := 0
	y := 0

	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		if len(line) == 0 {
			continue
		}
		if slices.Contains(line, "^") {
			y = lineIndex
			x = slices.Index(line, "^")
		}
		lineIndex++
		result = append(result, line)
	}
	return result, x, y
}
