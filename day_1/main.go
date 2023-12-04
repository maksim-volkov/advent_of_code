package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var NumMap = map[string]byte{
	"one":   byte('1'),
	"two":   byte('2'),
	"three": byte('3'),
	"four":  byte('4'),
	"five":  byte('5'),
	"six":   byte('6'),
	"seven": byte('7'),
	"eight": byte('8'),
	"nine":  byte('9'),
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var sum int
	s := bufio.NewScanner(f)
	for s.Scan() {
		l := s.Bytes()
		fmt.Printf("%s ", string(l))
		s, e := findCalibration(l)
		if e != nil {
			panic(e)
		}
		sum += s
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}

func findCalibration(data []byte) (int, error) {
	var (
		start, end byte
		nums       []byte
		lit        []byte
	)

	for _, c := range data {
		if c >= 48 && c <= 57 {
			nums = append(nums, c)
		}
		if len(lit) == 5 {
			lit = nil
		}
		lit = append(lit, c)
		v, ok := NumMap[string(lit)]
		if ok {
			nums = append(nums, v)
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

	c, err := strconv.Atoi(string([]byte{start, end}))
	if err != nil {
		return 0, err
	}
	fmt.Printf("%v %v\n", []byte{start, end}, c)
	return c, nil
}
