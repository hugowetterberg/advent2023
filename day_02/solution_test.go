package day_02_test

import (
	"io"
	"log/slog"
	"strings"
	"testing"

	"github.com/hugowetterberg/advent2023"
	"github.com/hugowetterberg/advent2023/day_02"
	"github.com/hugowetterberg/advent2023/internal"
)

const sampleInput = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

type testCase struct {
	Fn     func(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error)
	Expect int
}

func TestParts(t *testing.T) {
	parts := map[string]testCase{
		"One": {
			Fn:     day_02.PartOne,
			Expect: 8,
		},
		"Two": {
			Fn:     day_02.PartTwo,
			Expect: 2286,
		},
	}

	for name, c := range parts {
		t.Run(name, func(t *testing.T) {
			logger := slog.New(internal.NewLogHandler(t, slog.LevelWarn))

			res, err := c.Fn(logger, strings.NewReader(sampleInput))
			if err != nil {
				t.Fatal(err.Error())
			}

			if res.N != c.Expect {
				t.Errorf("got %d, expected %d", res.N, c.Expect)
			}
		})
	}
}
