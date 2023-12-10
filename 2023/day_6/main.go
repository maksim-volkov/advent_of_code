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

func winningCombos(time, distance int) int {
	var winningCount int
	for t := 0; t < time; t++ {
		tmpDistance := (time - t) * t
		if tmpDistance > distance {
			winningCount++
		}
	}
	return winningCount
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var (
		times          []int
		distances      []int
		oneTimeStr     string
		oneDistanceStr string
		oneTime        int
		oneDistance    int
	)
	for _, t := range strings.Split(lines[0], " ")[1:] {
		if t == "" {
			continue
		}
		times = append(times, toNum(t))
		oneTimeStr = oneTimeStr + t
	}

	for _, d := range strings.Split(lines[1], " ")[1:] {
		if d == "" {
			continue
		}
		distances = append(distances, toNum(d))
		oneDistanceStr = oneDistanceStr + d
	}

	oneTime = toNum(oneTimeStr)
	oneDistance = toNum(oneDistanceStr)

	res := 1
	for i, t := range times {
		res *= winningCombos(t, distances[i])
	}

	fmt.Println(res)
	fmt.Println(winningCombos(oneTime, oneDistance))
}
