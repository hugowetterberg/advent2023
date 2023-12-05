package day_05

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"

	"github.com/hugowetterberg/advent2023"
)

func PartTwo(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	almanac, err := parseAlmanac(r)
	if err != nil {
		return nil, fmt.Errorf("parse almanac: %w", err)
	}

	if len(almanac.Seeds)%2 != 0 {
		return nil, errors.New("the almanac must contain an even number of seeds")
	}

	res := advent2023.IntegerResult{
		Description: "Lowest location",
		N:           math.MaxInt,
	}

	chain, err := almanac.MappingChain("location")
	if err != nil {
		return nil, fmt.Errorf("chain to location: %w", err)
	}

	for i := 0; i < len(almanac.Seeds); i += 2 {
		start := almanac.Seeds[i]
		len := almanac.Seeds[i+1]

		for j := 0; j < len; j++ {
			res.N = min(res.N, chain(start+j))
		}
	}

	return &res, nil
}
