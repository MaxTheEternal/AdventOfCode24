package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type equation struct {
	result int
	nums   []int
}

func Day7() {
	file := "./Day7/day7_input.txt"
	fmt.Println("Day 7")
	fmt.Printf("Part 1: %v\n", partOne(file))
	fmt.Printf("Part 2: %v\n", partTwo(file))

}

func partOne(file string) int {
	equations := readFile(file)
	counter := 0
	for _, v := range equations {
		if calcPossiblity(v) {
			counter += v.result
		}
	}
	return counter
}

func partTwo(file string) int {
	equations := readFile(file)
	counter := 0

	for _, v := range equations {
		if calcPossiblityWithConcat(v) {
			counter += v.result
		}
	}
	return counter
}

func calcPossiblityWithConcat(eq equation) bool {

	if len(eq.nums) == 0 {
		log.Fatal("How did I get here")
	}
	if len(eq.nums) == 1 {
		return eq.nums[0] == eq.result
	}

	multiplied := equation{eq.result, append([]int{eq.nums[0] * eq.nums[1]}, eq.nums[2:]...)}
	summed := equation{eq.result, append([]int{eq.nums[0] + eq.nums[1]}, eq.nums[2:]...)}

	concat, err := strconv.Atoi(strconv.Itoa(eq.nums[0]) + strconv.Itoa(eq.nums[1]))
	if err != nil {
		log.Fatal(err)
	}

	concated := equation{eq.result, append([]int{concat}, eq.nums[2:]...)}

	return calcPossiblityWithConcat(multiplied) || calcPossiblityWithConcat(summed) || calcPossiblityWithConcat(concated)
}

func calcPossiblity(eq equation) bool {

	if len(eq.nums) == 0 {
		log.Fatal("How did I get here")
	}
	if len(eq.nums) == 1 {
		return eq.nums[0] == eq.result
	}

	multiplied := equation{eq.result, append([]int{eq.nums[0] * eq.nums[1]}, eq.nums[2:]...)}
	summed := equation{eq.result, append([]int{eq.nums[0] + eq.nums[1]}, eq.nums[2:]...)}

	return calcPossiblity(multiplied) || calcPossiblity(summed)
}

func readFile(file string) []equation {
	openfile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	results := []equation{}
	scanner := bufio.NewScanner(openfile)

	for scanner.Scan() {

		line := scanner.Text()
		parts := strings.Split(line, " ")

		res, err := strconv.Atoi(strings.TrimSuffix(parts[0], ":"))
		if err != nil {
			log.Fatal(err)
		}

		nums := []int{}
		for i := 1; i < len(parts); i++ {
			num, err := strconv.Atoi(parts[i])
			if err != nil {
				log.Fatal(err)
			}
			nums = append(nums, num)
		}
		results = append(results, equation{res, nums})

	}

	return results
}
