package day08

import (
	"bytes"
	"io"
	"log/slog"
	"strings"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
	"gopkg.in/typ.v4/arrays"
)

type Day struct{}

var _ solutions.Day = Day{}

func (Day) Part1(file io.Reader) (any, error) {
	m, err := Parse(file)
	if err != nil {
		return nil, err
	}

	var count int
	for _, positions := range m.AntennaPositions {
		for i, pos := range positions {
			for _, other := range positions[i+1:] {

				delta := other.Sub(pos)
				anti1 := pos.Sub(delta)
				anti2 := other.Add(delta)

				if m.ContainsPoint(anti1) {
					if m.Grid.Get(anti1.X, anti1.Y) != NodeAnti {
						m.Grid.Set(anti1.X, anti1.Y, NodeAnti)
						count++
					}
				}

				if m.ContainsPoint(anti2) {
					if m.Grid.Get(anti2.X, anti2.Y) != NodeAnti {
						m.Grid.Set(anti2.X, anti2.Y, NodeAnti)
						count++
					}
				}
			}
		}
	}

	slog.Debug("", "map", m.String())

	return count, nil
}

func (Day) Part2(file io.Reader) (any, error) {
	m, err := Parse(file)
	if err != nil {
		return nil, err
	}

	var count int
	for _, positions := range m.AntennaPositions {
		for i, pos := range positions {
			if m.Grid.Get(pos.X, pos.Y) != NodeAnti {
				m.Grid.Set(pos.X, pos.Y, NodeAnti)
				count++
			}

			for _, other := range positions[i+1:] {

				delta := other.Sub(pos)
				if delta == (Vec2{0, 0}) {
					panic("oh no zero delta")
				}

				for anti1 := pos.Sub(delta); m.ContainsPoint(anti1); anti1 = anti1.Sub(delta) {
					if m.Grid.Get(anti1.X, anti1.Y) != NodeAnti {
						m.Grid.Set(anti1.X, anti1.Y, NodeAnti)
						count++
					}
				}

				for anti2 := other.Add(delta); m.ContainsPoint(anti2); anti2 = anti2.Add(delta) {
					if m.Grid.Get(anti2.X, anti2.Y) != NodeAnti {
						m.Grid.Set(anti2.X, anti2.Y, NodeAnti)
						count++
					}
				}
			}
		}
	}

	slog.Debug("", "map", m.String())

	return count, nil
}

type Node byte

const (
	NodeEmpty Node = '.'
	NodeAnti  Node = '#'
)

type Map struct {
	Grid             arrays.Array2D[Node]
	AntennaPositions map[Node][]Vec2
}

func (m Map) String() string {
	var sb strings.Builder
	sb.Grow((m.Grid.Width() + 1) * m.Grid.Height()) // +1 to account for newlines
	for y := range m.Grid.Height() {
		if y > 0 {
			sb.WriteByte('\n')
		}
		for _, node := range m.Grid.Row(y) {
			sb.WriteByte(byte(node))
		}
	}
	return sb.String()
}

func (m Map) ContainsPoint(v Vec2) bool {
	if v.X < 0 || v.X >= m.Grid.Width() ||
		v.Y < 0 || v.Y >= m.Grid.Height() {
		return false
	}
	return true
}

func Parse(file io.Reader) (Map, error) {
	b, err := io.ReadAll(file)
	if err != nil {
		return Map{}, err
	}
	lines := bytes.Split(bytes.TrimSpace(b), []byte{'\n'})
	width := len(lines[0])
	height := len(lines)
	grid := arrays.New2DFilled(width, height, NodeEmpty)
	antennasPerType := map[Node][]Vec2{}

	for y, line := range lines {
		for x, char := range line {
			if char == '.' || char == '#' {
				continue
			}
			antenna := Node(char)
			grid.Set(x, y, antenna)
			antennasPerType[antenna] = append(antennasPerType[antenna], Vec2{x, y})
		}
	}
	return Map{
		Grid:             grid,
		AntennaPositions: antennasPerType,
	}, nil
}
