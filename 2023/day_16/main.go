package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

var (
	connections map[rune]map[image.Point]Directions
	up          = image.Point{0, -1}
	down        = image.Point{0, 1}
	left        = image.Point{-1, 0}
	right       = image.Point{1, 0}
)

type Node struct {
	point     image.Point
	direction image.Point
}

type Direction struct {
	point image.Point
	move  bool
}

func (d Direction) from() image.Point {
	switch d.point {
	case up:
		return down
	case down:
		return up
	}
	return d.point
}

type Directions []Direction

func init() {
	connections = make(map[rune]map[image.Point]Directions)
	connections['|'] = map[image.Point]Directions{
		left:  Directions{{up, true}, {down, true}, {left, false}, {right, false}},
		right: Directions{{up, true}, {down, true}, {left, false}, {right, false}},
		up:    Directions{{up, false}, {down, true}, {left, false}, {right, false}},
		down:  Directions{{up, true}, {down, false}, {left, false}, {right, false}},
	}

	connections['-'] = map[image.Point]Directions{
		left:  Directions{{up, false}, {down, false}, {left, true}, {right, false}},
		right: Directions{{up, false}, {down, false}, {left, false}, {right, true}},
		up:    Directions{{up, false}, {down, false}, {left, true}, {right, true}},
		down:  Directions{{up, false}, {down, false}, {left, true}, {right, true}},
	}

	connections['/'] = map[image.Point]Directions{
		left:  Directions{{up, false}, {down, true}, {left, false}, {right, false}},
		right: Directions{{up, true}, {down, false}, {left, false}, {right, false}},
		up:    Directions{{up, false}, {down, false}, {left, true}, {right, false}},
		down:  Directions{{up, false}, {down, false}, {left, false}, {right, true}},
	}

	connections['\\'] = map[image.Point]Directions{
		left:  Directions{{up, true}, {down, false}, {left, false}, {right, false}},
		right: Directions{{up, false}, {down, true}, {left, false}, {right, false}},
		up:    Directions{{up, false}, {down, false}, {left, false}, {right, true}},
		down:  Directions{{up, false}, {down, false}, {left, true}, {right, false}},
	}

	connections['.'] = map[image.Point]Directions{
		left:  Directions{{up, false}, {down, false}, {left, true}, {right, false}},
		right: Directions{{up, false}, {down, false}, {left, false}, {right, true}},
		up:    Directions{{up, false}, {down, true}, {left, false}, {right, false}},
		down:  Directions{{up, true}, {down, false}, {left, false}, {right, false}},
	}
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(f)), "\n")

	contraption := make(map[image.Point]rune)
	var starts []Node
	for y, l := range input {
		for x, r := range l {
			contraption[image.Point{x, y}] = r
			if y == 0 {
				starts = append(starts, Node{image.Point{x, y}, down})
			}
			if x == 0 {
				starts = append(starts, Node{image.Point{x, y}, right})
			}
			if y == len(input)-1 {
				starts = append(starts, Node{image.Point{x, y}, up})
			}
			if x == len(input[0])-1 {
				starts = append(starts, Node{image.Point{x, y}, left})
			}
		}
	}

	bfs := func(start Node) int {
		powered := make(map[image.Point]struct{})
		queue := []Node{start}
		explored := map[Node]struct{}{start: {}}

		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]

			char, ok := contraption[node.point]
			// Out of border.
			if !ok {
				continue
			}
			directions := connections[char]
			nextDirections := directions[node.direction]

			powered[node.point] = struct{}{}

			for _, direction := range nextDirections {
				if direction.move {
					nextDirection := direction.from()
					nextPoint := node.point.Add(direction.point)
					next := Node{nextPoint, nextDirection}

					if _, ok := explored[next]; !ok {
						explored[next] = struct{}{}
						queue = append(queue, next)
					}
				}
			}
		}

		return len(powered)
	}

	// Part 1.
	fmt.Println(bfs(Node{image.Point{0, 0}, right}))

	// Part 2.
	maxPowered := 0
	for _, start := range starts {
		maxPowered = max(maxPowered, bfs(start))
	}
	fmt.Println(maxPowered)
}
