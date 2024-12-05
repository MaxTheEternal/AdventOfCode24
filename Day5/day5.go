package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file := "./Day5/day5_input.txt"
	fmt.Printf("Part 1: %v\n", PartOne(file))
}

func PartTwo(file string) int {
	a, b, c := readInput(file)
	return sumOfMiddlePages(listOfUpdates(a, b, c, false))
}
func PartOne(file string) int {
	a, b, c := readInput(file)
	return sumOfMiddlePages(listOfUpdates(a, b, c, true))
}

func sumOfMiddlePages(pagelist [][]int) int {
	sum := 0
	for _, pages := range pagelist {
		sum += pages[len(pages)/2]
	}
	return sum
}

func listOfUpdates(lefts []int, rights []int, pagelist [][]int, correctOnes bool) [][]int {
	result := [][]int{}

	for _, pages := range pagelist {
		rule := correctOnes
		for pageIndex, page := range pages {
			beforeNums := []int{}
			for rightIndex, rightValue := range rights {
				if rightValue == page {
					beforeNums = append(beforeNums, lefts[rightIndex])
				}
			}
			for _, v := range beforeNums {
				index := slices.Index(pages, v)
				if index > pageIndex {
					rule = !correctOnes
				}

			}

		}
		if rule {
			result = append(result, pages)
		}
	}
	return result
}

func readInput(file string) ([]int, []int, [][]int) {
	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	left := []int{}
	right := []int{}
	pages := [][]int{}

	scanner := bufio.NewScanner(openFile)
	rules := true
	pageIndex := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			rules = false

		} else if rules {
			nums := strings.Split(line, "|")
			if len(nums) != 2 {
				log.Fatal("Rules read incorrectly")
			}
			num1, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatal(err)
			}
			num2, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatal(err)
			}
			left = append(left, num1)
			right = append(right, num2)

		} else {
			nums := strings.Split(line, ",")
			page := []int{}
			for _, v := range nums {
				num, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				page = append(page, num)
			}
			pages = append(pages, page)
			pageIndex++

		}
	}
	return left, right, pages
}
