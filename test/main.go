package main

import "fmt"

type MyC struct {
	name string
	age  int
}

func (m *MyC) ShowAge() {
	fmt.Printf("My age is %d\n", m.age)
}

func main() {

	a := MyC{
		name: "goo",
		age:  1,
	}
	fmt.Println(a)

	// b := &a
	// a = MyC{
	// 	name: "g2",
	// 	age:  11,
	// }
	// fmt.Println(b)
	(&a).ShowAge()
}
