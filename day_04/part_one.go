package day_04

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
	scanner := bufio.NewScanner(r)

	res := advent2023.IntegerResult{
		Description: "Total points",
	}

	var i int

	for scanner.Scan() {
		card, err := parseRow(scanner.Bytes())
		if err != nil {
			return nil, fmt.Errorf("invalid row at line %d: %w",
				i+1, err)
		}

		if card.WinCount == 0 {
			continue
		}

		res.N += 1 << (card.WinCount - 1)

		i++
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &res, nil
}

type scratchCard struct {
	ID             int
	WinningNumbers []int
	Numbers        []int
	WinCount       int
}

func WinCount(winners, numbers []int) int {
	var c int

	for _, n := range numbers {
		if !slices.Contains(winners, n) {
			continue
		}

		c++
	}

	return c
}

var (
	preamble  = []byte("Card ")
	headerSep = []byte(": ")
)

func parseRow(line []byte) (*scratchCard, error) {
	line, ok := bytes.CutPrefix(line, preamble)
	if !ok {
		return nil, fmt.Errorf("expected the line to start with %q",
			string(preamble))
	}

	idData, numbers, ok := bytes.Cut(line, headerSep)
	if !ok {
		return nil, fmt.Errorf(
			"expected a header and numbers separated by %q", headerSep)
	}

	idData = bytes.TrimSpace(idData)

	id, err := strconv.Atoi(string(idData))
	if err != nil {
		return nil, fmt.Errorf("invalid card ID: %w", err)
	}

	card := scratchCard{
		ID: id,
	}

	var (
		buf        []byte
		numSection int
	)

	push := func() error {
		if len(buf) == 0 {
			return nil
		}

		n, err := strconv.Atoi(string(buf))
		if err != nil {
			return fmt.Errorf("invalid card number: %w", err)
		}

		buf = buf[0:0]

		switch numSection {
		case 0:
			card.WinningNumbers = append(card.WinningNumbers, n)
		case 1:
			card.Numbers = append(card.Numbers, n)
		default:
			return errors.New("too many card section separators ('|')")
		}

		return nil
	}

	for _, b := range numbers {
		switch b {
		case ' ':
			err := push()
			if err != nil {
				return nil, err
			}
		case '|':
			err := push()
			if err != nil {
				return nil, err
			}

			numSection++
		default:
			buf = append(buf, b)
		}
	}

	err = push()
	if err != nil {
		return nil, err
	}

	card.WinCount = WinCount(card.WinningNumbers, card.Numbers)

	return &card, nil
}
