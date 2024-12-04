package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file := "./Day4/day4_input.txt"
	lines := readFileIntoLines(file)
	count1 := countXMas(lines)
	count2 := countXMasPart2(lines)
	fmt.Printf("XMAS Counted: %v\n", count1)
	fmt.Printf("Crosses Counted: %v\n", count2)
}

func countXMasPart2(lines []string) int {
	matches := 0

	leny := len(lines)
	if leny == 0 {
		return 0
	}

	lenx := len(lines[0])
	for y := 0; y < leny; y++ {
		for x := 0; x < lenx; x++ {
			char := string(lines[y][x])
			if char == "A" {
				if checkCrossMass(lines, x, y, lenx, leny) {
					matches++
				}

			}
		}
	}
	return matches
}

func checkCrossMass(lines []string, x, y, lenx, leny int) bool {
	if x == 0 || y == 0 {
		return false
	}
	if x == lenx-1 || y == leny-1 {
		return false
	}
	cross := string(lines[y+1][x-1]) + string(lines[y+1][x+1]) + string(lines[y-1][x-1]) + string(lines[y-1][x+1])

	switch cross {
	case "MSMS":
		return true
	case "MMSS":
		return true
	case "SSMM":
		return true
	case "SMSM":
		return true
	default:
		return false
	}

}

func countXMas(lines []string) int {
	matches := 0

	leny := len(lines)
	if leny == 0 {
		return 0
	}
	lenx := len(lines[0])
	for y := 0; y < leny; y++ {
		for x := 0; x < lenx; x++ {
			char := string(lines[y][x])
			if char == "X" {
				matches += checkWord(lines, "XMAS", x, y, lenx, leny)
			} else if char == "S" {
				matches += checkWord(lines, "SAMX", x, y, lenx, leny)
			}
		}
	}
	return matches
}

func checkWord(lines []string, word string, x, y, lenx, leny int) int {
	matches := 0
	if x < lenx-3 {
		straight := lines[y][x : x+4]
		// fmt.Println("straight: " + straight)
		if straight == word {
			// fmt.Println(straight + " straight")
			matches += 1
		}
	}
	if y < leny-3 {
		down := string(lines[y][x]) + string(lines[y+1][x]) + string(lines[y+2][x]) + string(lines[y+3][x])
		// fmt.Println("Down: " + down)
		if down == word {
			// fmt.Println(down + " down")
			matches += 1
		}
	}
	if x < lenx-3 && y < leny-3 {
		diagonal := string(lines[y][x]) + string(lines[y+1][x+1]) + string(lines[y+2][x+2]) + string(lines[y+3][x+3])
		// fmt.Println("Diagonal R: " + diagonal)
		if diagonal == word {
			// fmt.Println(diagonal + " diag R")
			matches += 1
		}
	}
	if x > 2 && y < leny-3 {
		diagonal := string(lines[y][x]) + string(lines[y+1][x-1]) + string(lines[y+2][x-2]) + string(lines[y+3][x-3])
		// fmt.Println("Diagonal L: " + diagonal)
		if diagonal == word {
			// fmt.Println(diagonal + " diag L")
			matches += 1
		}
	}

	return matches
}

func readFileIntoLines(file string) []string {
	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]string, 0, 140)
	scanner := bufio.NewScanner(openFile)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}
