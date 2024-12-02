package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file := "./Day2/day2_input.txt"
	value, err := SafeReports(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Safe rows: %v\n", value)

}
func SafeReports(path string) (int, error) {

	data, err := getData(path)
	if err != nil {
		return 0, err
	}
	fmt.Print(data)
	fmt.Println()

	safeRows := 0

	for _, v := range data {
		if rowIsSafe(v) {
			safeRows += 1
			fmt.Printf("Row %v is safe\n", v)
		}
	}

	return safeRows, nil

}

func rowIsSafe(row []int) bool {
	if len(row) == 0 {
		return false
	}
	if len(row) == 1 {
		return true
	}

	return checkAscending(row) || checkDescending(row)
}
func checkAscending(row []int) bool {

	count := 0

	for i := 0; i < len(row)-1; i++ {
		if row[i] > row[i+1] {
			dist := absolut(row[i] - row[i+1])
			if dist > 3 || dist == 0 {
				count += 1
				continue
			}
			count += 1
			continue
		}
		dist := absolut(row[i] - row[i+1])
		if dist > 3 || dist == 0 {
			count += 1
			continue
		}
	}
	return count <= 1
}
func checkDescending(row []int) bool {

	count := 0

	for i := 0; i < len(row)-1; i++ {
		if row[i] < row[i+1] {
			count += 1
			continue
		}
		dist := absolut(row[i] - row[i+1])
		if dist > 3 || dist == 0 {
			count += 1
			continue
		}
	}
	return count <= 1
}

func absolut(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getData(path string) ([][]int, error) {
	openFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	values := [][]int{}

	fileScanner := bufio.NewScanner(openFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbers := strings.Split(line, " ")

		row := []int{}
		for i := range numbers {
			number, err := strconv.Atoi(numbers[i])
			if err != nil {
				return nil, err
			}
			row = append(row, number)
		}

		values = append(values, row)
	}
	openFile.Close()

	return values, nil
}
