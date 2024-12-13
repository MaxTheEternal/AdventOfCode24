package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
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
	start := time.Now()
	fmt.Println("Part 1: ", PartOne(file))
	fmt.Println("Took: ", time.Since(start))
	star2 := time.Now()
	fmt.Println("Part 2: ", PartTwo(file))
	fmt.Println("Took: ", time.Since(star2))
}

func PartOne(file string) int {
	mashines := readFile(file, false)
	sum := 0

	for _, m := range mashines {
		sum += bruteForcePrice(m)
	}
	return sum
}
func PartTwo(file string) int {
	mashines := readFile(file, true)
	sum := 0

	for _, m := range mashines {
		sum += bruteForcePriceExtreme(m)
	}
	return sum
}
func bruteForcePriceExtreme(m mashine) int {
	prices := []int{}
	a := 0
	for true {
		xa := m.xTarget - (a * m.a.xDir)
		ya := m.yTarget - (a * m.a.yDir)

		if xa == 0 && ya == 0 {
			prices = append(prices, a*3)
		}
		if xa < 0 || ya < 0 {
			break
		}

		b := 0
		for true {
			x := m.xTarget - (a * m.a.xDir) - (b * m.b.xDir)
			y := m.yTarget - (a * m.a.yDir) - (b * m.b.yDir)

			if x == 0 && y == 0 {
				prices = append(prices, a*3+b*1)
			}
			if x < 0 || y < 0 {
				break
			}
			b++
		}
		a++
	}

	if len(prices) == 0 {
		return 0
	}
	smallest := prices[0]
	for _, v := range prices {
		if v < smallest {
			smallest = v
		}
	}
	return smallest
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

func readFile(file string, convert bool) []mashine {

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

		dif := 10000000000000
		if convert {
			px = px + dif
			py = py + dif
		}

		aButton := button{3, ax, ay}
		bButton := button{1, bx, by}

		automat := mashine{aButton, bButton, px, py}
		result = append(result, automat)

	}
	return result
}
