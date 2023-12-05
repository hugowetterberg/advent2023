package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/hugowetterberg/advent2023"
	"github.com/hugowetterberg/advent2023/day_01"
	"github.com/hugowetterberg/advent2023/day_02"
	"github.com/hugowetterberg/advent2023/day_03"
)

type integerFunc func(logger *slog.Logger, r io.Reader) (*advent2023.IntegerResult, error)

func main() {
	solutions := map[string]integerFunc{
		"day_1_1": day_01.PartOne,
		"day_1_2": day_01.PartTwo,
		"day_2_1": day_02.PartOne,
		"day_2_2": day_02.PartTwo,
		"day_3_1": day_03.PartOne,
		"day_3_2": day_03.PartTwo,
	}

	printSolutionNames := func() {
		for name := range solutions {
			fmt.Println("*", name)
		}
	}

	if len(os.Args) < 2 {
		fmt.Println("please supply a solution name, one of:")
		printSolutionNames()
		os.Exit(1)
	}

	fn, ok := solutions[os.Args[1]]
	if !ok {
		fmt.Println("please supply a valid solution name, one of:")
		printSolutionNames()
		os.Exit(1)
	}

	res, err := fn(slog.Default(), os.Stdin)
	if err != nil {
		fmt.Printf("failed to run: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s: %d\n", res.Description, res.N)
}
