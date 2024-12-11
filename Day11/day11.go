package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day11() {
	fmt.Println("Day 11")
	file := "./Day11/day11_input.txt"
	fmt.Println("Part 1: ", len(PartOne(file, 25)))
	// start := time.Now()
	// fmt.Println("Part 2: ", len(PartOne(file, 75)))
	// fmt.Println("Took: ", time.Since(start))

}

func PartOne(file string, amount int) []string {
	result := readFile(file)
	for i := 0; i < amount; i++ {
		newStones := []string{}
		for _, stone := range result {
			s1, s2 := transform(stone)
			newStones = append(newStones, s1)
			if s2 != "" {
				newStones = append(newStones, s2)
			}
		}
		result = newStones
	}
	return result
}

func transform(stone string) (string, string) {
	if stone == "0" {
		return "1", ""
	}
	if len(stone)%2 == 0 {
		middle := len(stone) / 2
		s1 := stone[:middle]
		s2 := stone[middle:]
		num, _ := strconv.Atoi(s2)
		return s1, strconv.Itoa(num)
	}
	num, err := strconv.Atoi(stone)
	if err != nil {
		log.Fatal(err)
	}
	num = num * 2024
	return strconv.Itoa(num), ""
}

func readFile(file string) []string {
	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(openFile)
	scanner.Scan()
	line := scanner.Text()

	nums := strings.Split(line, " ")
	return nums
}
