package day_05

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"runtime"
	"sync"

	"github.com/hugowetterberg/advent2023"
)

type seedRange struct {
	Start  int
	Length int
}

func PartTwo(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	almanac, err := parseAlmanac(r)
	if err != nil {
		return nil, fmt.Errorf("parse almanac: %w", err)
	}

	if len(almanac.Seeds)%2 != 0 {
		return nil, errors.New("the almanac must contain an even number of seeds")
	}

	chain, err := almanac.MappingChain("location")
	if err != nil {
		return nil, fmt.Errorf("chain to location: %w", err)
	}

	rangeJobs := make(chan seedRange)
	rangeResult := make(chan int)

	var wg sync.WaitGroup

	go func() {
		const batchSize = 1 << 20

		for i := 0; i < len(almanac.Seeds); i += 2 {
			fullRange := seedRange{
				Start:  almanac.Seeds[i],
				Length: almanac.Seeds[i+1],
			}

			for ss := 0; ss < fullRange.Length; ss += batchSize {
				rangeJobs <- seedRange{
					Start:  fullRange.Start + ss,
					Length: min(batchSize, fullRange.Length-ss),
				}
			}
		}

		close(rangeJobs)
	}()

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)

		go func() {
			for r := range rangeJobs {
				if r.Length == 0 {
					continue
				}

				n := math.MaxInt

				for j := 0; j < r.Length; j++ {
					n = min(n, chain(r.Start+j))
				}

				rangeResult <- n
			}

			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(rangeResult)
	}()

	res := advent2023.IntegerResult{
		Description: "Lowest location",
		N:           math.MaxInt,
	}

	for n := range rangeResult {
		res.N = min(res.N, n)
	}

	return &res, nil
}
