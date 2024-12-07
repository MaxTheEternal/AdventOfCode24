package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

type recordedPosition struct {
	X    int
	Y    int
	XDir int
	YDir int
}
type safeCounter struct {
	m       sync.Mutex
	counter int
}

func (c *safeCounter) Inc() {
	c.m.Lock()
	c.counter++
	c.m.Unlock()
}

func Day6() {
	file := "./Day6/day6_input.txt"
	fmt.Println("Day 6")
	fmt.Printf("Part 1: %v\n", partOne(file))
	start := time.Now()
	fmt.Printf("Part 2: %v\n", PartTwo(file))
	elapsed := time.Since(start)
	fmt.Printf("It took: %v\n", elapsed)

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

func PartTwo(path string) int {
	matrix, guardX, guardY := readFile(path)
	var wg sync.WaitGroup
	coutner := safeCounter{sync.Mutex{}, 0}

	for yIndex, line := range matrix {
		for xIndex, char := range line {
			if char == "#" {
				continue
			}
			if xIndex == guardX && yIndex == guardY {
				continue
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				if isLooping(matrix, guardX, guardY, xIndex, yIndex) {
					coutner.Inc()
				}
			}()
		}
	}

	wg.Wait()

	return coutner.counter
}

func isLooping(matrix [][]string, x, y, xBlock, yBlock int) bool {
	positions := []recordedPosition{}
	running := true

	xDir := 0
	yDir := -1

	yLen := len(matrix)
	if yLen == 0 {
		log.Fatal("Matirx shouldnt be empty")
	}
	xlen := len(matrix[0])

	for running {
		nextX := x + xDir
		nextY := y + yDir

		if nextX == xlen || nextX < 0 {
			return false
		}
		if nextY == yLen || nextY < 0 {
			return false
		}

		// fmt.Printf("X: %v Y: %v\n", x, y)
		currentPosition := recordedPosition{x, y, xDir, yDir}
		if slices.Contains(positions, currentPosition) {
			return true
		}
		positions = append(positions, currentPosition)

		virtualBlocker := nextX == xBlock && nextY == yBlock
		blockerIsNext := matrix[nextY][nextX] == "#"
		if blockerIsNext || virtualBlocker {
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
		}
	}
	return true
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
