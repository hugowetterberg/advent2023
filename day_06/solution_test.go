package day_06_test

import (
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/hugowetterberg/advent2023"
	"github.com/hugowetterberg/advent2023/day_06"
	"github.com/hugowetterberg/advent2023/internal"
)

const sampleInput = `Time:      7  15   30
Distance:  9  40  200`

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
			Fn:     day_06.PartOne,
			Input:  sampleInputReader,
			Expect: 288,
		},
		"One": {
			Fn:     day_06.PartOne,
			Input:  fullInputReader,
			Expect: 170000,
		},
		"TwoSample": {
			Fn:     day_06.PartTwo,
			Input:  sampleInputReader,
			Expect: 71503,
		},
		"Two": {
			Fn:     day_06.PartTwo,
			Input:  fullInputReader,
			Expect: 20537782,
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
