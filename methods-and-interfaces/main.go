package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// a method is just a function with a receiver argument
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pointer receiver
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v2 := Vertex{3, 4}
	// 簡寫自 (&v2).Scale(10)
	// 實際上 receiver receiver
	v2.Scale(10)
	fmt.Println(v2.Abs())

	v3 := &Vertex{3, 4}
	ScaleFunc(v3, 8)
	fmt.Print(*v3)

	// Read to 6
}
