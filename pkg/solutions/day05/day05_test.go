package day05

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseInput(t *testing.T) {
	input := `47|53
97|13
97|61
97|47

75,47,61,53,29
97,61,53,29,13`

	rules, updates, err := parseInput(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	wantRules := []OrderingRule{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
	}

	wantUpdates := []Update{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
	}

	if diff := cmp.Diff(wantRules, rules); diff != "" {
		t.Errorf("rules: (-want, +got)\n%s", diff)
	}

	if diff := cmp.Diff(wantUpdates, updates); diff != "" {
		t.Errorf("updates: (-want, +got)\n%s", diff)
	}
}
