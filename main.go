package main

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
	"github.com/charmbracelet/log"
	"github.com/spf13/pflag"
)

var flags = struct {
	file string
	help bool
}{}

func init() {
	handler := log.NewWithOptions(os.Stderr, log.Options{Level: log.InfoLevel})
	slog.SetDefault(slog.New(handler))

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: go run . <day> [flags]

Flags:
`)
		pflag.PrintDefaults()
	}

	pflag.BoolVarP(&flags.help, "help", "h", false, "Show this help text")
	pflag.StringVarP(&flags.file, "file", "f", "", `Input file (defaults to "inputs/day%02d.txt" using the "<day>" argument)`)
}

func main() {
	pflag.Parse()
	if flags.help {
		pflag.Usage()
		os.Exit(0)
	}

	if pflag.NArg() != 1 {
		slog.Error("Wrong number of arguments")
		pflag.Usage()
		os.Exit(1)
	}

	dayNum, err := strconv.Atoi(pflag.Arg(0))
	if err != nil {
		slog.Error("Failed to parse argument", "error", err)
		os.Exit(1)
	}

	var day solutions.Day
	switch dayNum {
	case 3:
		day = solutions.Day03{}
	default:
		slog.Error("Invalid day", "day", dayNum)
		os.Exit(1)
	}

	if flags.file == "" {
		flags.file = fmt.Sprintf("inputs/day%02d.txt", dayNum)
	}

	file, err := os.Open(flags.file)
	if err != nil {
		slog.Error("Failed to open file", "error", err)
		os.Exit(1)
	}
	defer file.Close()

	slog.Debug("Executing AoC solution", "day", dayNum, "file", flags.file)

	runPart(day, dayNum, 1, file)

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		slog.Error("Failed to seek file", "error", err)
		os.Exit(1)
	}

	runPart(day, dayNum, 2, file)
}

func runPart(day solutions.Day, dayNum, part int, file io.Reader) {
	var solution any
	var err error
	switch part {
	case 1:
		solution, err = day.Part1(file)
	default:
		solution, err = day.Part2(file)
	}

	if err == errors.ErrUnsupported {
		slog.Warn("Part has not been implemented", "day", dayNum, "part", part)
		return
	}

	if err != nil {
		slog.Error("Failed to calculate solution", "day", dayNum, "part", part, "error", err)
		os.Exit(1)
	}

	slog.Info("", "day", dayNum, "part", part, "solution", fmt.Sprintf("%v\n", solution))
}
