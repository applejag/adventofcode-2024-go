package solutions

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"regexp"
	"strconv"
)

var day03Regex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

type Day03 struct{}

func (Day03) Part1(file io.Reader) (any, error) {
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		matches := day03Regex.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matches) == 0 {
			return nil, fmt.Errorf("no matches on line: %q", scanner.Text())
		}
		for _, match := range matches {
			lhs, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, fmt.Errorf("parse LHS: %s: %w", match[0], err)
			}
			rhs, err := strconv.Atoi(match[2])
			if err != nil {
				return nil, fmt.Errorf("parse RHS: %s: %w", match[0], err)
			}
			sum += lhs * rhs
		}
	}
	return sum, nil
}

var day03DoDontRegex = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)

func (Day03) Part2(file io.Reader) (any, error) {
	scanner := bufio.NewScanner(file)
	var sum int
	enabled := true
	for scanner.Scan() {
		matches := day03DoDontRegex.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matches) == 0 {
			return nil, fmt.Errorf("no matches on line: %q", scanner.Text())
		}
		for _, match := range matches {
			switch match[0] {
			case "do()":
				slog.Debug("Match", "full", match[0], "enabled", fmt.Sprintf("%t->true", enabled))
				enabled = true
				continue
			case "don't()":
				slog.Debug("Match", "full", match[0], "enabled", fmt.Sprintf("%t->false", enabled))
				enabled = false
				continue
			}
			slog.Debug("Match", "full", match[0], "enabled", enabled, "lhs", match[1], "rhs", match[2])
			if !enabled {
				continue
			}
			lhs, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, fmt.Errorf("parse LHS: %s: %w", match[0], err)
			}
			rhs, err := strconv.Atoi(match[2])
			if err != nil {
				return nil, fmt.Errorf("parse RHS: %s: %w", match[0], err)
			}
			sum += lhs * rhs
		}
	}
	return sum, nil
}
