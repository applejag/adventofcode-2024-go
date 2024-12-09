package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
	"github.com/applejag/adventofcode-2024-go/pkg/solutions/day01"
	"github.com/applejag/adventofcode-2024-go/pkg/solutions/day02"
	"github.com/applejag/adventofcode-2024-go/pkg/solutions/day03"
	"github.com/applejag/adventofcode-2024-go/pkg/solutions/day04"
	"github.com/applejag/adventofcode-2024-go/pkg/solutions/day05"
	"github.com/applejag/adventofcode-2024-go/pkg/solutions/day06"
	"github.com/applejag/adventofcode-2024-go/pkg/solutions/day07"
	"github.com/applejag/adventofcode-2024-go/pkg/solutions/day08"
	"github.com/applejag/adventofcode-2024-go/pkg/solutions/day09"
	"github.com/charmbracelet/log"
	"github.com/spf13/pflag"
)

var days = map[int]solutions.Day{
	1:  day01.Day{},
	2:  day02.Day{},
	3:  day03.Day{},
	4:  day04.Day{},
	5:  day05.Day{},
	6:  day06.Day{},
	7:  day07.Day{},
	8:  day08.Day{},
	9:  day09.Day{},
	10: solutions.UnimplementedDay{},
	11: solutions.UnimplementedDay{},
	12: solutions.UnimplementedDay{},
	13: solutions.UnimplementedDay{},
	14: solutions.UnimplementedDay{},
	15: solutions.UnimplementedDay{},
	16: solutions.UnimplementedDay{},
	17: solutions.UnimplementedDay{},
	18: solutions.UnimplementedDay{},
	19: solutions.UnimplementedDay{},
	20: solutions.UnimplementedDay{},
	21: solutions.UnimplementedDay{},
	22: solutions.UnimplementedDay{},
	23: solutions.UnimplementedDay{},
	24: solutions.UnimplementedDay{},
	25: solutions.UnimplementedDay{},
}

var flags = struct {
	part    int
	file    string
	help    bool
	verbose bool
}{}

var logger *log.Logger

func init() {
	logger = log.NewWithOptions(os.Stderr, log.Options{Level: log.InfoLevel})
	slog.SetDefault(slog.New(logger))

	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: go run . <day> [flags]

Flags:
`)
		pflag.PrintDefaults()
	}

	pflag.IntVarP(&flags.part, "part", "p", 0, "Part to execute, where 0 means run both part 1 and part 2")
	pflag.StringVarP(&flags.file, "file", "f", "", `Input file (defaults to "inputs/day%02d.txt" using the "<day>" argument)`)
	pflag.BoolVarP(&flags.help, "help", "h", false, "Show this help text")
	pflag.BoolVarP(&flags.verbose, "verbose", "v", false, "Show debug logs")
}

func main() {
	pflag.Parse()
	if flags.help {
		pflag.Usage()
		os.Exit(0)
	}

	if flags.verbose {
		logger.SetLevel(log.DebugLevel)
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
	day, ok := days[dayNum]
	if !ok {
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

	slog.Debug("Opened file", "file", flags.file)

	runPart(day, dayNum, 1, file)

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		slog.Error("Failed to seek file", "error", err)
		os.Exit(1)
	}

	runPart(day, dayNum, 2, file)
}

func runPart(day solutions.Day, dayNum, part int, file io.Reader) {
	if flags.part != 0 && flags.part != part {
		slog.Debug("Skipping part because of the --part flag")
		return
	}
	var solution any
	var err error
	switch part {
	case 1:
		solution, err = day.Part1(file)
	default:
		solution, err = day.Part2(file)
	}

	if err == solutions.ErrNotImplemented {
		slog.Warn("Part has not been implemented", "day", dayNum, "part", part)
		return
	}

	if err != nil {
		slog.Error("Failed to calculate solution", "day", dayNum, "part", part, "error", err)
		os.Exit(1)
	}

	slog.Info("", "day", dayNum, "part", part, "solution", fmt.Sprintf("%v\n", solution))
}
