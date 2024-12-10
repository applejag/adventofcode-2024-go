package day10

import "fmt"

type Vec2 struct {
	X int
	Y int
}

func (v Vec2) String() string {
	return fmt.Sprintf("{%d %d}", v.X, v.Y)
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

func (v Vec2) Scale(factor int) Vec2 {
	return Vec2{v.X * factor, v.Y * factor}
}
