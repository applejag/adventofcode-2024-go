package solutions

import "io"

var Days = map[int]Day{
	1: Day01{},
	2: Day02{},
	3: Day03{},
	4: Day04{},
}

type Day interface {
	Part1(file io.Reader) (any, error)
	Part2(file io.Reader) (any, error)
}
