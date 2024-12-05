package testutil

import (
	"log/slog"
	"strings"
	"testing"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
	"github.com/charmbracelet/log"
)

func AssertPart1(t *testing.T, day solutions.Day, want any, input string) {
	t.Helper()
	got, err := day.Part1(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func AssertPart2(t *testing.T, day solutions.Day, want any, input string) {
	t.Helper()
	got, err := day.Part2(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}

func SetLogger(t *testing.T) {
	origLogger := slog.Default()
	t.Cleanup(func() {
		slog.SetDefault(origLogger)
	})
	logger := log.NewWithOptions(testLogWriter{t}, log.Options{Level: log.DebugLevel})
	slog.SetDefault(slog.New(logger))
}

type testLogWriter struct {
	t *testing.T
}

func (w testLogWriter) Write(b []byte) (int, error) {
	w.t.Log(strings.TrimSuffix(string(b), "\n"))
	return len(b), nil
}
