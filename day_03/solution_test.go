package day_03_test

import (
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/hugowetterberg/advent2023"
	"github.com/hugowetterberg/advent2023/day_03"
	"github.com/hugowetterberg/advent2023/internal"
)

const sampleInput = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

type testCase struct {
	Fn     func(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error)
	Input  func() io.Reader
	Expect int
}

func TestParts(t *testing.T) {
	fullInput, err := os.ReadFile("input.txt")
	if err != nil {
		t.Fatalf("failed to read input file: %v", err)
	}

	sampleInputReader := internal.ByteReaderFunc([]byte(sampleInput))
	fullInputReader := internal.ByteReaderFunc(fullInput)

	parts := map[string]testCase{
		"OneSample": {
			Fn:     day_03.PartOne,
			Input:  sampleInputReader,
			Expect: 4361,
		},
		"One": {
			Fn:     day_03.PartOne,
			Input:  fullInputReader,
			Expect: 533775,
		},
		"TwoSample": {
			Fn:     day_03.PartTwo,
			Input:  sampleInputReader,
			Expect: 467835,
		},
		"Two": {
			Fn:     day_03.PartTwo,
			Input:  fullInputReader,
			Expect: 78236071,
		},
	}

	for name, c := range parts {
		t.Run(name, func(t *testing.T) {
			logger := slog.New(internal.NewLogHandler(t, slog.LevelWarn))

			res, err := c.Fn(logger, c.Input())
			if err != nil {
				t.Fatal(err.Error())
			}

			if res.N != c.Expect {
				t.Errorf("got %d, expected %d", res.N, c.Expect)
			}
		})
	}
}
