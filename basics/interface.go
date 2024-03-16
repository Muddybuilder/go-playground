package main

import "fmt"

type Point struct {
	x float32
	y float32
}

type Vehicle struct {
	velocity float32
	Point
}

type Spaceship interface {
	fly()
	land()
	position() Point
}

func (v *Vehicle) fly() {
	v.velocity = 10
}

func (v *Vehicle) land() {
	v.velocity = 0
}

func (v Vehicle) position() Point {
	return Point{v.x, v.y}
}

func main() {
	v := Vehicle{0, Point{0, 0}}

	v.fly()
	fmt.Println(v.velocity)
	v.land()
	fmt.Println(v.velocity)
}
