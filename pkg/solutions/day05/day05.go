package day05

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
)

type Day struct{}

var _ solutions.Day = Day{}

func (Day) Part1(file io.Reader) (any, error) {
	rules, updates, err := parseInput(file)
	if err != nil {
		return nil, err
	}

	var sum int
	for _, update := range updates {
		if ValidateRules(rules, update) {
			sum += update.Middle()
		}
	}

	return sum, nil
}

func (Day) Part2(file io.Reader) (any, error) {
	rules, updates, err := parseInput(file)
	if err != nil {
		return nil, err
	}

	var sum int
	for _, update := range updates {
		if !ValidateRules(rules, update) {
			fixed := FixUpdate(rules, update)
			sum += fixed.Middle()
		}
	}

	return sum, nil
}

type OrderingRule struct {
	Before int
	After  int
}

func ValidateRules(rules []OrderingRule, update Update) bool {
	return FindInvalidRule(rules, update) == nil
}

func FindInvalidRule(rules []OrderingRule, update Update) *OrderingRule {
	for i, rule := range rules {
		if !rule.Validate(update) {
			return &rules[i]
		}
	}
	return nil
}

func FixUpdate(rules []OrderingRule, update Update) Update {
	update = slices.Clone(update)
	for range 100 {
		rule := FindInvalidRule(rules, update)
		if rule == nil {
			return update
		}
		idxBefore := slices.Index(update, rule.Before)
		idxAfter := slices.Index(update, rule.After)
		update[idxBefore], update[idxAfter] = update[idxAfter], update[idxBefore]
	}
	panic(fmt.Errorf("failed to fix update after 100 attempts"))
}

func (rule OrderingRule) Validate(update Update) bool {
	for i, n := range update {
		if rule.Before == n {
			prevNumbers := update[:i]
			if slices.Contains(prevNumbers, rule.After) {
				return false
			}
		}
		if rule.After == n {
			nextNumbers := update[i+1:]
			if slices.Contains(nextNumbers, rule.Before) {
				return false
			}
		}
	}
	return true
}

type Update []int

func (u Update) Middle() int {
	return u[len(u)/2]
}

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
