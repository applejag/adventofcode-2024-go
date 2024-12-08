package day05

import (
	"slices"
	"strings"
	"testing"

	"github.com/applejag/adventofcode-2024-go/pkg/testutil"
	"github.com/google/go-cmp/cmp"
)

func TestPart1(t *testing.T) {
	testutil.AssertPart1(t, Day{}, 143, `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`)
}

func TestPart2(t *testing.T) {
	testutil.AssertPart2(t, Day{}, 123, `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`)
}

func TestValidate(t *testing.T) {
	rules := []OrderingRule{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	}

	tests := []struct {
		name   string
		update Update
		want   bool
	}{
		{
			name:   "first ok",
			update: Update{75, 47, 61, 53, 29},
			want:   true,
		},
		{
			name:   "second ok",
			update: Update{97, 61, 53, 29, 13},
			want:   true,
		},
		{
			name:   "third ok",
			update: Update{75, 29, 13},
			want:   true,
		},
		{
			name:   "forth not ok",
			update: Update{75, 97, 47, 61, 53},
			want:   false,
		},
		{
			name:   "fifth not ok",
			update: Update{61, 13, 29},
			want:   false,
		},
		{
			name:   "sixth not ok",
			update: Update{97, 13, 75, 29, 47},
			want:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateRules(rules, tc.update)
			t.Logf("update: %v", tc.update)
			if tc.want != got {
				t.Errorf("want %t, got %t", tc.want, got)
			}
		})
	}
}

func TestMiddle(t *testing.T) {
	update := Update{75, 47, 61, 53, 29}
	want := 61
	got := update.Middle()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

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

func TestFixUpdate(t *testing.T) {
	rules := []OrderingRule{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	}

	tests := []struct {
		name   string
		update Update
		want   Update
	}{
		{
			name:   "forth",
			update: Update{75, 97, 47, 61, 53},
			want:   Update{97, 75, 47, 61, 53},
		},
		{
			name:   "fifth",
			update: Update{61, 13, 29},
			want:   Update{61, 29, 13},
		},
		{
			name:   "sixth",
			update: Update{97, 13, 75, 29, 47},
			want:   Update{97, 75, 47, 29, 13},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("update: %v", tc.update)
			got := FixUpdate(rules, tc.update)
			if !slices.Equal(tc.want, got) {
				t.Errorf("wrong output\nwant: %v\ngot:  %v", tc.want, got)
			}
		})
	}
}
