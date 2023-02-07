package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	top *Node
}

func (s *Stack) Push(value int) {
	s.top = &Node{value, s.top}
}

func (s *Stack) Pop() int {
	if s.top == nil {
		return 0
	}
	value := s.top.value
	s.top = s.top.next
	return value
}

func (s *Stack) Peek() int {
	if s.top == nil {
		return 0
	}
	return s.top.value
}
func (s *Stack) Clear() {
	s.top = nil
}
func (s *Stack) Contains() bool {
	if s.top == nil {
		return false
	} else {
		return true
	}
}
func (s *Stack) Increment() {
	cur := s.top
	for cur != nil {
		cur.value += 1
		cur = cur.next
	}
}
func (s *Stack) PrintAll() {
	cur := s.top
	for cur != nil {
		fmt.Println(cur.value)
		cur = cur.next
	}
}

func (s *Stack) Reverse() {
	cur := s.top
	var arr []int
	for cur != nil {
		arr = append(arr, cur.value)
		cur = cur.next
	}
	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}
}

func main() {
	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Peek())
	fmt.Println(s.Pop())
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Pop()
	fmt.Println(s.Peek())
	s.Push(10)
	s.Push(20)
	s.Push(30)
	s.Increment()
	s.PrintAll()
	s.ReversePrint()
	//fmt.Println(s.Peek())
	//fmt.Println(s.Contains())
	//fmt.Println(s.Peek())
	s.Clear()
	fmt.Println(s.Contains())

	//s.Pop()
	//fmt.Println(s.Peek())
}
