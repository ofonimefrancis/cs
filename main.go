package main

import (
	"fmt"

	"github.com/ofonimefrancis/cs-fundamentals/foundations/DS/stacks"
)

func main() {
	s := stacks.New()
	s.Push(12)
	s.Push(90)
	s.Push(81)
	s.Push(12)
	s.Push(21)
	s.Push(50)

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
