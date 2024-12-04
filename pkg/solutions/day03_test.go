package solutions

import (
	"strings"
	"testing"
)

func TestDay03(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	got, err := Day03Part1(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	want := 161
	if got != want {
		t.Fatalf("got %d, want %d", got, want)
	}
}
