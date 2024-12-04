package solutions

import "io"

var Days = map[int]Day{
	3: Day03{},
}

type Day interface {
	Part1(file io.Reader) (any, error)
	Part2(file io.Reader) (any, error)
}
