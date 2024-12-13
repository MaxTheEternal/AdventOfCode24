package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type button struct {
	cost int
	xDir int
	yDir int
}

type mashine struct {
	a       button
	b       button
	xTarget int
	yTarget int
}

func Day13() {
	fmt.Println("Day 13")
	file := "./Day13/day13_input.txt"
	fmt.Println("Part 1: ", PartOne(file))
}

func PartOne(file string) int {
	mashines := readFile(file)
	sum := 0

	for _, m := range mashines {
		sum += bruteForcePrice(m)
	}
	return sum
}

func bruteForcePrice(m mashine) int {
	for a := 0; a <= 100; a++ {
		xa := m.xTarget - (a * m.a.xDir)
		ya := m.yTarget - (a * m.a.yDir)

		if xa == 0 && ya == 0 {
			return a * 3
		}
		if xa < 0 || ya < 0 {
			break
		}

		for b := 0; b <= 100; b++ {
			x := m.xTarget - (a * m.a.xDir) - (b * m.b.xDir)
			y := m.yTarget - (a * m.a.yDir) - (b * m.b.yDir)

			if x == 0 && y == 0 {
				return a*3 + b*1
			}
			if x < 0 || y < 0 {
				break
			}
		}
	}
	return 0
}

func readFile(file string) []mashine {

	openfile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	result := []mashine{}
	scanner := bufio.NewScanner(openfile)
	for scanner.Scan() {
		a := scanner.Text()
		scanner.Scan()
		b := scanner.Text()
		scanner.Scan()
		p := scanner.Text()
		scanner.Scan()

		aVals := strings.Split(strings.TrimPrefix(a, "Button A: X+"), ", Y+")
		bVals := strings.Split(strings.TrimPrefix(b, "Button B: X+"), ", Y+")
		prize := strings.Split(strings.TrimPrefix(p, "Prize: X="), ", Y=")

		ax, _ := strconv.Atoi(aVals[0])
		ay, _ := strconv.Atoi(aVals[1])
		bx, _ := strconv.Atoi(bVals[0])
		by, _ := strconv.Atoi(bVals[1])
		px, _ := strconv.Atoi(prize[0])
		py, _ := strconv.Atoi(prize[1])

		aButton := button{3, ax, ay}
		bButton := button{1, bx, by}

		automat := mashine{aButton, bButton, px, py}
		result = append(result, automat)

	}
	return result
}
