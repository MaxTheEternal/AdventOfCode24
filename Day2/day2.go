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
	return rowIsSafeInternal(row, 0)
}
func rowIsSafeInternal(row []int, count int) bool {

	if count >= 1 {
		return false
	}
	if len(row) == 0 {
		return false
	}
	if len(row) == 1 {
		return true
	}
	//	ascending order
	if row[0] < row[1] {
		for i := 0; i < len(row)-1; i++ {
			diff := row[i+1] - row[i]
			if diff < 1 || diff > 3 {
				return false
			}
		}
	} else {
		for i := 0; i < len(row)-1; i++ {
			diff := row[i] - row[i+1]
			if diff < 1 || diff > 3 {
				return false
			}
		}
	}
	return true
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
