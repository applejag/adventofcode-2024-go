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
