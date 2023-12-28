package main

import (
	"fmt"
	"os"
	"strings"
)

type comparator func([][]rune, int, int, bool) (bool, bool)

func rowCmp(pattern [][]rune, i, j int, smudge bool) (bool, bool) {
	equal := true
	diffs := 0
	for k := 0; k < len(pattern[0]); k++ {
		if pattern[i][k] != pattern[j][k] {
			equal = false
			diffs++
			if diffs > 1 {
				break
			}
		}
	}
	if smudge {
		if diffs == 1 {
			return true, true
		}
		return equal, false
	}
	return equal, false
}

func columnCmp(pattern [][]rune, i, j int, smudge bool) (bool, bool) {
	equal := true
	diffs := 0
	for k := 0; k < len(pattern); k++ {
		if pattern[k][i] != pattern[k][j] {
			equal = false
			diffs++
			if diffs > 1 {
				break
			}
		}
	}
	if smudge {
		if diffs == 1 {
			return true, true
		}
		return equal, false
	}
	return equal, false
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(f)), "\n\n")

	var (
		pattern          [][]rune
		sum1, sum2       int
		matches          int
		smudgedAmount    int
		start, end, size int
		mirror           func(comparator, int, int, bool) []int
	)

	mirror = func(cmp comparator, start, end int, smudge bool) []int {
		if equal, smudged := cmp(pattern, start, end, smudge); equal && end-start == 1 {
			if smudged {
				smudgedAmount++
			}
			matches++
			if smudge {
				if smudgedAmount == 1 {
					return []int{start, end, matches}
				}
				return nil
			} else {
				return []int{start, end, matches}
			}
		}
		for prev := start; prev < end; prev++ {
			for next := end; next > prev; next-- {
				if equal, smudged := cmp(pattern, prev, next, smudge); equal {
					if smudged {
						smudgedAmount++
					}
					matches++
					if next-prev != 1 {
						ret := mirror(cmp, prev+1, next-1, smudge)
						if ret == nil {
							matches = 0
							smudgedAmount = 0
							continue
						}
						return ret
					} else {
						if smudge {
							if smudgedAmount == 1 {
								return []int{prev, next, matches}
							}
							matches = 0
							smudgedAmount = 0
							continue
						} else {
							return []int{prev, next, matches}
						}
					}
				} else if matches != 0 {
					return nil
				}
			}
		}
		return nil
	}

	reflections := func(cmp comparator, to int, smudge bool) int {
		start = 0
		for {
			matches = 0
			smudgedAmount = 0
			line := mirror(cmp, start, to-1, smudge)
			if line != nil {
				start = line[0]
				end = line[1]
				size = line[2]
				if end+size == to || start-size == -1 {
					return end
				}
				start = end
				continue
			}
			break
		}
		return 0
	}

	for i, patterns := range input {
		_ = i
		pattern = nil
		for _, p := range strings.Split(patterns, "\n") {
			var row []rune
			for _, pp := range p {
				row = append(row, pp)
			}
			pattern = append(pattern, row)
		}
		sum1 += reflections(columnCmp, len(pattern[0]), false) + 100*reflections(rowCmp, len(pattern), false)
		sum2 += reflections(columnCmp, len(pattern[0]), true) + 100*reflections(rowCmp, len(pattern), true)
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}
