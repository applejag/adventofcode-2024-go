package solutions

import (
	"bytes"
	"io"
	"slices"
)

type Day04 struct{}

func (Day04) readGrid(file io.Reader) ([][]byte, error) {
	b, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(bytes.TrimSpace(b), []byte("\n"))
	return lines, nil
}

func (Day04) columns(lines [][]byte) [][]byte {
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

func (Day04) diagonals(lines [][]byte) [][]byte {
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

func (d Day04) invertedDiagonals(lines [][]byte) [][]byte {
	reversedLines := make([][]byte, len(lines))
	copy(reversedLines, lines)
	slices.Reverse(reversedLines)
	return d.diagonals(reversedLines)
}

func (d Day04) Part1(file io.Reader) (any, error) {
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

func (Day04) Part2(file io.Reader) (any, error) {
	return nil, ErrNotImplemented
}
