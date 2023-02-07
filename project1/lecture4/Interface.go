package main

import "fmt"

type List interface {
	Run(int)
	DoubleIt(int) int
}

type ArrayList struct {
	data []int
}

func (a *ArrayList) Run(val int) {
	a.data = append(a.data, val)
}

func (a *ArrayList) DoubleIt(index int) int {
	return a.data[index] * 2
}

type Vector struct {
	data []int
}

func (v *Vector) Run(val int) {
	v.data = append(v.data, val)
}

func (v *Vector) DoubleIt(index int) int {
	return v.data[index] * 2
}

func main() {
	var list List
	list = &ArrayList{}
	list.Run(1)
	list.Run(2)
	fmt.Println(list.DoubleIt(0))
	fmt.Println(list.DoubleIt(1))

	list = &Vector{}
	list.Run(3)
	list.Run(4)
	fmt.Println(list.DoubleIt(0))
	fmt.Println(list.DoubleIt(1))
}
