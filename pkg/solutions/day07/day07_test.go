package day07

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/applejag/adventofcode-2024-go/pkg/testutil"
)

func TestPart1(t *testing.T) {
	testutil.AssertPart1(t, Day{}, int64(3749), `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`)
}

func TestPart2(t *testing.T) {
	testutil.AssertPart2(t, Day{}, int64(11387), `
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
`)
}

func TestPermutationsPart2(t *testing.T) {
	eq := Equation{
		Result:   7290,
		Operands: []int64{6, 8, 6, 15},
	}
	want := []Operator{OpMul, OpConcat, OpMul}

	var permutations [][]Operator
	for perm := range eq.PermutationsIterPart2() {
		permutations = append(permutations, slices.Clone(perm))
	}

	if !slices.ContainsFunc(permutations, func(e []Operator) bool {
		return slices.Equal(e, want)
	}) {
		var sb strings.Builder
		for _, perm := range permutations {
			fmt.Fprintln(&sb, perm)
		}
		t.Errorf("missing %v, all:\n%s", want, sb.String())
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		a, b int64
		want int64
	}{
		{a: 15, b: 6, want: 156},
		{a: 17, b: 8, want: 178},
		{a: 1, b: 1, want: 11},
		{a: 100, b: 100, want: 100100},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("%d..%d", tc.a, tc.b), func(t *testing.T) {
			got := concat(tc.a, tc.b)
			if got != tc.want {
				t.Errorf("%d || %d = ?\nwant: %d\ngot:  %d", tc.a, tc.b, tc.want, got)
			}
		})
	}
}
