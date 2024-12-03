package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file := "./Day3/day3_input.txt"
	sum := correctMuls(file)
	fmt.Printf("\nSum is: %v\n", sum)
}

func correctMuls(file string) int {
	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	sum := 0

	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")

		matches := r.FindAllString(line, -1)

		for _, v := range matches {

			v = strings.TrimPrefix(v, "mul(")
			v = strings.TrimSuffix(v, ")")

			nums := strings.Split(v, ",")
			if len(nums) != 2 {
				log.Fatal("Mold didnt contain 2 values")
			}

			number1, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatal(err)

			}
			number2, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatal(err)
			}
			sum += number1 * number2
		}
	}
	return sum
}
