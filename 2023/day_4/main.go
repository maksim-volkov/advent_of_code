package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// calculatePower will check how many have numbers are present in winning map,
// will return integer representing amount of matches (where -1 indicates no matches).
func calculatePower(winning map[int]struct{}, have []int) int {
	power := -1
	for _, n := range have {
		if _, ok := winning[n]; ok {
			power++
		}
	}
	return power
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	re := regexp.MustCompile(`Card +(\d+): (.*) \| (.*)`)

	var (
		cost        float64
		haveNumbers []int
		allWinning  []map[int]struct{}
		allHave     [][]int
		allMatches  []int
	)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		winningNumbers := make(map[int]struct{})
		haveNumbers = nil

		// Parse input line.
		line := scanner.Text()
		data := re.FindStringSubmatch(line)
		if len(data) != 4 {
			panic(fmt.Sprintf("Wrong game input format: '%s'", line))
		}

		// Ignore card id.
		_, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}

		// Process winning numbers.
		for _, n := range strings.Split(data[2], " ") {
			if n == "" {
				continue
			}
			i, err := strconv.Atoi(strings.TrimSpace(n))
			if err != nil {
				panic(err)
			}
			winningNumbers[i] = struct{}{}
		}

		//Process have numbers.
		allWinning = append(allWinning, winningNumbers)
		for _, n := range strings.Split(data[3], " ") {
			if n == "" {
				continue
			}
			i, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			haveNumbers = append(haveNumbers, i)
		}
		allHave = append(allHave, haveNumbers)

		// Calculate cost of card and remember amount of matches.
		power := calculatePower(winningNumbers, haveNumbers)
		if power >= 0 {
			c := math.Pow(2.0, float64(power))
			cost += c
		}
		allMatches = append(allMatches, power+1)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(int(cost))

	// Calculate amount of cards. First initialize list with 1 (cause we have at least one original card).
	allCardCount := make([]int, len(allMatches))
	for i := 0; i < len(allCardCount); i++ {
		allCardCount[i] = 1
	}
	for i, n := range allMatches {
		for j := i + 1; j < i+1+n; j++ {
			allCardCount[j] += allCardCount[i]
		}
	}

	cost = 0
	for _, c := range allCardCount {
		cost += float64(c)
	}

	fmt.Println(int(cost))
}
