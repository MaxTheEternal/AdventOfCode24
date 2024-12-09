package day9

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day9() {
	file := "./Day9/day9_input.txt"
	fmt.Println("Day 9")
	fmt.Printf("Part 1: %v\n", PartOne(file))
	fmt.Printf("Part 2: %v\n", PartTwo(file))

}

func PartOne(file string) int {
	files := compactFiles(readFile(file))
	return calcCheckSum(files)
}

func PartTwo(file string) int {
	files := compactWholeFiles(readFile(file))
	fmt.Println(files)
	return calcCheckSum(files)
}

func calcCheckSum(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == -1 {
			continue
		}
		sum += (i * nums[i])
	}
	return sum
}

func compactWholeFiles(filesystem []int) []int {

	endPointer := len(filesystem) - 1

	for endPointer > 0 {
		currentDigit := filesystem[endPointer]
		if currentDigit == -1 {
			endPointer--
			continue
		}
		count := 0
		lastOccurence := endPointer
		for filesystem[endPointer] == currentDigit && endPointer > 0 {
			count += 1
			lastOccurence = endPointer
			endPointer--
		}

		freeIndex, err := findFreeSpaces(filesystem, count, lastOccurence)
		if err != nil {
			continue
		}

		for c := range count {
			filesystem[freeIndex+c] = filesystem[lastOccurence+c]
			filesystem[lastOccurence+c] = -1
		}
	}
	return filesystem
}

func findFreeSpaces(filesystem []int, amount int, lastCheck int) (int, error) {
	for i := 0; i < lastCheck; i++ {
		if filesystem[i] == -1 {
			freeCount := 0

			if i+amount > lastCheck {
				return 0, errors.New("No sufficient space found")
			}

			for c := range amount {
				if filesystem[i+c] == -1 {
					freeCount++
				}
			}

			if freeCount == amount {
				return i, nil
			}
		}
	}
	return 0, errors.New("No suffiecient space found")
}

func compactFiles(fileSystem []int) []int {
	beginning := 0
	end := len(fileSystem) - 1

	for beginning != end {
		if fileSystem[beginning] != -1 {
			beginning++
			continue
		}
		if fileSystem[end] == -1 {
			end--
			continue
		}

		fileSystem[beginning] = fileSystem[end]
		beginning++
		end--

	}

	if fileSystem[len(fileSystem)-1] == -1 {
		return fileSystem[:end]
	}
	return fileSystem[:end+1]
}

func readFile(file string) []int {
	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(openFile)

	scanner.Scan()
	nums := strings.Split(scanner.Text(), "")

	fileSystem := make([]int, 0, len(nums))

	fileIndex := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {

			count, err := strconv.Atoi(nums[i])
			if err != nil {
				log.Fatal(err)
			}
			for range count {
				fileSystem = append(fileSystem, fileIndex)
			}
			fileIndex++
		} else {
			count, err := strconv.Atoi(nums[i])
			if err != nil {
				log.Fatal(err)
			}
			for range count {

				fileSystem = append(fileSystem, -1)
			}
		}
	}

	return fileSystem

}
