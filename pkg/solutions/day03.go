package solutions

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var day03Regex = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func Day03Part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	var sum int
	for scanner.Scan() {
		matches := day03Regex.FindAllStringSubmatch(scanner.Text(), -1)
		if len(matches) == 0 {
			return 0, fmt.Errorf("no matches on line: %q", scanner.Text())
		}
		for _, match := range matches {
			lhs, err := strconv.Atoi(match[1])
			if err != nil {
				return 0, fmt.Errorf("parse LHS: %s: %w", match[0], err)
			}
			rhs, err := strconv.Atoi(match[2])
			if err != nil {
				return 0, fmt.Errorf("parse RHS: %s: %w", match[0], err)
			}
			sum += lhs * rhs
		}
	}
	return sum, nil
}
