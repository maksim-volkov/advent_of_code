package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	MULTIPLIER = 17
	DIVIDER    = 256
)

type Lens struct {
	label string
	focal int
}

type Lenses []Lens

func (all *Lenses) find(label string) int {
	for i, l := range *all {
		if l.label == label {
			return i
		}
	}
	return -1
}

func (all *Lenses) replace(lens Lens) {
	if all == nil {
		*all = append(*all, lens)
		return
	}
	index := all.find(lens.label)
	if index >= 0 {
		(*all)[index] = lens
	} else {
		*all = append(*all, lens)
	}
}

func (all *Lenses) remove(lens Lens) {
	if all == nil {
		return
	}
	index := all.find(lens.label)
	if index >= 0 {
		*all = append((*all)[:index], (*all)[index+1:]...)
	}
}

func focusingPower(index int, box Lenses) int {
	var power int
	for i, l := range box {
		power += (index + 1) * (i + 1) * l.focal
	}
	return power
}

func hash(in string) int {
	var current int
	for _, r := range []byte(in) {
		current += int(r)
		current *= MULTIPLIER
		current %= DIVIDER
	}
	return current
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.Split(strings.TrimSpace(string(f)), ",")
	re := regexp.MustCompile(`([a-z]+)([=-])([0-9]*)`)

	var (
		sum1, sum2 int
		lens       Lens
	)
	boxes := make([]Lenses, 256)

	for _, in := range input {
		//Part 1
		sum1 += hash(in)

		//Part 2
		command := re.FindAllStringSubmatch(in, -1)
		if len(command) != 1 {
			panic(fmt.Sprintf("Unexpected command format '%s'", in))
		}
		label := command[0][1]
		boxIndex := hash(label)
		op := command[0][2]
		switch op {
		case "=":
			if len(command[0]) != 4 {
				panic(fmt.Sprintf("Unexpected command format '%s' for operation '='", in))
			}
			focal, err := strconv.Atoi(command[0][3])
			if err != nil {
				panic(fmt.Sprintf("Unexpected focal length value: %v", err))
			}
			lens = Lens{label, focal}
			boxes[boxIndex].replace(lens)
		case "-":
			lens = Lens{label, -1}
			boxes[boxIndex].remove(lens)
		default:
			panic(fmt.Sprintf("Unexpected operation '%s'", op))
		}
	}
	//Part 1
	fmt.Println(sum1)

	//Part 2
	for i, b := range boxes {
		if b != nil && len(b) > 0 {
			sum2 += focusingPower(i, b)
		}
	}
	fmt.Println(sum2)
}
