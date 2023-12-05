package day_05_test

import (
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/hugowetterberg/advent2023"
	"github.com/hugowetterberg/advent2023/day_05"
	"github.com/hugowetterberg/advent2023/internal"
)

const sampleInput = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

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
			Fn:     day_05.PartOne,
			Input:  sampleInputReader,
			Expect: 35,
		},
		"One": {
			Fn:     day_05.PartOne,
			Input:  fullInputReader,
			Expect: 218513636,
		},
		"TwoSample": {
			Fn:     day_05.PartTwo,
			Input:  sampleInputReader,
			Expect: 46,
		},
	}

	if !testing.Short() {
		// Brute-forced this one, putting it behind a short guard.
		parts["Two"] = testCase{
			Fn:     day_05.PartTwo,
			Input:  fullInputReader,
			Expect: 81956384,
		}
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
