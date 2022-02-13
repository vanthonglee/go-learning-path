package main

import (
	"fmt"
	"reflect"
)

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{height: 5.0, base: 5.0}
	sq := square{sideLength: 2}

	printArea(t)
	printArea(sq)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func printArea(s shape) {
	fmt.Println("area of", reflect.TypeOf(s).Name(), "is", s.getArea())
}
