package day_04

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"

	"github.com/hugowetterberg/advent2023"
)

func PartTwo(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	scanner := bufio.NewScanner(r)

	res := advent2023.IntegerResult{
		Description: "Total scratch cards",
	}

	var (
		i     int
		cards []scratchCard
	)

	for scanner.Scan() {
		card, err := parseRow(scanner.Bytes())
		if err != nil {
			return nil, fmt.Errorf("invalid row at line %d: %w",
				i+1, err)
		}

		cards = append(cards, *card)

		i++
	}

	for i := 0; i < len(cards); i++ {
		res.N += calculateScore(cards, i)
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &res, nil
}

func calculateScore(cards []scratchCard, idx int) int {
	if cards[idx].WinCount == 0 {
		return 1
	}

	var (
		bound = min(len(cards), idx+cards[idx].WinCount+1)
		score = 1
	)

	for i := idx + 1; i < bound; i++ {
		score += calculateScore(cards, i)
	}

	return score
}
