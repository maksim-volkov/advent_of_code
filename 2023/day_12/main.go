package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toNum(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var cache = make(map[string]int)

func count(spring string, workingCount []int) (arrangements int) {
	defer func() {
		recover()
		cache[fmt.Sprint(spring, workingCount)] = arrangements
	}()
	if c, ok := cache[fmt.Sprint(spring, workingCount)]; ok {
		return c
	}

	if len(spring) == 0 {
		if len(workingCount) == 0 {
			return 1
		}
		return 0
	}

	if spring[0] == '?' || spring[0] == '.' {
		arrangements += count(spring[1:], workingCount)
	}

	if (spring[0] == '#' || spring[0] == '?') &&
		!strings.Contains(spring[:workingCount[0]], ".") &&
		spring[workingCount[0]] != '#' {
		arrangements += count(spring[workingCount[0]+1:], workingCount[1:])
	}

	return
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var sum1, sum2 int

	for _, l := range lines {
		groups := strings.Split(l, " ")
		if len(groups) != 2 {
			panic(fmt.Sprintf("Wrong input line format: '%s'", l))
		}

		var working []int
		for _, n := range strings.Split(groups[1], ",") {
			working = append(working, toNum(n))
		}
		cache = make(map[string]int)
		sum1 += count(groups[0]+"?", working)

		for _, n := range strings.Split(strings.Repeat(","+groups[1], 4)[1:], ",") {
			working = append(working, toNum(n))
		}
		cache = make(map[string]int)
		sum2 += count(strings.Repeat(groups[0]+"?", 5), working)
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
