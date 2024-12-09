package day09

import (
	"slices"
	"strings"
	"testing"

	"github.com/applejag/adventofcode-2024-go/pkg/testutil"
)

func TestPart1(t *testing.T) {
	testutil.AssertPart1(t, Day{}, 1928, `2333133121414131402`)
}

func TestParse(t *testing.T) {
	got, err := ParseDigits(strings.NewReader("23331"))
	if err != nil {
		t.Fatal(err)
	}
	want := []byte{2, 3, 3, 3, 1}
	if !slices.Equal(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}
