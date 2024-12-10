package day10

import (
	"testing"

	"github.com/applejag/adventofcode-2024-go/pkg/testutil"
)

func TestPart1(t *testing.T) {
	testutil.AssertPart1(t, Day{}, 36, `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`)
}

func TestPart2(t *testing.T) {
	testutil.AssertPart2(t, Day{}, 81, `
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732
`)
}
