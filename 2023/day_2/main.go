package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	RED_LIMIT   = 12
	GREEN_LIMIT = 13
	BLUE_LIMIT  = 14
)

type GameSet struct {
	red   int
	green int
	blue  int
}

func (gs *GameSet) isPossible() bool {
	return gs.red <= RED_LIMIT && gs.green <= GREEN_LIMIT && gs.blue <= BLUE_LIMIT
}

func (gs *GameSet) String() string {
	return fmt.Sprintf("red %d, green %d, blue %d", gs.red, gs.green, gs.blue)
}

type Game struct {
	id   int
	sets []GameSet
}

func (g *Game) String() string {
	var sets []string
	for _, s := range g.sets {
		sets = append(sets, fmt.Sprintf("%s", s.String()))
	}
	return fmt.Sprintf("id: %d, sets: %s", g.id, strings.Join(sets, "; "))
}

func parseColor(re *regexp.Regexp, s string) int {
	color := re.FindStringSubmatch(s)
	if color == nil {
		return 0
	}
	if len(color) != 2 {
		panic(fmt.Sprintf("Wrong color format: '%s'", s))
	}
	amount, err := strconv.Atoi(color[1])
	if err != nil {
		panic(err)
	}
	return amount
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var (
		sum   int
		games []Game
	)
	re := regexp.MustCompile(`Game ([0-9]+):(.*)`)
	reRed := regexp.MustCompile(`([0-9]+) red`)
	reBlue := regexp.MustCompile(`([0-9]+) blue`)
	reGreen := regexp.MustCompile(`([0-9]+) green`)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		data := re.FindStringSubmatch(line)
		if len(data) != 3 {
			panic(fmt.Sprintf("Wrong game input format: '%s'", line))
		}
		id, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}
		g := Game{id: id}
		sets := strings.Split(strings.TrimSpace(data[2]), ";")
		for _, s := range sets {
			red := parseColor(reRed, s)
			blue := parseColor(reBlue, s)
			green := parseColor(reGreen, s)
			g.sets = append(g.sets, GameSet{red, green, blue})
		}
		games = append(games, g)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	var skip bool
	for _, g := range games {
		for _, s := range g.sets {
			if !s.isPossible() {
				skip = true
				break
			}
		}
		if !skip {
			sum += g.id
		}
		skip = false
	}

	fmt.Println(sum)

	var minRed, minBlue, minGreen int
	sum = 0
	for _, g := range games {
		minRed, minBlue, minGreen = 0, 0, 0
		for _, s := range g.sets {
			minRed = max(minRed, s.red)
			minBlue = max(minBlue, s.blue)
			minGreen = max(minGreen, s.green)
		}
		sum += minRed * minBlue * minGreen
	}
	fmt.Println(sum)
}
