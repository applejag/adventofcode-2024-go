package day03

import (
	"testing"

	"github.com/applejag/adventofcode-2024-go/pkg/testutil"
)

func TestDay03Part1(t *testing.T) {
	testutil.SetLogger(t)
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	testutil.AssertPart1(t, Day{}, 161, input)
}

func TestDay03Part2(t *testing.T) {
	testutil.SetLogger(t)
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	testutil.AssertPart2(t, Day{}, 48, input)
}
