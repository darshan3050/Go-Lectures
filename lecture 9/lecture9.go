package main

import "fmt"

// import "fmt"

type Inter_Face interface {
	Area() float64

	Perimeter() float64
}
type Circle struct {
	radius int
}

func AddCircle(radius int) *Circle {
	var newCircle Circle
	newCircle.radius = radius
	return &newCircle
}
func (c *Circle) Area() float64 {
	radius := float64(c.radius)
	return 3.14 * radius * radius
}
func (c *Circle) Perimeter() float64 {
	radius := float64(c.radius)
	return 3.14 * 2 * radius
}
func AddRectangle(length, breadth float64) *Rectangle {
	var newRectangle Rectangle
	newRectangle.length = length
	newRectangle.breadth = breadth
	return &newRectangle
}

type Rectangle struct {
	length  float64
	breadth float64
}

func (r *Rectangle) Area() float64 {
	return r.length * r.breadth
}
func (r *Rectangle) Perimeter() float64 {
	return 2 * r.length * r.breadth
}
func main() {
	circle1 := AddCircle(12)
	circle2 := AddCircle(24)
	AreaOfCircle1 := circle1.Area()
	PerimeterOfCircle2 := circle2.Perimeter()
	fmt.Println(AreaOfCircle1)
	fmt.Println(PerimeterOfCircle2)
	rectangle1 := AddRectangle(50, 12)
	rectangle2 := AddRectangle(50, 24)
	AreaOfrectangle1 := rectangle1.Area()
	PerimeterOfrectangle2 := rectangle2.Perimeter()
	fmt.Println(AreaOfrectangle1)
	fmt.Println(PerimeterOfrectangle2)
}
