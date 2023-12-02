package day_01

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log/slog"

	"github.com/hugowetterberg/advent2023"
)

const (
	asciiZero = 48
	asciiNine = 57
)

func PartOne(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	lines := bufio.NewScanner(r)

	var (
		total int
		linum int
		tuple [2]int
	)

	for lines.Scan() {
		numbers := tuple[0:0]

		linum++

		for _, b := range lines.Bytes() {
			if b < asciiZero || b > asciiNine {
				continue
			}

			n := int(b) - asciiZero

			if len(numbers) < 2 {
				numbers = append(numbers, n)
			} else {
				numbers[1] = n
			}
		}

		if len(numbers) == 0 {
			logger.Warn("invalid line, no numbers",
				"line_number", linum)

			continue
		}

		total += numbers[0]*10 + numbers[len(numbers)-1]
	}

	err := lines.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &advent2023.IntegerResult{
		Description: "Calibration value sum",
		N:           total,
	}, nil
}

func PartTwo(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	lines := bufio.NewScanner(r)

	var (
		total int
		linum int
		tuple [2]int
	)

	for lines.Scan() {
		numbers := tuple[0:0]

		linum++

		line := lines.Bytes()

		for i := 0; i < len(line); i++ {
			n, ok := number(line[i:])
			if !ok {
				continue
			}

			if len(numbers) < 2 {
				numbers = append(numbers, n)
			} else {
				numbers[1] = n
			}
		}

		if len(numbers) == 0 {
			logger.Warn("invalid line, no numbers",
				"line_number", linum)

			continue
		}

		total += numbers[0]*10 + numbers[len(numbers)-1]
	}

	err := lines.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &advent2023.IntegerResult{
		Description: "Calibration value sum",
		N:           total,
	}, nil
}

var numWords = map[int][]byte{
	0: []byte("zero"),
	1: []byte("one"),
	2: []byte("two"),
	3: []byte("three"),
	4: []byte("four"),
	5: []byte("five"),
	6: []byte("six"),
	7: []byte("seven"),
	8: []byte("eight"),
	9: []byte("nine"),
}

func number(data []byte) (int, bool) {
	if data[0] >= asciiZero && data[0] <= asciiNine {
		return int(data[0]) - asciiZero, true
	}

	for number, word := range numWords {
		if bytes.HasPrefix(data, word) {
			return number, true
		}
	}

	return 0, false
}
