package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
	"unicode"
)

func toNum(runes []rune) int {
	n, err := strconv.Atoi(string(runes))
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var (
		y, sum int
	)
	symbolCoordinates := make(map[image.Point]rune)
	boundariesCoordinates := make(map[image.Point]struct{})
	boundaryCorrections := []image.Point{{0, 1}, {1, 0}, {1, 1}, {-1, 0}, {0, -1}, {-1, -1}, {1, -1}, {-1, 1}}

	// Build table with symbol coordinates
	var input []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for x, r := range line {
			if r != '.' && !unicode.IsDigit(r) {
				symbolCoordinates[image.Point{x, y}] = r
			}
		}
		y++
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var num []rune
	parts := make(map[image.Point][]int)
	for y, line := range input {
		// There is a chanse number was right at the end of the line (so no symbol after it),
		// so we need to check it here.
		if num != nil {
			for b := range boundariesCoordinates {
				if _, ok := symbolCoordinates[b]; ok {
					v := toNum(num)
					parts[b] = append(parts[b], v)
					sum += v
					break
				}
			}
			boundariesCoordinates = make(map[image.Point]struct{})
			num = nil
		}
		for x, r := range line {
			if !unicode.IsDigit(r) {
				if num != nil {
					for b := range boundariesCoordinates {
						if _, ok := symbolCoordinates[b]; ok {
							v := toNum(num)
							parts[b] = append(parts[b], v)
							sum += v
							break
						}
					}
					boundariesCoordinates = make(map[image.Point]struct{})
					num = nil
				}
			} else {
				num = append(num, r)
				for _, correction := range boundaryCorrections {
					boundariesCoordinates[image.Point{x, y}.Add(correction)] = struct{}{}
				}
			}
		}
	}

	fmt.Println(sum)

	sum = 0
	for coordinate, values := range parts {
		if v, ok := symbolCoordinates[coordinate]; ok && v == '*' {
			if len(values) == 2 {
				sum += values[0] * values[1]
			}
		}
	}

	fmt.Println(sum)
}
