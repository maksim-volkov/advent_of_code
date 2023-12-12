package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strings"
)

var (
	charConnections map[rune]map[image.Point]bool
	up              = image.Point{0, -1}
	down            = image.Point{0, 1}
	left            = image.Point{-1, 0}
	right           = image.Point{1, 0}
)

func init() {
	charConnections = make(map[rune]map[image.Point]bool)
	// Items in list represent connection presensce in {up, down, left, right}.
	charConnections['|'] = map[image.Point]bool{up: true, down: true, left: false, right: false}
	charConnections['-'] = map[image.Point]bool{up: false, down: false, left: true, right: true}
	charConnections['L'] = map[image.Point]bool{up: true, down: false, left: false, right: true}
	charConnections['J'] = map[image.Point]bool{up: true, down: false, left: true, right: false}
	charConnections['7'] = map[image.Point]bool{up: false, down: true, left: true, right: false}
	charConnections['F'] = map[image.Point]bool{up: false, down: true, left: false, right: true}
	charConnections['S'] = map[image.Point]bool{up: true, down: true, left: true, right: true}
	charConnections['.'] = map[image.Point]bool{up: false, down: false, left: false, right: false}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var (
		start     image.Point
		next      image.Point
		foundNext bool
		path      []image.Point
		area      int
		walk      func(image.Point, image.Point, image.Point) int
	)

	grid := make(map[image.Point]map[image.Point]bool)
	for y, l := range lines {
		for x, r := range l {
			if r == 'S' {
				start = image.Point{x, y}
			}
			grid[image.Point{x, y}] = charConnections[r]
		}
	}

	// Return value 0 meand that loop was found, -1 means deadend.
	walk = func(previousStart, _start, end image.Point) int {
		if _start == end {
			return 0
		}
		directions, ok := grid[_start]
		if !ok {
			return -1
		}
		for correction, ok := range directions {
			if ok {
				nextStart := _start.Add(correction)
				// Do not go back.
				if nextStart == previousStart {
					continue
				}
				// Calculate area of the loop.
				area += _start.X*nextStart.Y - _start.Y*nextStart.X

				result := walk(_start, nextStart, end)
				if result == 0 {
					path = append(path, _start)
					break
				}
			}
		}
		return 0
	}

	// Since start can connect to any other location, do the first step here to find correct next symbol
	// that can be connected to start if moving backwards.
	for correction, ok := range grid[start] {
		if foundNext {
			break
		}
		if ok {
			next = start.Add(correction)
			for nextCorrection, ok := range grid[next] {
				if ok {
					newNext := next.Add(nextCorrection)
					if newNext == start {
						foundNext = true
						break
					}
				}
			}
		}
	}
	path = append(path, start)
	area += start.X*next.Y - start.Y*next.X

	// Calculate all the next steps in the loop starting not from S, but from next valid character.
	walk(start, next, start)
	fmt.Println(len(path) / 2)
	fmt.Println((int(math.Abs(float64(area)))-len(path))/2 + 1)
}
