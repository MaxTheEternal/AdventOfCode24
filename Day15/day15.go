package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const emptyspace byte = 0x00
const box byte = 0x05
const wall byte = 0x09
const robot byte = 0x01

const up byte = 0x01
const down byte = 0x02
const right byte = 0x03
const left byte = 0x04

func Day15() {
	file := "./Day15/day15_input.txt"
	fmt.Println("Part 1: ", PartOne(file))
}

func PartOne(file string) int {
	warehouse, moves, x, y := readFile(file)

	for _, m := range moves {
		x, y = moveRobot(m, &warehouse, x, y)
	}

	sum := 0
	// fmt.Println("Result")
	for i, row := range warehouse {
		for j, cell := range row {
			// fmt.Print(cell)
			if cell == box {
				sum += (i * 100) + j
			}
		}
		// fmt.Println()
		// fmt.Println(sum)
	}
	return sum
}

func moveRobot(m byte, warehouse *[][]byte, x, y int) (int, int) {
	var xDir int
	var yDir int

	switch m {
	case up:
		xDir = 0
		yDir = -1
	case down:
		xDir = 0
		yDir = 1
	case left:
		xDir = -1
		yDir = 0
	case right:
		xDir = 1
		yDir = 0
	}

	newX := x + xDir
	newY := y + yDir
	switch (*warehouse)[newY][newX] {
	case emptyspace:
		(*warehouse)[newY][newX] = robot
		(*warehouse)[y][x] = emptyspace
		return newX, newY
	case wall:
		return x, y
	case box:
		if pushBox(warehouse, newX, newY, xDir, yDir) {
			(*warehouse)[newY][newX] = robot
			(*warehouse)[y][x] = emptyspace
			return newX, newY
		} else {
			return x, y
		}
	}
	return x, y
}

func pushBox(warehouse *[][]byte, x, y, xDir, yDir int) bool {

	if (*warehouse)[y+yDir][x+xDir] == wall {
		return false
	}
	if (*warehouse)[y+yDir][x+xDir] == box {
		return pushBox(warehouse, x+xDir, y+yDir, xDir, yDir)
	}
	if (*warehouse)[y+yDir][x+xDir] == emptyspace {
		(*warehouse)[y+yDir][x+xDir] = box
		return true
	}
	return false
}

func readFile(file string) ([][]byte, []byte, int, int) {

	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	x := 0
	y := 0
	lineIndex := 0
	warehouse := [][]byte{}
	moves := []byte{}

	instructions := false
	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			instructions = true
			continue
		}
		if !instructions {
			chars := strings.Split(line, "")
			row := make([]byte, len(chars))
			for i, c := range chars {
				switch c {

				case "@":
					row[i] = robot
					x = i
					y = lineIndex
				case "#":
					row[i] = wall
				case "O":
					row[i] = box
				case ".":
					row[i] = emptyspace
				}
			}
			warehouse = append(warehouse, row)
			lineIndex++
		} else {
			chars := strings.Split(line, "")
			for _, c := range chars {
				switch c {
				case ">":
					moves = append(moves, right)
				case "<":
					moves = append(moves, left)
				case "^":
					moves = append(moves, up)
				case "v":
					moves = append(moves, down)
				}
			}
		}
	}
	return warehouse, moves, x, y
}
