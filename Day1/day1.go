package day1

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1() {

	file := "./Day1/day1_input.txt"

	fmt.Println("Day 1")
	totalDistance := CalculateTotalDistance(file)
	fmt.Println("Part 1")
	fmt.Printf("Total Distance: %v\n", totalDistance)

	totalSimilarity := SimilarityScore(file)
	fmt.Println("Part 2")
	fmt.Printf("Total Similarity: %v\n", totalSimilarity)
}

func SimilarityScore(file string) int {
	list1, list2, err := readFileIntoLists(file)
	if err != nil {
		log.Fatal(err)
	}

	rightMap := make(map[int]int)
	for _, v := range list2 {
		if count, ok := rightMap[v]; ok == true {
			rightMap[v] = count + 1
		} else {
			rightMap[v] = 1
		}
	}

	totalSimilarity := 0
	for _, v := range list1 {
		totalSimilarity += (v * rightMap[v])
	}

	return totalSimilarity

}

func CalculateTotalDistance(file string) int {
	list1, list2, err := readFileIntoLists(file)
	if err != nil {
		log.Fatal(err)
	}
	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0

	for i, _ := range list1 {
		distance := absolut(list1[i] - list2[i])
		totalDistance += distance
	}
	return totalDistance
}

func absolut(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func readFileIntoLists(filePath string) ([]int, []int, error) {

	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := strings.Split(line, "   ")

		number1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatal(err)

		}
		number2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatal(err)

		}

		list1 = append(list1, number1)
		list2 = append(list2, number2)
	}
	readFile.Close()

	if len(list1) != len(list2) {
		return nil, nil, errors.New("Lists of uneven length")
	}
	return list1, list2, nil

}
