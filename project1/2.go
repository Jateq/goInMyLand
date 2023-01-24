package main

import "fmt"

func main() {
	var size int = 3
	m := make([]int, 4)
	for i, val := range m {
		val = val + i
	}
	m[size] = 50
	fmt.Println(m[size])
}
