package main

import (
	"fmt"
	"math"
	"strconv"
)

type Vertex struct {
	X, Y float64
}

// a method is just a function with a receiver argument
func (v *Vertex) Abs() float64 {
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

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Print(f)
}

// type switches
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bypes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

// implement Stringer interface defined by the fmt package
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	str := ""
	for i := 0; i < len(ip); i++ {
		str += strconv.Itoa(int(ip[i]))
		if i != (len(ip) - 1) {
			str += "."
		}
	}
	return str
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
	fmt.Println(*v3)

	var a Abser
	f2 := MyFloat(-math.Sqrt2)
	v4 := Vertex{3, 4}

	a = f2
	a = &v4

	// error
	// Vertex (the value type) doesn't implement Abser because the Abs method is define only on *Vertex (the pointer type)
	// a = v4

	fmt.Println(a.Abs())

	var i I = &T{
		S: "hello",
	}
	i.M()

	var i2 I
	i2 = &T{"hello"}
	describe(i2)
	i2.M()

	i2 = F(math.Pi)
	describe(i2)
	i2.M()

	var i3 I
	var t *T
	i3 = t
	describe(i3)
	i3.M()

	// An empty interface may hold values of any type
	var i4 interface{}
	describe(i4)

	i4 = 42
	describe(i4)

	i4 = "Hello"
	describe(i4)

	// type assertions
	var i5 interface{} = "hello"
	s := i5.(string)
	fmt.Println(s)

	s, ok := i5.(string)
	fmt.Println(s, ok)

	f3, ok := i5.(int)
	fmt.Println(f3, ok)

	do(21)
	do("hello")
	do(true)

	a2 := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a2, z)

	ip := IPAddr{1, 2, 3, 4}
	fmt.Println(ip)

	// Read to 18
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
