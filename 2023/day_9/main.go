package main

import (
	"fmt"
	"os"
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

func predict(in string) (int, int) {
	var (
		history                  [][]int
		start                    []int
		stop                     bool
		next                     []int
		prediction1, prediction2 int
	)
	for _, r := range strings.Split(in, " ") {
		start = append(start, toNum(string(r)))
	}

	history = append(history, start)

	for !stop {
		next = nil
		for i := 0; i < len(start)-1; i++ {
			diff := start[i+1] - start[i]
			next = append(next, diff)
		}
		stop = true
		for _, n := range next {
			if n != 0 {
				stop = false
				break
			}
		}
		if stop {
			next = append(next, 0)
			next = append([]int{0}, next...)
		}
		start = next
		history = append(history, next)
	}

	for i := len(history) - 1; i > 0; i-- {
		o := history[i][len(history[i])-1]
		n := history[i-1][len(history[i-1])-1]
		prediction1 = n + o
		history[i-1] = append(history[i-1], prediction1)
	}

	for i := len(history) - 1; i > 0; i-- {
		o := history[i][0]
		n := history[i-1][0]
		prediction2 = n - o
		history[i-1] = append([]int{prediction2}, history[i-1]...)
	}

	return prediction1, prediction2
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var (
		wg         sync.WaitGroup
		sum1, sum2 atomic.Int64
	)
	for _, l := range lines {
		wg.Add(1)
		go func(in string) {
			defer wg.Done()
			s1, s2 := predict(in)
			sum1.Add(int64(s1))
			sum2.Add(int64(s2))
		}(l)
	}

	wg.Wait()

	fmt.Println(sum1.Load())
	fmt.Println(sum2.Load())
}
