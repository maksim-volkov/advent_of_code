package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var (
		universe []image.Point
		steps    []int
		sum      int
	)

	calculate := func(expansion int) {
		deltaY := 0
		for y, line := range lines {
			if !strings.Contains(line, "#") {
				deltaY += expansion
			}
			deltaX := 0
			for x, char := range line {
				var column []byte
				for _, l := range lines {
					column = append(column, l[x])
				}
				if !strings.Contains(string(column), "#") {
					deltaX += expansion
				}

				if char == '#' {
					for _, galaxy := range universe {
						steps = append(steps, int(math.Abs(float64(x+deltaX-galaxy.X))+math.Abs(float64(y+deltaY-galaxy.Y))))
					}
					universe = append(universe, image.Point{x + deltaX, y + deltaY})
				}
			}
		}
	}

	calculate(1)
	for _, s := range steps {
		sum += s
	}

	fmt.Println(sum)

	steps = nil
	universe = nil
	sum = 0

	calculate(999999)
	for _, s := range steps {
		sum += s
	}

	fmt.Println(sum)
}
