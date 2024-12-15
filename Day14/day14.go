package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type robot struct {
	xPos int
	yPos int
	xDir int
	yDir int
}

func (r *robot) Move(xLen, yLen int) {
	r.xPos = (r.xPos + r.xDir) % xLen
	if r.xPos < 0 {
		r.xPos = r.xPos + xLen
	}
	r.yPos = (r.yPos + r.yDir) % yLen
	if r.yPos < 0 {
		r.yPos = r.yPos + yLen
	}
}

func Day14() {
	file := "./Day14/day14_input.txt"
	fmt.Println("Day 14")
	fmt.Println("Part 1: ", PartOne(file, 100, 101, 103))
	// fmt.Println("Part 2: ")
	// Part2(file, 100, 101, 103)
}

func PartOne(file string, seconds, xLen, yLen int) int {
	robots := readFile(file)
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, bot := range robots {
		// fmt.Println(bot)
		newX := (((bot.xDir * seconds) % xLen) + bot.xPos) % xLen
		if newX < 0 {
			newX = newX + xLen
		}
		newY := (((bot.yDir * seconds) % yLen) + bot.yPos) % yLen
		if newY < 0 {
			newY = newY + yLen
		}
		// fmt.Println("X and Y ", newX, newY)
		if newX < xLen/2 {
			if newY < yLen/2 {
				q1++
				// fmt.Println("Q1")
			} else if newY > yLen/2 {
				q3++
				// fmt.Println("Q3")
			}
		} else if newX > xLen/2 {
			if newY < yLen/2 {
				q2++
				// fmt.Println("Q2")
			} else if newY > yLen/2 {
				q4++
				// fmt.Println("Q4")
			}
		}
	}

	return q1 * q2 * q3 * q4

}

// func Part2(file string, seconds, xLen, yLen int) {
// 	robots := readFile(file)
//
// 	for s := range seconds {
// 		fmt.Println("After Second:", s)
// 		grid := make([][]string, yLen)
// 		for i := range yLen {
// 			row := make([]string, xLen)
// 			for j := range xLen {
// 				row[j] = " "
// 			}
// 			grid[i] = row
// 		}
// 		for _, r := range robots {
// 			r.Move(xLen, yLen)
// 			grid[r.yPos][r.xPos] = "#"
// 		}
//
// 		for _, line := range grid {
// 			fmt.Println(line)
// 		}
//
// 	}
// }

func readFile(file string) []robot {

	openfile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	result := []robot{}
	scanner := bufio.NewScanner(openfile)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(strings.TrimPrefix(line, "p="), " ")
		pos := strings.Split(parts[0], ",")
		xPos, _ := strconv.Atoi(pos[0])
		yPos, _ := strconv.Atoi(pos[1])
		vel := strings.Split(strings.TrimPrefix(parts[1], "v="), ",")
		xDir, _ := strconv.Atoi(vel[0])
		yDir, _ := strconv.Atoi(vel[1])

		result = append(result, robot{xPos, yPos, xDir, yDir})
	}
	return result
}
