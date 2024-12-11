package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type safeCounter struct {
	m       sync.Mutex
	counter int
}

func (c *safeCounter) Inc() {
	c.m.Lock()
	c.counter++
	c.m.Unlock()
}
func (c *safeCounter) Add(n int) {
	c.m.Lock()
	c.counter += n
	c.m.Unlock()
}

func Day11() {
	fmt.Println("Day 11")
	file := "./Day11/day11_input.txt"
	// start := time.Now()
	// fmt.Println("Part 1: ", len(PartOne(file, 25)))
	// fmt.Println("Strings Took: ", time.Since(start))

	start := time.Now()
	fmt.Println("Part 1: ", PartTwo(file, 25))
	fmt.Println("ints Took: ", time.Since(start))

	start = time.Now()
	fmt.Println("Part 2: ", PartTwo(file, 75))
	fmt.Println("ints Took: ", time.Since(start))
}

func PartTwo(file string, amount int) int {
	start := readFile(file)
	nums := make([]int, len(start))
	for i, s := range start {
		nums[i], _ = strconv.Atoi(s)
	}
	var wg sync.WaitGroup
	coutner := safeCounter{sync.Mutex{}, 0}

	for _, num := range nums {
		wg.Add(1)
		go func(n, a int) {
			defer wg.Done()
			c := safeCounter{sync.Mutex{}, 0}
			calcTotalStones(n, a, &c)
			coutner.Add(c.counter)
		}(num, amount)
	}

	wg.Wait()

	return coutner.counter
}

func calcTotalStones(n, amount int, counter *safeCounter) {
	if amount == 0 {
		counter.Inc()
		return
	}
	if n == 0 {
		calcTotalStones(1, amount-1, counter)
		return
	}
	nLen := len(strconv.Itoa(n))
	if nLen%2 == 0 {
		num1 := n / (IntPow(10, (nLen / 2)))
		num2 := n % (IntPow(10, (nLen / 2)))
		calcTotalStones(num1, amount-1, counter)
		calcTotalStones(num2, amount-1, counter)
		return
	}
	calcTotalStones((n * 2024), amount-1, counter)
}

func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	if m == 1 {
		return n
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func PartOne(file string, amount int) []string {
	result := readFile(file)
	for i := 0; i < amount; i++ {
		newStones := []string{}
		for _, stone := range result {
			s1, s2 := transform(stone)
			newStones = append(newStones, s1)
			if s2 != "" {
				newStones = append(newStones, s2)
			}
		}
		result = newStones
	}
	return result
}

func transform(stone string) (string, string) {
	if stone == "0" {
		return "1", ""
	}
	if len(stone)%2 == 0 {
		middle := len(stone) / 2
		s1 := stone[:middle]
		s2 := stone[middle:]
		num, _ := strconv.Atoi(s2)
		return s1, strconv.Itoa(num)
	}
	num, err := strconv.Atoi(stone)
	if err != nil {
		log.Fatal(err)
	}
	num = num * 2024
	return strconv.Itoa(num), ""
}

func readFile(file string) []string {
	openFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(openFile)
	scanner.Scan()
	line := scanner.Text()

	nums := strings.Split(line, " ")
	return nums
}
