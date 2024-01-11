package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	side float32
}

func (square Square) Area() float32 {
	return square.side * square.side
}

type Circle struct {
	radius float32
}

func (circle Circle) Area() float32 {
	return float32(math.Pi) * circle.radius
}

func Shapes() {
	square := Square{10.5}
	circle := Circle{5.7}

	printShapeArea(square)
	printShapeArea(circle)
}

func printShapeArea(shape Shape) {
	fmt.Printf("Плозадь фигуры равна %.2f см\n", shape.Area())
}
