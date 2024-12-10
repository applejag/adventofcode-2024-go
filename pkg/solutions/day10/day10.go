package day10

import (
	"bytes"
	"io"
	"iter"
	"slices"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
	"gopkg.in/typ.v4/arrays"
)

type Day struct{}

var _ solutions.Day = Day{}

func (Day) Part1(file io.Reader) (any, error) {
	grid, err := ParseGrid(file)
	if err != nil {
		return nil, err
	}
	var sum int
	for vec := range IterTrailheads(grid) {
		sum += TraverseTrailPart1(grid, vec)
	}
	return sum, nil
}

func (Day) Part2(file io.Reader) (any, error) {
	grid, err := ParseGrid(file)
	if err != nil {
		return nil, err
	}
	var sum int
	for vec := range IterTrailheads(grid) {
		sum += TraverseTrailPart2(grid, vec)
	}
	return sum, nil
}

func IterTrailheads(grid Grid) iter.Seq[Vec2] {
	return func(yield func(Vec2) bool) {
		for y := range grid.Height() {
			for x, char := range grid.Row(y) {
				if char != '0' {
					continue
				}

				if !yield(Vec2{x, y}) {
					return
				}
			}
		}
	}
}

type Grid = arrays.Array2D[byte]

func TraverseTrailPart1(grid Grid, pos Vec2) int {
	next := []Vec2{pos}
	var visited []Vec2
	var peaks int

	for len(next) > 0 {
		if len(visited) > 10000 {
			panic("possible infinite loop")
		}
		currentPos := next[len(next)-1]
		next = next[:len(next)-1]
		visited = append(visited, currentPos)
		currentValue := grid.Get(currentPos.X, currentPos.Y)

		if currentValue == '9' {
			peaks++
			continue
		}

		if p := currentPos.Add(Vec2{1, 0}); !slices.Contains(visited, p) && GridPosIsNext(grid, currentValue, p) {
			next = append(next, p)
		}
		if p := currentPos.Add(Vec2{0, 1}); !slices.Contains(visited, p) && GridPosIsNext(grid, currentValue, p) {
			next = append(next, p)
		}
		if p := currentPos.Add(Vec2{-1, 0}); !slices.Contains(visited, p) && GridPosIsNext(grid, currentValue, p) {
			next = append(next, p)
		}
		if p := currentPos.Add(Vec2{0, -1}); !slices.Contains(visited, p) && GridPosIsNext(grid, currentValue, p) {
			next = append(next, p)
		}
	}

	return peaks
}

func TraverseTrailPart2(grid Grid, pos Vec2) int {
	next := []Vec2{pos}
	var visited []Vec2
	var peaks int

	var queued []Vec2

	for len(next) > 0 {
		if len(visited) > 10000 {
			panic("possible infinite loop")
		}
		currentPos := next[len(next)-1]
		next = next[:len(next)-1]
		visited = append(visited, currentPos)
		currentValue := grid.Get(currentPos.X, currentPos.Y)

		if currentValue == '9' {
			peaks++
			continue
		}

		queued = queued[0:0]
		if p := currentPos.Add(Vec2{1, 0}); !slices.Contains(visited, p) && GridPosIsNext(grid, currentValue, p) {
			queued = append(queued, p)
		}
		if p := currentPos.Add(Vec2{0, 1}); !slices.Contains(visited, p) && GridPosIsNext(grid, currentValue, p) {
			queued = append(queued, p)
		}
		if p := currentPos.Add(Vec2{-1, 0}); !slices.Contains(visited, p) && GridPosIsNext(grid, currentValue, p) {
			queued = append(queued, p)
		}
		if p := currentPos.Add(Vec2{0, -1}); !slices.Contains(visited, p) && GridPosIsNext(grid, currentValue, p) {
			queued = append(queued, p)
		}

		switch len(queued) {
		case 0:
			// do nothing
		case 1:
			next = append(next, queued[0])
		default:
			// branch
			for _, p := range queued {
				peaks += TraverseTrailPart2(grid, p)
			}
			return peaks
		}
	}

	return peaks
}

func GridPosIsNext(grid Grid, currentValue byte, pos Vec2) bool {
	if !GridContains(grid, pos) {
		return false
	}
	return grid.Get(pos.X, pos.Y) == currentValue+1
}

func GridContains(grid Grid, pos Vec2) bool {
	return pos.X >= 0 && pos.X < grid.Width() && pos.Y >= 0 && pos.Y < grid.Height()
}

func ParseGrid(file io.Reader) (Grid, error) {
	b, err := io.ReadAll(file)
	if err != nil {
		return Grid{}, err
	}
	lines := bytes.Split(bytes.TrimSpace(b), []byte{'\n'})
	width := len(lines[0])
	height := len(lines)
	grid := arrays.New2D[byte](width, height)

	for y, line := range lines {
		copy(grid.Row(y), line)
	}
	return grid, nil
}
