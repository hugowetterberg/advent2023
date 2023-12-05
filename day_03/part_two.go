package day_03

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

	// We're using an offset to limit how far back we look for adjecent
	// numbers in the numbers slice.
	var numOffset int

	for r := 0; r < len(schematic); r++ {
		row := schematic[r]

		for c := 0; c < len(row); c++ {
			if schematic[r][c] != '*' {
				continue
			}

			adjacent, o := adjacentNumbers(
				r, c, numbers[numOffset:],
			)

			numOffset += o
			if o > 0 {
				println(numOffset)
			}

			if len(adjacent) != 2 {
				continue
			}

			res.N += adjacent[0].N * adjacent[1].N
		}
	}

	err := scanner.Err()
	if err != nil {
		return nil, fmt.Errorf("read input: %w", err)
	}

	return &res, nil
}

func adjacentNumbers(
	row, col int, numbers []number,
) ([]number, int) {
	var res []number

	var offset int

	for i := range numbers {
		if numbers[i].Row < row-1 {
			offset = i
			continue
		}

		if numbers[i].Row > row+1 {
			break
		}

		if !numbers[i].AdjacentTo(row, col) {
			continue
		}

		res = append(res, numbers[i])
	}

	return res, offset
}
