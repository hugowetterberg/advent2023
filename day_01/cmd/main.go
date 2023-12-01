package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/hugowetterberg/advent2023/day_01"
)

func main() {
	var partTwo bool

	flag.BoolVar(&partTwo, "part-two", false, "run part two")

	flag.Parse()

	fn := day_01.PartOne
	if partTwo {
		fn = day_01.PartTwo
	}

	total, err := fn(slog.Default(), os.Stdin)
	if err != nil {
		fmt.Printf("failed to run: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Calibration value sum: %d\n", total)
}
