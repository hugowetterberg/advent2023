package day_01_test

import (
	"log/slog"
	"strings"
	"testing"

	"github.com/hugowetterberg/advent2023/day_01"
	"github.com/hugowetterberg/advent2023/internal"
)

func TestPartOne(t *testing.T) {
	logger := slog.New(internal.NewLogHandler(t, slog.LevelWarn))

	in := strings.NewReader(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

	res, err := day_01.PartOne(logger, in)
	if err != nil {
		t.Fatal(err.Error())
	}

	if res.N != 142 {
		t.Errorf("got %d, expected %d", res.N, 142)
	}
}

func TestPartTwo(t *testing.T) {
	logger := slog.New(internal.NewLogHandler(t, slog.LevelWarn))

	in := strings.NewReader(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)

	res, err := day_01.PartTwo(logger, in)
	if err != nil {
		t.Fatal(err.Error())
	}

	if res.N != 281 {
		t.Errorf("got %d, expected %d", res.N, 281)
	}
}
