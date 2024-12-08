package day06

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log/slog"

	"github.com/applejag/adventofcode-2024-go/pkg/solutions"
	"gopkg.in/typ.v4/arrays"
)

type Day struct{}

var _ solutions.Day = Day{}

func (Day) Part1(file io.Reader) (any, error) {
	m, err := ParseMap(file)
	if err != nil {
		return nil, err
	}
	if err := m.GuardMove(); !errors.Is(err, ErrMoveOutOfBounds) {
		return nil, err
	}
	var sum int
	for y := range m.Grid.Height() {
		row := m.Grid.Row(y)
		sum += bytes.Count(row, []byte{'X'})
	}

	return sum, nil
}

func (Day) Part2(file io.Reader) (any, error) {
	m, err := ParseMap(file)
	if err != nil {
		return nil, err
	}

	var loopCount int
	for x := range m.Grid.Width() {
		for y := range m.Grid.Height() {
			if m.Grid.Get(x, y) != '.' {
				continue
			}
			clone := m.Clone()
			clone.Grid.Set(x, y, '#')
			if err := clone.GuardMove(); errors.Is(err, ErrMoveLoop) {
				slog.Debug("Loop", "x", x, "y", y)
				loopCount++
			}
		}
	}

	return loopCount, nil
}

type Vec2 struct {
	X int
	Y int
}

type Facing byte

const (
	FacingUp Facing = 1 << iota
	FacingDown
	FacingLeft
	FacingRight
)

func (f Facing) Delta() Vec2 {
	switch f {
	case FacingUp:
		return Vec2{0, -1}
	case FacingDown:
		return Vec2{0, 1}
	case FacingLeft:
		return Vec2{-1, 0}
	case FacingRight:
		return Vec2{1, 0}
	default:
		panic(fmt.Sprintf("unexpected day06.Facing: %#v", f))
	}
}

type Guard struct {
	Pos    Vec2
	Facing Facing
}

type Map struct {
	Grid         arrays.Array2D[byte]
	Guard        Guard
	GuardHistory arrays.Array2D[Facing]
}

func (m Map) Clone() Map {
	clone := m
	clone.Grid = clone.Grid.Clone()
	clone.GuardHistory = clone.GuardHistory.Clone()
	return clone
}

var (
	ErrMoveOutOfBounds = errors.New("move out of bounds")
	ErrMoveObstructed  = errors.New("move obstructed")
	ErrMoveStepsLimit  = errors.New("move steps limit")
	ErrMoveLoop        = errors.New("move loop")
)

func (m *Map) GuardMove() error {
	for range 10000 {
		err := m.guardTryMoveOnce()
		if errors.Is(err, ErrMoveObstructed) {
			m.guardTurn()
		} else if err != nil {
			return err
		}
	}
	return ErrMoveStepsLimit
}

func (m *Map) guardTryMoveOnce() error {
	delta := m.Guard.Facing.Delta()
	newPos := Vec2{
		X: m.Guard.Pos.X + delta.X,
		Y: m.Guard.Pos.Y + delta.Y,
	}
	if newPos.X < 0 || newPos.X >= m.Grid.Width() ||
		newPos.Y < 0 || newPos.Y >= m.Grid.Height() {
		return ErrMoveOutOfBounds
	}
	char := m.Grid.Get(newPos.X, newPos.Y)
	if char == '#' {
		return ErrMoveObstructed
	}
	pastFacing := m.GuardHistory.Get(newPos.X, newPos.Y)

	if pastFacing&m.Guard.Facing != 0 {
		return ErrMoveLoop
	}

	m.GuardHistory.Set(newPos.X, newPos.Y, pastFacing|m.Guard.Facing)
	m.Grid.Set(newPos.X, newPos.Y, 'X')
	m.Guard.Pos = newPos
	return nil
}

func (m *Map) guardTurn() {
	switch m.Guard.Facing {
	case FacingUp:
		m.Guard.Facing = FacingRight
	case FacingDown:
		m.Guard.Facing = FacingLeft
	case FacingLeft:
		m.Guard.Facing = FacingUp
	case FacingRight:
		m.Guard.Facing = FacingDown
	default:
		panic(fmt.Sprintf("unexpected Facing: %#v", m.Guard.Facing))
	}
}

func ParseMap(file io.Reader) (Map, error) {
	b, err := io.ReadAll(file)
	if err != nil {
		return Map{}, err
	}
	lines := bytes.Split(bytes.TrimSpace(b), []byte{'\n'})
	width := len(lines[0])
	height := len(lines)

	grid := arrays.New2DFilled[byte](width, height, '.')
	var guard Guard

	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '#':
				grid.Set(x, y, '#')
			case '^':
				guard.Pos = Vec2{x, y}
				guard.Facing = FacingUp
				grid.Set(x, y, 'X')
			}
		}
	}

	return Map{
		Grid:         grid,
		Guard:        guard,
		GuardHistory: arrays.New2DFilled(width, height, Facing(0)),
	}, nil
}
