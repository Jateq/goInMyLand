package main

import (
	"fmt"
	"github.com/Jateq/jateqsgopackage/workspace"
)

func main() {
	s := workspace.Student{Name: "Temer"}
	fmt.Println(s.GetName())
}
