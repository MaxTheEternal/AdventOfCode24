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

type config struct {
	v int
	n int
}

var cache = make(map[config]int)

func Day11() {
	fmt.Println("Day 11")
	file := "./Day11/day11_input.txt"
	// start := time.Now()
	// fmt.Println("Part 1: ", len(PartOne(file, 25)))
	// fmt.Println("Strings Took: ", time.Since(start))

	start := time.Now()
	fmt.Println("Part 1: ", PartTwo(file, 25))
	fmt.Println("ints cheated Took: ", time.Since(start))

	start = time.Now()
	fmt.Println("Part 2: ", PartTwo(file, 75))
	fmt.Println("ints cheated Took: ", time.Since(start))
}

func PartTwo(file string, amount int) int {
	start := readFile(file)
	sum := 0
	for _, s := range start {
		v, _ := strconv.Atoi(s)
		sum += calcTotalStones(v, amount)

	}
	return sum
}

func calcTotalStones(v, n int) int {
	if n == 0 {
		return 1
	}
	if r, ok := cache[config{v, n}]; ok {
		return r
	}
	if v == 0 {
		res := calcTotalStones(1, n-1)
		cache[config{v, n}] = res
		return res
	}
	s := strconv.Itoa(v)
	if len(s)%2 == 0 {
		a, _ := strconv.Atoi(s[:len(s)/2])
		b, _ := strconv.Atoi(s[len(s)/2:])
		res := calcTotalStones(a, n-1) + calcTotalStones(b, n-1)
		cache[config{v, n}] = res
		return res
	}

	res := calcTotalStones((v * 2024), n-1)
	cache[config{v, n}] = res
	return res
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
