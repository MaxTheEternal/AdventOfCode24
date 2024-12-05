package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3() {
	file := "./Day3/day3_input.txt"
	fmt.Println("Day 3")
	fmt.Printf("Part 1: %v\n", correctMuls(file))
	fmt.Printf("Part 2: %v\n", correctMulsWithInstructions(file))
}

func correctMulsWithInstructions(file string) int {
	lines := readFile(file)
	sum := 0
	muls := getMulsWithInstructions(lines)
	for _, v := range muls {
		sum += calcProductFromMul(v)
	}
	return sum
}

func correctMuls(file string) int {
	lines := readFile(file)
	sum := 0
	muls := getAllMuls(lines)
	for _, v := range muls {
		sum += calcProductFromMul(v)
	}
	return sum
}

func readFile(file string) []string {
	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := []string{}
	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return lines

}
func getMulsWithInstructions(lines []string) []string {
	muls := []string{}
	const doString string = "do()"
	const dontString string = "don't()"

	enabled := true

	// /(mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\))/gm
	r, err := regexp.Compile(`(mul\([0-9]{1,3},[0-9]{1,3}\)|do\(\)|don't\(\))`)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		for _, v := range matches {
			if v == doString {
				enabled = true
			} else if v == dontString {
				enabled = false
			} else if enabled {
				muls = append(muls, v)
			}

		}
	}
	return muls
}
func getAllMuls(lines []string) []string {
	muls := []string{}

	r, _ := regexp.Compile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)

	for _, line := range lines {
		matches := r.FindAllString(line, -1)
		muls = append(muls, matches...)
	}
	return muls
}

func calcProductFromMul(mul string) int {
	mul = strings.TrimPrefix(mul, "mul(")
	mul = strings.TrimSuffix(mul, ")")

	nums := strings.Split(mul, ",")
	if len(nums) != 2 {
		log.Fatal("Mol didnt contain 2 values")
	}

	number1, err := strconv.Atoi(nums[0])
	if err != nil {
		log.Fatal(err)

	}
	number2, err := strconv.Atoi(nums[1])
	if err != nil {
		log.Fatal(err)
	}
	return number1 * number2
}
