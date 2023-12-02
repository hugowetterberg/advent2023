package day_02

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strconv"
	"strings"

	"github.com/hugowetterberg/advent2023"
)

const (
	colourRed   = "red"
	colourGreen = "green"
	colourBlue  = "blue"
)

func PartOne(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	scanner := bufio.NewScanner(r)

	res := advent2023.IntegerResult{
		Description: "Sum of games",
	}

	var linum int

	counts := map[string]int{
		colourRed:   12,
		colourGreen: 13,
		colourBlue:  14,
	}

	for scanner.Scan() {
		linum++

		game, err := parseGameLine(scanner.Text(), counts)
		if err != nil {
			logger.Warn("malformed line",
				"err", err,
				"line_number", linum)
		}

		if game.Possible {
			res.N += game.ID
		}
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &res, nil
}

func PartTwo(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	scanner := bufio.NewScanner(r)

	res := advent2023.IntegerResult{
		Description: "Sum of games",
	}

	var linum int

	for scanner.Scan() {
		linum++

		game, err := parseGameLine(scanner.Text(), nil)
		if err != nil {
			logger.Warn("malformed line",
				"err", err,
				"line_number", linum)
		}

		maxObserved := make(map[string]int)

		for _, round := range game.Rounds {
			for colour, count := range round {
				maxObserved[colour] = max(maxObserved[colour], count)
			}
		}

		power := 1
		for _, count := range maxObserved {
			power *= count
		}

		res.N += power
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &res, nil
}

type gameData struct {
	ID       int
	Rounds   []map[string]int
	Possible bool
}

func parseGameLine(line string, counts map[string]int) (*gameData, error) {
	heading, data, ok := strings.Cut(line, ": ")
	if !ok {
		return nil, errors.New(
			"expected a heading that ends with ': '")
	}

	format, idString, _ := strings.Cut(heading, " ")

	if format != "Game" {
		return nil, errors.New(
			"expected the line to start with 'Game'")
	}

	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, fmt.Errorf("malformed game ID: %w", err)
	}

	res := gameData{
		ID:       id,
		Possible: true,
	}

	rounds := strings.Split(data, "; ")
	if len(rounds) == 0 {
		return nil, errors.New(
			"expected each game to have one or more rounds separated by '; '")
	}

	for roundN, round := range rounds {
		revealedCubes := strings.Split(round, ", ")

		roundCounts := make(map[string]int)

		for countN, cubes := range revealedCubes {
			nString, colour, ok := strings.Cut(cubes, " ")
			if !ok {
				return nil, fmt.Errorf(
					"round %d, count %d: the cube count must be a number and a colour separated by a space",
					roundN+1, countN+1)
			}

			n, err := strconv.Atoi(nString)
			if err != nil {
				return nil, fmt.Errorf(
					"round %d, count %d: invalid count number: %w",
					roundN+1, countN+1, err)
			}

			if counts != nil {
				maxN, ok := counts[colour]
				if !ok {
					return nil, fmt.Errorf(
						"round %d, count %d: invalid colour %q",
						roundN+1, countN+1, colour)
				}

				if n > maxN {
					res.Possible = false
				}
			}

			roundCounts[colour] = n
		}

		res.Rounds = append(res.Rounds, roundCounts)
	}

	return &res, nil
}
