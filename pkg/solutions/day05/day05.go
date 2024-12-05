package day05

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
)

type Day struct{}

var _ solutions.Day = Day{}

func (Day) Part1(file io.Reader) (any, error) {

	return nil, solutions.ErrNotImplemented
}

func (Day) Part2(file io.Reader) (any, error) {
	return nil, solutions.ErrNotImplemented
}

type OrderingRule struct {
	Before int
	After  int
}

type Update []int

func parseInput(file io.Reader) ([]OrderingRule, []Update, error) {
	scanner := bufio.NewScanner(file)

	var rules []OrderingRule
	for scanner.Scan() && len(scanner.Bytes()) != 0 {
		var rule OrderingRule
		if _, err := fmt.Sscanf(scanner.Text(), "%d|%d", &rule.Before, &rule.After); err != nil {
			return nil, nil, fmt.Errorf("parse rule: %w", err)
		}
		rules = append(rules, rule)
	}

	var updates []Update
	for scanner.Scan() {
		var update Update
		for _, s := range strings.Split(scanner.Text(), ",") {
			num, err := strconv.Atoi(s)
			if err != nil {
				return nil, nil, fmt.Errorf("parse update num: %w", err)
			}
			update = append(update, num)
		}
		updates = append(updates, update)
	}

	return rules, updates, nil
}
