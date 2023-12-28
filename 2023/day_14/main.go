package main

import (
	"fmt"
	"os"
	"strings"
)

func tilt(columns []string) []string {
	shift := func(column []rune) ([]rune, int) {
		shifts := 0
		for j := len(column) - 1; j > 0; j-- {
			if j > 0 {
				curr := column[j]
				prev := column[j-1]
				if string([]rune{prev, curr}) == "O." {
					column[j] = 'O'
					column[j-1] = '.'
					j++
					shifts++
				}
			}
		}
		return column, shifts
	}

	for i, _ := range columns {
		shifts := -1
		var column []rune
		for shifts != 0 {
			column, shifts = shift([]rune(columns[i]))
			columns[i] = string(column)
		}
	}
	return columns
}

func rotate(dish []string) []string {
	result := make([]string, len(dish[0]))
	for r := range dish {
		for c := range dish[r] {
			result[c] += string(dish[len(dish)-r-1][c])
		}
	}
	return result
}

func load(dish []string) int {
	var l int
	for i, s := range rotate(dish) {
		l += strings.Count(s, "O") * (i + 1)
	}
	return l
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	dish := rotate(strings.Fields(string(f)))

	dish = tilt(dish)
	fmt.Println(load(dish))

	cycles, cache := 1000000000, map[string]int{}
	for i := 0; i < cycles; i++ {
		if s, ok := cache[fmt.Sprint(dish)]; ok {
			i = cycles - (cycles-i)%(i-s)
		}
		cache[fmt.Sprint(dish)] = i
		for range []string{"north", "west", "south", "east"} {
			dish = rotate(tilt(dish))
		}
	}
	fmt.Println(load(dish))
}
