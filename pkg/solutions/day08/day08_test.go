package day08

import (
	"testing"

	"github.com/applejag/adventofcode-2024-go/pkg/testutil"
)

func TestPart1(t *testing.T) {
	testutil.AssertPart1(t, Day{}, 14, `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)
}

func TestPart2(t *testing.T) {
	testutil.AssertPart2(t, Day{}, 34, `
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`)
}
