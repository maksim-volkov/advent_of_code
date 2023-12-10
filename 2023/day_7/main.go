package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	cardsWeights map[string]int
	part2        bool
)

func setup() {
	cardsWeights = nil
	var cards []string
	if part2 {
		cards = []string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}
	} else {
		cards = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
	}
	cardsWeights = make(map[string]int)
	for w, c := range cards {
		cardsWeights[c] = w
	}
}

type HandType uint8

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfKind
	FullHouse
	FourOfKind
	FiveOfKind
)

type Hand struct {
	t     HandType
	cards []string
	bid   int
}

type Hands []Hand

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hands) Less(i, j int) bool {
	if h[i].t != h[j].t {
		return h[i].t < h[j].t
	} else {
		for k := 0; k < len(h[i].cards); k++ {
			if cardsWeights[h[i].cards[k]] != cardsWeights[h[j].cards[k]] {
				return cardsWeights[h[i].cards[k]] < cardsWeights[h[j].cards[k]]
			}
		}
	}
	return false
}

func NewHand(s string) Hand {
	r := strings.Split(s, " ")
	if len(r) != 2 {
		panic(fmt.Sprintf("Wrong format for hand: ''", s))
	}
	bid := toNum(r[1])
	cards := strings.Split(r[0], "")
	t := handType(cards)

	return Hand{t, cards, bid}
}

func handType(cards []string) HandType {
	cardsCount := make(map[string]int)
	for _, c := range cards {
		if _, ok := cardsCount[c]; ok {
			cardsCount[c] += 1
		} else {
			cardsCount[c] = 1
		}
	}

	if part2 {
		amountOfJokers := cardsCount["J"]
		delete(cardsCount, "J")
		var (
			m    int
			card string
		)
		for k, v := range cardsCount {
			if v > m {
				m = v
				card = k
			}
		}
		cardsCount[card] += amountOfJokers
	}

	var t HandType
	switch len(cardsCount) {
	case 1:
		t = FiveOfKind
	case 2:
		for _, v := range cardsCount {
			if v == 1 || v == 4 {
				t = FourOfKind
			} else if v == 2 || v == 3 {
				t = FullHouse
			}
			break
		}
	case 3:
		for _, v := range cardsCount {
			if v == 3 {
				t = ThreeOfKind
			} else if v == 2 {
				t = TwoPair
			}
		}
	case 4:
		t = OnePair
	case 5:
		t = HighCard
	default:
		panic("Unknown hand type!!!")
	}
	return t
}

func toNum(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var (
		hands Hands
		sum   int
	)

	// Part one.
	setup()
	for _, l := range lines {
		hands = append(hands, NewHand(l))
	}

	sort.Sort(hands)

	for i, h := range hands {
		sum += (i + 1) * h.bid
	}

	fmt.Println(sum)

	// Part two.
	hands = nil
	part2 = true
	setup()
	for _, l := range lines {
		hands = append(hands, NewHand(l))
	}

	sort.Sort(hands)

	sum = 0
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}

	fmt.Println(sum)
}
