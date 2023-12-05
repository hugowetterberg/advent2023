package day_05

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"slices"
	"strconv"

	"github.com/hugowetterberg/advent2023"
)

func PartOne(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	almanac, err := parseAlmanac(r)
	if err != nil {
		return nil, fmt.Errorf("parse almanac: %w", err)
	}

	res := advent2023.IntegerResult{
		Description: "Lowest location",
	}

	locations, err := almanac.SeedsTo("location")
	if err != nil {
		return nil, fmt.Errorf("map to location: %w", err)
	}

	for i, n := range locations {
		if i == 0 {
			res.N = n

			continue
		}

		res.N = min(res.N, n)
	}

	return &res, nil
}

type almanac struct {
	Seeds       []int
	ForwardMap  map[string]mapping
	BackwardMap map[string]mapping
}

func (a almanac) SeedsTo(name string) ([]int, error) {
	var (
		chain []mapping
		end   mapping
	)

	end, ok := a.BackwardMap[name]
	if !ok {
		return nil, fmt.Errorf("unknown destination %q", name)
	}

	current := end

	for {
		chain = append(chain, current)

		if current.Source == "seed" {
			break
		}

		next, ok := a.BackwardMap[current.Source]
		if !ok {
			return nil, fmt.Errorf("missing mappings for %q",
				current.Source)
		}

		current = next
	}

	slices.Reverse(chain)

	numbers := make([]int, len(a.Seeds))
	copy(numbers, a.Seeds)

	for _, m := range chain {
		for i, n := range numbers {
			numbers[i] = m.Map(n)
		}
	}

	return numbers, nil
}

func (a almanac) SeedTo(seed int, name string) ([]int, error) {
	var (
		chain []mapping
		end   mapping
	)

	end, ok := a.BackwardMap[name]
	if !ok {
		return nil, fmt.Errorf("unknown destination %q", name)
	}

	current := end

	for {
		chain = append(chain, current)

		if current.Source == "seed" {
			break
		}

		next, ok := a.BackwardMap[current.Source]
		if !ok {
			return nil, fmt.Errorf("missing mappings for %q",
				current.Source)
		}

		current = next
	}

	slices.Reverse(chain)

	numbers := make([]int, len(a.Seeds))
	copy(numbers, a.Seeds)

	for _, m := range chain {
		for i, n := range numbers {
			numbers[i] = m.Map(n)
		}
	}

	return numbers, nil
}

func (a almanac) MappingChain(name string) (func(int) int, error) {
	var (
		chain []func(int) int
		end   mapping
	)

	end, ok := a.BackwardMap[name]
	if !ok {
		return nil, fmt.Errorf("unknown destination %q", name)
	}

	current := end

	for {
		chain = append(chain, current.Map)

		if current.Source == "seed" {
			break
		}

		next, ok := a.BackwardMap[current.Source]
		if !ok {
			return nil, fmt.Errorf("missing mappings for %q",
				current.Source)
		}

		current = next
	}

	slices.Reverse(chain)

	return func(n int) int {
		for _, fn := range chain {
			n = fn(n)
		}

		return n
	}, nil
}

type mapping struct {
	Source      string
	Destination string
	Ranges      []mappingRange
}

func (m mapping) Map(n int) int {
	for _, r := range m.Ranges {
		if n < r.SourceStart || n >= r.SourceStart+r.Len {
			continue
		}

		return r.DestStart + n - r.SourceStart
	}

	return n
}

type mappingRange struct {
	DestStart   int
	SourceStart int
	Len         int
}

type state int

const (
	stateNone state = iota
	stateAwaitMap
	stateMap
)

func parseAlmanac(r io.Reader) (*almanac, error) {
	scanner := bufio.NewScanner(r)

	var (
		s              state
		seeds          []int
		currentMapping *mapping
	)

	forwardMap := make(map[string]mapping)
	backwardMap := make(map[string]mapping)

	push := func() {
		if currentMapping == nil {
			return
		}

		forwardMap[currentMapping.Source] = *currentMapping
		backwardMap[currentMapping.Destination] = *currentMapping

		currentMapping = nil
	}

	for scanner.Scan() {
		switch s {
		case stateNone:
			list, err := parseSeeds(scanner.Bytes())
			if err != nil {
				return nil, fmt.Errorf(
					"failed to parse seeds: %w", err)
			}

			seeds = list
			s = stateAwaitMap
		case stateAwaitMap:
			// Ignore blank lines
			if len(scanner.Bytes()) == 0 {
				break
			}

			m, err := parseMappingHeader(scanner.Bytes())
			if err != nil {
				return nil, fmt.Errorf(
					"failed to parse mapping header: %w", err)
			}

			s = stateMap
			currentMapping = m
		case stateMap:
			if len(scanner.Bytes()) == 0 {
				push()

				s = stateAwaitMap

				break
			}

			r, err := parseMappingRange(scanner.Bytes())
			if err != nil {
				return nil, fmt.Errorf(
					"failed to parse mapping range: %w", err)
			}

			currentMapping.Ranges = append(currentMapping.Ranges, r)
		}
	}

	push()

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &almanac{
		Seeds:       seeds,
		ForwardMap:  forwardMap,
		BackwardMap: backwardMap,
	}, nil
}

var (
	seedPrefix = []byte("seeds: ")
	spaceSep   = []byte(" ")
	mapSuffix  = []byte(" map:")
	mapDirSep  = []byte("-to-")
)

func parseSeeds(b []byte) ([]int, error) {
	numberList, ok := bytes.CutPrefix(b, seedPrefix)
	if !ok {
		return nil, fmt.Errorf("expected seed list to start with %q",
			string(seedPrefix))
	}

	var seeds []int

	numberData := bytes.Split(numberList, spaceSep)
	for _, d := range numberData {
		n, err := strconv.Atoi(string(d))
		if err != nil {
			return nil, fmt.Errorf("invalid seed number: %w",
				err)
		}

		seeds = append(seeds, n)
	}

	return seeds, nil
}

func parseMappingHeader(b []byte) (*mapping, error) {
	directionLine, ok := bytes.CutSuffix(b, mapSuffix)
	if !ok {
		return nil, fmt.Errorf("expected map header to end with %q",
			string(mapSuffix))
	}

	source, destination, ok := bytes.Cut(directionLine, mapDirSep)
	if !ok {
		return nil, fmt.Errorf("expected map source and destination to be separated by %q",
			string(mapDirSep))
	}

	return &mapping{
		Source:      string(source),
		Destination: string(destination),
	}, nil
}

func parseMappingRange(b []byte) (mappingRange, error) {
	numberData := bytes.Split(b, spaceSep)
	if len(numberData) != 3 {
		return mappingRange{}, errors.New(
			"expected mapping to consist of three numbers")
	}

	var numbers [3]int

	for i, d := range numberData {
		n, err := strconv.Atoi(string(d))
		if err != nil {
			return mappingRange{}, fmt.Errorf("invalid mapping number: %w",
				err)
		}

		numbers[i] = n
	}

	return mappingRange{
		DestStart:   numbers[0],
		SourceStart: numbers[1],
		Len:         numbers[2],
	}, nil
}
