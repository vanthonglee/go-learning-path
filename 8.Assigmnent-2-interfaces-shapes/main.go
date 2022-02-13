package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

type shape interface {
	printArea()
}

func main() {
	t := triangle{24, 13}
	sq := square{5}

	t.printArea()
	sq.printArea()
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (sq square) getArea() float64 {
	return sq.sideLength * sq.sideLength
}

func (t triangle) printArea() {
	fmt.Println("Area of this triangle:", t.getArea())
}

func (sq square) printArea() {
	fmt.Println("Area of this square:", sq.getArea())
}
