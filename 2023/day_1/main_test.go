package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordToNumber(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		in  string
		out string
	}{
		{"", ""},
		{"five6fivethree2three", "f5e6f5et3e2t3e"},
		{"326sevenfivenseven1kctgmnqtwonefq", "326s7nf5ens7n1kctgmnqt2o1efq"},
		{"sone1oneight", "so1e1o1e8t"},
		{"two1twooo", "t2o1t2ooo"},
		{"two1nine", "t2o1n9e"},
		{"eightwothree", "e8t2ot3e"},
		{"abcone2threexyz", "abco1e2t3exyz"},
		{"xtwone3four", "xt2o1e3f4r"},
		{"4nineeightseven2", "4n9ee8ts7n2"},
		{"zoneight234", "zo1e8t234"},
		{"7pqrstsixteen", "7pqrsts6xteen"},
	}

	for _, c := range cases {
		n := wordToNumber(c.in)
		assert.Equal(c.out, n)
	}
}

func TestFindCalibration(t *testing.T) {

	assert := assert.New(t)

	cases := []struct {
		in  string
		out int
		err error
	}{
		{"", 0, nil},
		{"1", 11, nil},
		{"seven", 77, nil},
		{"326sevenfivenseven1kctgmnqtwonefq", 31, nil},
		{"dfsdfonesdfs", 11, nil},
		{"two1nine", 29, nil},
		{"eightwothree", 83, nil},
		{"abcone2threexyz", 13, nil},
		{"xtwone3four", 24, nil},
		{"4nineeightseven2", 42, nil},
		{"zoneight234", 14, nil},
		{"7pqrstsixteen", 76, nil},
	}

	for _, c := range cases {
		n, e := findCalibration(c.in)
		assert.Equal(c.err, e)
		assert.Equal(c.out, n)
	}
}
