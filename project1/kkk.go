package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println("with val")
	for i, val := range arr {
		fmt.Println(i, val)
	}
	fmt.Println(arr)
}
