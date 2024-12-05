package day04

import (
	"bytes"
	"strings"
	"testing"

	"github.com/applejag/adventofcode-2024-go/pkg/testutil"
)

func TestDay04Part1(t *testing.T) {
	testutil.SetLogger(t)
	input := `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	testutil.AssertPart1(t, Day{}, 18, input)
}

func TestDay04Columns(t *testing.T) {
	testutil.SetLogger(t)
	diagonals := Day{}.columns(bytes.Split(bytes.TrimSpace([]byte(`
abcd
efgh
ijkl
`)), []byte{'\n'}))
	want := strings.TrimSpace(`
aei
bfj
cgk
dhl
`)
	got := string(bytes.Join(diagonals, []byte{'\n'}))
	if want != got {
		t.Errorf("want:\n  | %s\ngot:\n  | %s",
			strings.ReplaceAll(want, "\n", "\n  | "),
			strings.ReplaceAll(got, "\n", "\n  | "))
	}
}

func TestDay04Diagonals(t *testing.T) {
	testutil.SetLogger(t)
	diagonals := Day{}.diagonals(bytes.Split(bytes.TrimSpace([]byte(`
abcd
efgh
ijkl
`)), []byte{'\n'}))
	want := strings.TrimSpace(`
a
eb
ifc
jgd
kh
l
`)
	got := string(bytes.Join(diagonals, []byte{'\n'}))
	if want != got {
		t.Errorf("want:\n  | %s\ngot:\n  | %s",
			strings.ReplaceAll(want, "\n", "\n  | "),
			strings.ReplaceAll(got, "\n", "\n  | "))
	}
}

func TestDay04DiagonalsInverted(t *testing.T) {
	testutil.SetLogger(t)
	lines := bytes.Split(bytes.TrimSpace([]byte(`
abcd
efgh
ijkl
`)), []byte{'\n'})
	diagonals := Day{}.invertedDiagonals(lines)
	want := strings.TrimSpace(`
i
ej
afk
bgl
ch
d
`)
	got := string(bytes.Join(diagonals, []byte{'\n'}))
	if want != got {
		t.Errorf("want:\n  | %s\ngot:\n  | %s",
			strings.ReplaceAll(want, "\n", "\n  | "),
			strings.ReplaceAll(got, "\n", "\n  | "))
	}
}

func TestDay04Part2(t *testing.T) {
	testutil.SetLogger(t)
	input := `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	testutil.AssertPart2(t, Day{}, 9, input)
}
