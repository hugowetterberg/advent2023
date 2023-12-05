package day_04_test

import (
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/hugowetterberg/advent2023"
	"github.com/hugowetterberg/advent2023/day_04"
	"github.com/hugowetterberg/advent2023/internal"
)

const sampleInput = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

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
			Fn:     day_04.PartOne,
			Input:  sampleInputReader,
			Expect: 13,
		},
		"One": {
			Fn:     day_04.PartOne,
			Input:  fullInputReader,
			Expect: 19855,
		},
		"TwoSample": {
			Fn:     day_04.PartTwo,
			Input:  sampleInputReader,
			Expect: 30,
		},
		"Two": {
			Fn:     day_04.PartTwo,
			Input:  fullInputReader,
			Expect: 10378710,
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
