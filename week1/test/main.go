package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

// func (m *MultiShape) perimeter() int {
// var perim int
//
// return l + w *2
// }

func main() {
	//var l,w  int= 6,7
	c := Circle{0, 0, 5}
	m := MultiShape{
		shapes: []Shape{&c, &c},
	}
	fmt.Println(m.area())

}
