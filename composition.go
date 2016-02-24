package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Perimeter() float64 {
	return (2 * r.width) + (2 * r.height)
}

type NamedRectangle struct {
	rectangle *Rectangle
	name      string
}

func (r *NamedRectangle) Area() float64 {
	return r.rectangle.Area()
}

func (r *NamedRectangle) Perimeter() float64 {
	return r.rectangle.Perimeter()
}

func (r *NamedRectangle) Name() string {
	return "My name is: " + r.name
}

func main() {
	r := &NamedRectangle{
		name: "Bob",
		rectangle: &Rectangle{
			width:  10,
			height: 20,
		},
	}
	fmt.Println(r.Name())
	fmt.Printf("Area: %v\n", r.Area())
	fmt.Printf("Perimeter: %v\n", r.Perimeter())
}
