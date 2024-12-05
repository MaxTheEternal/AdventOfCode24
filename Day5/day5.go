package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5() {
	file := "./Day5/day5_input.txt"
	fmt.Println("Day 5")
	fmt.Printf("Part 1: %v\n", PartOne(file))
	fmt.Printf("Par2: %v\n", PartTwo(file))
}

func PartTwo(file string) int {
	beforeMap, b := readInput(file)
	return sumOfMiddlePages(sortWrongUpdates(beforeMap, listOfUpdates(beforeMap, b, false)))
}
func PartOne(file string) int {
	beforeMap, b := readInput(file)
	return sumOfMiddlePages(listOfUpdates(beforeMap, b, true))
}

func sumOfMiddlePages(pagelist [][]int) int {
	sum := 0
	for _, pages := range pagelist {
		sum += pages[len(pages)/2]
	}
	return sum
}

func sortWrongUpdates(beforeMap map[int][]int, pagelist [][]int) [][]int {
	for _, pages := range pagelist {
		slices.SortFunc(pages, func(a, b int) int {
			befores := beforeMap[a]
			if slices.Contains(befores, b) {
				return 1
			} else {
				return -1
			}
		})
	}
	return pagelist
}

func listOfUpdates(beforemap map[int][]int, pagelist [][]int, correctOnes bool) [][]int {
	result := [][]int{}

	for _, pages := range pagelist {
		rule := correctOnes
		for pageIndex, page := range pages {
			beforeNums := beforemap[page]
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

func readInput(file string) (map[int][]int, [][]int) {
	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	beforeMap := make(map[int][]int)
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
			leftNum, err := strconv.Atoi(nums[0])
			if err != nil {
				log.Fatal(err)
			}
			rightNum, err := strconv.Atoi(nums[1])
			if err != nil {
				log.Fatal(err)
			}
			beforeMap[rightNum] = append(beforeMap[rightNum], leftNum)

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
	return beforeMap, pages
}
