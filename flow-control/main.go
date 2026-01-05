package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// if statement can start with a short statement to execute before the condition
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func exSqrt(x float64) float64 {
	z := 1.0
	count := 0

	for math.Abs(z*z-x) > 0.001 {
		z -= (z*z - x) / (2 * z)
		count++
	}

	return z
}

func tryDefer() {
	// The deferred call's arguments are evaluated immediately, but the function call is not executed util the surrounding function returns.
	n := 0
	defer fmt.Println("world " + strconv.Itoa(n))
	n++
	fmt.Print("hello ")
}

func main() {
	sum := 0
	// go 只有 for 迴圈
	// {} 是必要的
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	// for is go's while
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))

	fmt.Println(exSqrt(2))

	fmt.Print("Go runs on ")
	// go 的 switch case 不用加 break，會自動加入
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	// switch with no condition
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	tryDefer()

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
