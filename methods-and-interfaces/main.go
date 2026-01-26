package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
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

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it did't work",
	}
}

type ErrNagtiveSqrt float64

func (e ErrNagtiveSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func exSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNagtiveSqrt(x)
	}

	return math.Sqrt(x), nil
}

type MyReader struct{}

func (reader MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(p []byte) (int, error) {
	n, err := reader.r.Read(p)
	for i := 0; i < n; i++ {
		b := p[i]
		switch {
		case b >= 'a' && b <= 'z':
			p[i] = 'a' + (b-'a'+13)%26
		case b >= 'A' && b <= 'Z':
			p[i] = 'A' + (b-'A'+13)%26
		}
	}
	return n, err
}

type Image struct {
	w, h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.h, img.w)
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
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

	if err := run(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(exSqrt(2))
	fmt.Println(exSqrt(-2))

	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	s2 := strings.NewReader("Lbh penpxrq gur pbqr!")
	r2 := rot13Reader{s2}
	io.Copy(os.Stdout, &r2)

	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
