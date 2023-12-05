package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var NumStringReplacements = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var (
		sum1 int
		sum2 int
	)
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Text()
		s1, e := extractCalibration([]byte(l))
		if e != nil {
			panic(e)
		}
		s2, e := findCalibration(l)
		if e != nil {
			panic(e)
		}
		sum1 += s1
		sum2 += s2
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}

func wordToNumber(word string) string {
	for k, v := range NumStringReplacements {
		word = strings.ReplaceAll(word, k, v)
	}
	return word
}

func extractCalibration(data []byte) (int, error) {
	var (
		start, end byte
		nums       []byte
	)
	for _, c := range []byte(data) {
		if c >= 48 && c <= 57 {
			nums = append(nums, c)
		}
	}

	if len(nums) > 0 {
		start = nums[0]
		end = start
		if len(nums) > 1 {
			end = nums[len(nums)-1]
		}
	} else {
		return 0, nil
	}

	return strconv.Atoi(string([]byte{start, end}))
}

func findCalibration(data string) (int, error) {
	data = wordToNumber(data)
	return extractCalibration([]byte(data))
}
