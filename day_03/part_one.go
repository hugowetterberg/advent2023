package day_03

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"strconv"

	"github.com/hugowetterberg/advent2023"
)

const (
	asciiZero = 48
	asciiNine = 57
)

type number struct {
	N      int
	Row    int
	Column int
	Length int
}

func (n number) AdjacentTo(row, col int) bool {
	return row >= n.Row-1 && row <= n.Row+1 &&
		col >= n.Column-1 && col <= n.Column+n.Length
}

func PartOne(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error) {
	scanner := bufio.NewScanner(r)

	res := advent2023.IntegerResult{
		Description: "Sum of parts",
	}

	var (
		i         int
		columns   int
		schematic [][]byte
		numbers   []number
	)

	for scanner.Scan() {
		n, row, nums, err := parseRow(scanner.Bytes(), i, columns)
		if err != nil {
			return nil, fmt.Errorf("invalid row at line %d: %w",
				i+1, err)
		}

		columns = n
		schematic = append(schematic, row)
		numbers = append(numbers, nums...)

		i++
	}

	for _, num := range numbers {
		adjacent := adjacentToSymbols(num, schematic)
		if adjacent > 0 {
			res.N += num.N
		}
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &res, nil
}

func adjacentToSymbols(num number, schematic [][]byte) int {
	rows := len(schematic)
	if rows == 0 {
		return 0
	}

	cols := len(schematic[0])
	if cols == 0 {
		return 0
	}

	var hits int

	farCol := min(num.Column+num.Length, cols-1)
	farRow := min(num.Row+1, rows-1)

	for col := max(num.Column-1, 0); col <= farCol; col++ {
		for row := max(num.Row-1, 0); row <= farRow; row++ {
			b := schematic[row][col]
			if (b >= asciiZero && b <= asciiNine) || b == '.' {
				continue
			}

			hits++
		}
	}

	return hits
}

func parseRow(line []byte, row int, columns int) (int, []byte, []number, error) {
	data := make([]byte, len(line))

	length := copy(data, line)
	if columns != 0 && columns != length {
		return 0, nil, nil, errors.New("inconsistent line length")
	}

	var (
		num     *number
		numbers []number
		buf     []byte
	)

	push := func() error {
		if num == nil {
			return nil
		}

		n, err := strconv.Atoi(string(buf))
		if err != nil {
			return fmt.Errorf(
				"invalid number at column %d: %w",
				num.Column, err)
		}

		num.N = n
		num.Length = len(buf)

		numbers = append(numbers, *num)

		num = nil
		buf = buf[0:0]

		return nil
	}

	for i, b := range data {
		if b < asciiZero || b > asciiNine {
			err := push()
			if err != nil {
				return 0, nil, nil, err
			}

			continue
		}

		if num == nil {
			num = &number{
				Row:    row,
				Column: i,
			}
		}

		buf = append(buf, b)
	}

	err := push()
	if err != nil {
		return 0, nil, nil, err
	}

	return length, data, numbers, nil
}
