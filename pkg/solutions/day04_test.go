package solutions

import (
	"bytes"
	"strings"
	"testing"
)

func TestDay04Part1(t *testing.T) {
	setLogger(t)
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
	AssertPart1(t, Day04{}, 18, input)
}

func TestDay04Columns(t *testing.T) {
	setLogger(t)
	diagonals := Day04{}.columns(bytes.Split(bytes.TrimSpace([]byte(`
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
	setLogger(t)
	diagonals := Day04{}.diagonals(bytes.Split(bytes.TrimSpace([]byte(`
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
	setLogger(t)
	lines := bytes.Split(bytes.TrimSpace([]byte(`
abcd
efgh
ijkl
`)), []byte{'\n'})
	diagonals := Day04{}.invertedDiagonals(lines)
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
	setLogger(t)
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
	AssertPart2(t, Day04{}, 9, input)
}
