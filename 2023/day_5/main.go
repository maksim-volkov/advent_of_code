package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

func toNum(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func findLocation(location int, maps [][][3]int) int {
	for _, mapEntries := range maps {
		for _, entry := range mapEntries {
			if entry[1] <= location && location-entry[1] < entry[2] {
				location = location - entry[1] + entry[0]
				break
			}
		}
	}
	return location
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	var (
		seeds []int
		maps  [][][3]int
	)
	reSeeds := regexp.MustCompile(`seeds: (.*)`)

	// Collect seeds numbers.
	matches := reSeeds.FindStringSubmatch(lines[0])
	if len(matches) != 2 {
		panic(fmt.Sprintf("Wrong format for seeds: '%s'", lines[0]))
	}
	for _, s := range strings.Split(matches[1], " ") {
		seeds = append(seeds, toNum(s))
	}
	if len(seeds)%2 != 0 {
		panic(fmt.Sprintf("Uneven amount of seeds: %d", len(seeds)))
	}

	// Collect maps.
	for _, line := range lines[1:] {
		var entries [][3]int
		for _, mapEntry := range strings.Split(line, "\n")[1:] {
			var entry [3]int
			for i, e := range strings.Split(mapEntry, " ") {
				entry[i] = toNum(e)
			}
			entries = append(entries, entry)
		}
		maps = append(maps, entries)
	}

	// Calculate locations.
	lowest1 := math.MaxInt
	lowest2 := atomic.Int64{}
	var wg sync.WaitGroup
	for i, s := range seeds {
		lowest1 = slices.Min([]int{lowest1, findLocation(s, maps)})

		if i%2 == 0 {
			wg.Add(1)
			go func(start, end int) {
				defer wg.Done()
				lowest2.Store(math.MaxInt)
				for s := start; s < end; s++ {
					lowest2.Store(slices.Min([]int64{lowest2.Load(), int64(findLocation(s, maps))}))
				}
			}(seeds[i], seeds[i]+seeds[i+1])
		}
	}

	wg.Wait()

	fmt.Println(lowest1)
	fmt.Println(lowest2.Load())
}
