package solutions

import (
	"errors"
	"io"
)

var ErrNotImplemented = errors.New("not implemented")

type Day interface {
	Part1(file io.Reader) (any, error)
	Part2(file io.Reader) (any, error)
}

type UnimplementedDay struct{}

var _ Day = UnimplementedDay{}

func (u UnimplementedDay) Part1(file io.Reader) (any, error) {
	return nil, ErrNotImplemented
}

func (u UnimplementedDay) Part2(file io.Reader) (any, error) {
	return nil, ErrNotImplemented
}
