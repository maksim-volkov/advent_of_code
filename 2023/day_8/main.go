package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	left  = "L"
	right = "R"
)

type Tuple[T1, T2 any] struct {
	left  T1
	right T2
}

func NewTuple[T1, T2 any](a T1, b T2) *Tuple[T1, T2] {
	return &Tuple[T1, T2]{
		left:  a,
		right: b,
	}
}

func (t *Tuple[T1, T2]) Left() T1 {
	return t.left
}

func (t *Tuple[T1, T2]) Right() T2 {
	return t.right
}

func (t *Tuple[T1, T2]) String() string {
	var r []string
	r = append(r, fmt.Sprintf("%#v", t.left))
	r = append(r, fmt.Sprintf("%#v", t.right))
	return "(" + strings.Join(r, ", ") + ")"
}

// Calculate the greatest common divisor (GCD) using Euclid's algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Calculate the least common multiple (LCM) of two integers
func lcm(a, b int) int {
	// Check for zero values to avoid division by zero
	if a == 0 || b == 0 {
		return 0
	}

	// Calculate LCM using the formula
	return (a * b) / gcd(a, b)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	re := regexp.MustCompile(`(.*) = \((.*), (.*)\)`)

	var directions []string
	for _, d := range strings.Split(lines[0], "") {
		directions = append(directions, d)
	}

	directionsMap := make(map[string]*Tuple[string, string])
	for _, d := range lines[2:] {
		match := re.FindStringSubmatch(d)
		if len(match) != 4 {
			panic(fmt.Sprintf("Wrong map enty format: '%s'", d))
		}
		directionsMap[match[1]] = NewTuple(match[2], match[3])
	}

	find := func(originalStart, stop string) int {
		var (
			steps    int
			end      bool
			next     *Tuple[string, string]
			allSteps []int
		)
		for direction := range directionsMap {
			if strings.HasSuffix(direction, originalStart) {
				start := direction
				steps = 0
				end = false
				for !end {
					for _, d := range directions {
						if strings.HasSuffix(start, stop) {
							end = true
							allSteps = append(allSteps, steps)
							break
						}

						next = directionsMap[start]

						switch d {
						case left:
							start = next.Left()
						case right:
							start = next.Right()
						default:
							panic(fmt.Sprintf("Wrong direction %s", d))
						}

						steps++
					}
				}
			}
		}
		for _, s := range allSteps {
			steps = lcm(steps, s)
		}
		return steps
	}

	steps := find("AAA", "ZZZ")

	fmt.Println(steps)

	steps = find("A", "Z")

	fmt.Println(steps)
}
