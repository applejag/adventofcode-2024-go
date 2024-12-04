package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
	"github.com/charmbracelet/log"
	"github.com/spf13/pflag"
)

var flags = struct {
	file string
}{}

func init() {
	handler := log.NewWithOptions(os.Stderr, log.Options{Level: log.DebugLevel})
	slog.SetDefault(slog.New(handler))
}

func main() {
	pflag.Parse()

	if pflag.NArg() != 1 {
		slog.Error("Wrong number of arguments")
		pflag.Usage()
		os.Exit(1)
	}

	day, err := strconv.Atoi(pflag.Arg(0))
	if err != nil {
		slog.Error("Failed to parse argument", "error", err)
		os.Exit(1)
	}

	if flags.file == "" {
		flags.file = fmt.Sprintf("inputs/day%02d.txt", day)
	}

	file, err := os.Open(flags.file)
	if err != nil {
		slog.Error("Failed to open file", "error", err)
		os.Exit(1)
	}
	defer file.Close()

	slog.Debug("Executing AoC solution", "day", day, "file", flags.file)

	var solution any
	var solutionErr error

	switch day {
	case 3:
		solution, solutionErr = solutions.Day03Part1(file)
	default:
		slog.Error("Invalid day", "day", solutionErr)
		os.Exit(1)
	}

	if solutionErr != nil {
		slog.Error("Failed to calculate solution", "day", day, "error", solutionErr)
		os.Exit(1)
	}

	slog.Info("", "day", day, "part", 1, "solution", fmt.Sprintf("%v\n", solution))
}
