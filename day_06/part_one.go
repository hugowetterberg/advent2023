package day_06

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strconv"

	"github.com/hugowetterberg/advent2023"
)

var (
	colonSep = []byte(":")
	spaceSep = []byte(" ")
	timeHead = []byte("Time")
	distHead = []byte("Distance")
)

func PartOne(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	res := advent2023.IntegerResult{
		Description: "Error margin",
		N:           1,
	}

	records, err := parseRecords(r)
	if err != nil {
		return nil, err
	}

	for i, t := range records.Time {
		record := records.Distance[i]

		var winCond int

		for ht := 1; ht < t; ht++ {
			travel := ht * (t - ht)

			if travel > record {
				winCond++
			}
		}

		res.N *= winCond
	}

	return &res, nil
}

type races struct {
	Time     []int
	Distance []int
}

func parseRecords(r io.Reader) (*races, error) {
	var (
		linum     int
		times     []int
		distances []int
	)

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		linum++
		if linum > 2 {
			return nil, errors.New(
				"only expected two lines of input")
		}

		heading, numberLine, ok := bytes.Cut(scanner.Bytes(), colonSep)
		if !ok {
			return nil, fmt.Errorf(
				"line %d ins missing a heading", linum)
		}

		switch {
		case bytes.Equal(heading, timeHead):
			n, err := parseNumbers(numberLine)
			if err != nil {
				return nil, fmt.Errorf(
					"invalid times: %w", err)
			}

			times = n
		case bytes.Equal(heading, distHead):
			n, err := parseNumbers(numberLine)
			if err != nil {
				return nil, fmt.Errorf(
					"invalid distances: %w", err)
			}

			distances = n
		default:
			return nil, fmt.Errorf(
				"unknown heading %q", string(heading))
		}
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	if len(times) != len(distances) {
		return nil, errors.New("time and distance count mismatch")
	}

	return &races{
		Time:     times,
		Distance: distances,
	}, nil
}

func parseNumbers(d []byte) ([]int, error) {
	var res []int

	for {
		nd, remainder, _ := bytes.Cut(bytes.TrimSpace(d), spaceSep)

		n, err := strconv.Atoi(string(nd))
		if err != nil {
			return nil, fmt.Errorf("invalid number: %w", err)
		}

		res = append(res, n)

		d = remainder
		if len(d) == 0 {
			break
		}
	}

	return res, nil
}
