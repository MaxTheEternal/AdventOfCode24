package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type coordiante struct {
	x int
	y int
}

func Day8() {
	file := "./Day8/day8_input.txt"
	fmt.Println("Day 8")
	fmt.Printf("Part 1: %v\n", partOne(file))
	fmt.Printf("Part 2: %v\n", partTwo(file))

}

func partOne(file string) int {
	coordMap, xSize, ySize := readFile(file)

	antiNodes := []coordiante{}

	for _, value := range coordMap {
		if len(value) < 2 {
			continue
		}

		for i := 0; i < len(value); i++ {
			for j := i + 1; j < len(value); j++ {
				antenna1 := value[i]
				antenna2 := value[j]

				xDelta := antenna2.x - antenna1.x
				yDelta := antenna2.y - antenna1.y

				newX1 := antenna2.x + xDelta
				newY1 := antenna2.y + yDelta

				if newX1 > -1 && newX1 < xSize && newY1 > -1 && newY1 < ySize {
					antiNode1 := coordiante{newX1, newY1}
					if !slices.Contains(antiNodes, antiNode1) {
						antiNodes = append(antiNodes, antiNode1)
					}
				}

				newX2 := antenna1.x - xDelta
				newY2 := antenna1.y - yDelta

				if newX2 > -1 && newX2 < xSize && newY2 > -1 && newY2 < ySize {
					antiNode2 := coordiante{newX2, newY2}
					if !slices.Contains(antiNodes, antiNode2) {
						antiNodes = append(antiNodes, antiNode2)
					}
				}
			}
		}
	}
	return len(antiNodes)
}

func partTwo(file string) int {
	coordMap, xSize, ySize := readFile(file)

	antiNodes := []coordiante{}

	for _, value := range coordMap {
		if len(value) < 2 {
			continue
		}

		for i := 0; i < len(value); i++ {
			for j := i + 1; j < len(value); j++ {
				antenna1 := value[i]
				antenna2 := value[j]

				xDelta := antenna2.x - antenna1.x
				yDelta := antenna2.y - antenna1.y

				countSize := max((xSize/xDelta)+1, (ySize/yDelta)+1)

				for count := -countSize; count < countSize; count++ {
					newX := antenna1.x + count*xDelta
					newY := antenna1.y + count*yDelta
					if isInBounds(newX, newY, xSize, ySize) {
						antiNode1 := coordiante{newX, newY}
						if !slices.Contains(antiNodes, antiNode1) {
							antiNodes = append(antiNodes, antiNode1)
						}
					}
				}
			}
		}
	}
	return len(antiNodes)
}

func isInBounds(x, y, xSize, ySize int) bool {
	return x > -1 && x < xSize && y > -1 && y < ySize
}

func readFile(filepath string) (map[string][]coordiante, int, int) {

	openFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(openFile)
	coordMap := make(map[string][]coordiante)

	yIndex := 0
	xSize := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		chars := strings.Split(line, "")
		xSize = len(chars)

		for xIndex, char := range chars {
			if char != "." {
				coordMap[char] = append(coordMap[char], coordiante{xIndex, yIndex})
			}
		}
		yIndex++
	}

	return coordMap, xSize, yIndex
}
