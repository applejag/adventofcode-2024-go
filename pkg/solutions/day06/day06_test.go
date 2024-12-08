package day06

import (
	"strings"
	"testing"

	"github.com/applejag/adventofcode-2024-go/pkg/testutil"
)

func TestPart1(t *testing.T) {
	testutil.AssertPart1(t, Day{}, 41, `
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

func TestPart2(t *testing.T) {
	testutil.AssertPart2(t, Day{}, 6, `
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

func TestWalkLoop(t *testing.T) {
	input := `
####
#..#
#^.#
####`
	m, err := ParseMap(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}
	moveErr := m.GuardMove()
	if moveErr == nil {
		t.Errorf("want error %q, got error nil", ErrMoveLoop)
	} else if moveErr != ErrMoveLoop {
		t.Errorf("want error %q, got error %q", ErrMoveLoop, moveErr)
	}
}
