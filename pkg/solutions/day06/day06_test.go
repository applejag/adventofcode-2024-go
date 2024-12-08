package day06

import (
	"testing"

	"github.com/applejag/adventofcode-2024-go/pkg/testutil"
)

func TestPart1(t *testing.T) {
	testutil.AssertPart1(t, Day{}, 41,  `
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)
}
