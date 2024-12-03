package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	file := "./Day3/day3_input_test.txt"
	readFile(file)
}

func readFile(file string) []string {
	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(openFile)
	for scanner.Scan() {
		line := scanner.Text()
		r, _ := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")

		matches := r.FindAllString(line, -1)
		fmt.Print(matches)
	}
	return nil

}
