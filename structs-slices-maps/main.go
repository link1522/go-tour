package main

import (
	"fmt"
	"math"
	"strings"
)

// A struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

var (
	gv1 = Vertex{1, 2}
	// 可以直接指定填入某個欄位
	gv2 = Vertex{X: 1}
	gv3 = Vertex{}
	gp  = &Vertex{1, 2}
)

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)

	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			pic[i][j] = uint8(i * j)
		}
	}

	return pic
}

type Vertex2 struct {
	Lat, Long float64
}

var m map[string]Vertex2

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordsMap := make(map[string]int)
	for _, v := range words {
		_, ok := wordsMap[v]
		if !ok {
			// 不判斷也可以，因為預設就是 0
			wordsMap[v] = 0
		}
		wordsMap[v]++
	}
	return wordsMap
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	i, j := 42, 2701

	// The type *T is a pointer to a T value. Its zero value is nil.
	// The & operator generates a pointer to its operand.
	p := &i
	// The * operator denotes the pointer's underlying value.
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	fmt.Println(Vertex{1, 2})

	v := Vertex{2, 4}
	v.X = 3
	fmt.Println(v)

	v2 := Vertex{7, 8}
	v2p := &v2
	// 簡化自 (*v2p).X
	v2p.X = 1e9
	fmt.Println(v2)

	fmt.Println(gv1, gp, gv2, gv3)

	var a [2]string
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slices
	var s []int = primes[1:4]
	fmt.Println(s)

	// A slices does not store any data, it just describes a section of an underlying array.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a2 := names[0:2]
	b2 := names[1:3]
	fmt.Println(a2, b2)

	b2[0] = "XXX"
	fmt.Println(a2, b2)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true}
	fmt.Println(r)

	s2 := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
	}
	fmt.Println(s2)

	s3 := []int{2, 3, 5, 7, 11, 13}
	s3 = s3[1:4]
	fmt.Println(s3)

	s3 = s3[:2]
	fmt.Println(s3)

	s3 = s3[1:]
	fmt.Println(s3)

	s4 := []int{2, 3, 5, 7, 11, 13}
	s4 = s4[:0]
	printSlice(s4)

	s4 = s4[:4]
	printSlice(s4)

	s4 = s4[2:]
	printSlice(s4)

	var s5 []int
	printSlice(s5)

	a3 := make([]int, 5)
	fmt.Print("a3: ")
	printSlice(a3)

	b3 := make([]int, 0, 5)
	fmt.Print("b3: ")
	printSlice(b3)

	c3 := b3[:2]
	fmt.Print("c3: ")
	printSlice(c3)

	d3 := c3[2:5]
	fmt.Print("d3: ")
	printSlice(d3)

	board := [][]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var s6 []int
	printSlice(s6)

	s6 = append(s6, 0)
	printSlice(s6)

	s6 = append(s6, 1)
	printSlice(s6)

	s6 = append(s6, 2, 3, 4)
	printSlice(s6)

	var xx [1]int
	xx2 := xx[:0]
	xxp := &xx2
	fmt.Println(xxp)

	xx2 = append(xx2, 1, 2, 3, 4, 6)
	xxpNew := &xx2
	fmt.Println(xxpNew == xxp)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	pow2 := make([]int, 10)
	for i := range pow {
		pow2[i] = 1 << uint(i)
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

	m = make(map[string]Vertex2)
	m["Bell Labs"] = Vertex2{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	var m2 = map[string]Vertex2{
		"Bell Labs": Vertex2{
			40.68433, -74.39967,
		},
		"Google": Vertex2{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2)

	// 可以省略型別
	var m3 = map[string]Vertex2{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(m3)

	m4 := make(map[string]int)
	m4["answer"] = 42
	fmt.Println(m4["answer"])
	m4["answer"] = 48
	fmt.Println(m4["answer"])
	delete(m4, "answer")
	fmt.Println(m4["answer"])
	vv, ok := m4["answer"]
	fmt.Println("The value", vv, "Present?", ok)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}

	f := fibonacci()
	for range 10 {
		fmt.Println(f())
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// closure
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		v := a
		a, b = b, a+b
		return v
	}
}
