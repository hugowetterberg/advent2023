package day_06

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"

	"github.com/hugowetterberg/advent2023"
)

func PartTwo(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	res := advent2023.IntegerResult{
		Description: "Error margin",
		N:           1,
	}

	record, err := parseRecord(r)
	if err != nil {
		return nil, err
	}

	var winCond int

	for ht := 1; ht < record.Time; ht++ {
		travel := ht * (record.Time - ht)

		if travel > record.Distance {
			winCond++
		}
	}

	res.N *= winCond

	return &res, nil
}

type race struct {
	Time     int
	Distance int
}

func parseRecord(r io.Reader) (*race, error) {
	var (
		linum    int
		time     int
		distance int
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
			n, err := parseNumber(numberLine)
			if err != nil {
				return nil, fmt.Errorf(
					"invalid times: %w", err)
			}

			time = n
		case bytes.Equal(heading, distHead):
			n, err := parseNumber(numberLine)
			if err != nil {
				return nil, fmt.Errorf(
					"invalid distances: %w", err)
			}

			distance = n
		default:
			return nil, fmt.Errorf(
				"unknown heading %q", string(heading))
		}
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &race{
		Time:     time,
		Distance: distance,
	}, nil
}

func parseNumber(d []byte) (int, error) {
	n, err := strconv.Atoi(strings.ReplaceAll(string(d), " ", ""))
	if err != nil {
		return 0, fmt.Errorf("invalid number: %w", err)
	}

	return n, nil
}
