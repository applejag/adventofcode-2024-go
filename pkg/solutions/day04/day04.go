package day04

import (
	"bytes"
	"io"
	"slices"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
)

type Day struct{}

var _ solutions.Day = Day{}

func (Day) readGrid(file io.Reader) ([][]byte, error) {
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(bytes.TrimSpace(b), []byte("\n"))
	return lines, nil
}

func (Day) columns(lines [][]byte) [][]byte {
	columns := make([][]byte, len(lines[0]))
	for i := range columns {
		columns[i] = make([]byte, len(lines))
	}
	for y, line := range lines {
		for x, char := range line {
			columns[x][y] = char
		}
	}
	return columns
}

func (Day) diagonals(lines [][]byte) [][]byte {
	/*

		  0 1 2 3 4 x
		0 . . . . .
		1 . . . . .
		2 . . . . .
		3 . . . . .
		4 . . . . .
		y

		x,y
		0,0

		0,1
		1,0

		0,2
		1,1
		0,2

		0,3
		1,2
		2,1
		4,0

		0,4
		1,3
		2,2
		3,1
		4,0

		0,5
		1,4
		2,3
		3,2
		4,1
		5,0
	*/

	width := len(lines[0])
	height := len(lines)
	diagonals := make([][]byte, 0, width+height)

	for limit := range width + height - 1 {
		var diag []byte
		x := 0
		y := limit
		for x < width && y >= 0 {
			if y >= height {
				x++
				y--
				continue
			}
			if x >= width {
				x++
				y--
				continue
			}
			diag = append(diag, lines[y][x])

			x++
			y--
		}

		diagonals = append(diagonals, diag)
	}

	return diagonals
}

func (d Day) invertedDiagonals(lines [][]byte) [][]byte {
	reversedLines := make([][]byte, len(lines))
	copy(reversedLines, lines)
	slices.Reverse(reversedLines)
	return d.diagonals(reversedLines)
}

func (d Day) Part1(file io.Reader) (any, error) {
	lines, err := d.readGrid(file)
	if err != nil {
		return nil, err
	}

	xmas := []byte("XMAS")
	samx := []byte("SAMX")

	var count int
	for _, line := range lines {
		// left to right
		count += bytes.Count(line, xmas)

		// right to left
		count += bytes.Count(line, samx)
	}

	columns := d.columns(lines)
	for _, col := range columns {
		// top to bottom
		count += bytes.Count(col, xmas)

		// bottom to top
		count += bytes.Count(col, samx)
	}

	diagonals := d.diagonals(lines)
	for _, line := range diagonals {
		// down-left to up-right
		count += bytes.Count(line, xmas)

		// up-right to down-left
		count += bytes.Count(line, samx)
	}

	otherDiagonals := d.invertedDiagonals(lines)
	for _, col := range otherDiagonals {
		// up-left to down-right
		count += bytes.Count(col, xmas)

		// down-right to up-left
		count += bytes.Count(col, samx)
	}

	return count, nil
}

func (d Day) Part2(file io.Reader) (any, error) {
	lines, err := d.readGrid(file)
	if err != nil {
		return nil, err
	}
	var count int

	mas := [3]byte{'M', 'A', 'S'}
	sam := [3]byte{'S', 'A', 'M'}

	for y := 1; y < len(lines)-1; y++ {
		line := lines[y]
		for x := 1; x < len(line)-1; x++ {
			if lines[y][x] != 'A' {
				continue
			}

			var diag1 [3]byte
			diag1[0] = lines[y-1][x-1]
			diag1[1] = lines[y][x]
			diag1[2] = lines[y+1][x+1]

			var diag2 [3]byte
			diag2[0] = lines[y-1][x+1]
			diag2[1] = lines[y][x]
			diag2[2] = lines[y+1][x-1]

			if (diag1 == sam || diag1 == mas) &&
				(diag2 == sam || diag2 == mas) {
				count++
			}
		}
	}

	return count, nil
}
