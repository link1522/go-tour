package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// 參數型別都一樣可以簡寫
// 以下等同於 func add(x int, y int) int
func add(x, y int) int {
	return x + y
}

// 可以有多個返回值
func swap(x, y string) (string, string) {
	return y, x
}

// Naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool
var csharp bool = true

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Pi)

	fmt.Println(add(1, 2))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	var i int
	fmt.Println(c, python, java, i)

	var j, k int = 1, 2
	fmt.Println(j, k)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// 預設初始值
	var ii int
	var f float64
	var bo bool
	var s string
	fmt.Printf("%v %v %v %q\n", ii, f, bo, s)

	// type conversion
	var c, d int = 3, 4
	var f2 float64 = math.Sqrt(float64(c*c + d*d))
	var z uint = uint(f2)
	fmt.Println(c, d, z)

	const PI = 3.14
	fmt.Println(PI)
}
