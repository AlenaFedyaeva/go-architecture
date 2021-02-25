package main

import "fmt"

type Stack struct {
	list  *List
}

func NewStack(cap int) *Stack {
	return &Stack{
		list: &List{},
	}
}

func (s *Stack) Push(elem int) {
	s.list.Append(elem)
}
func (s *Stack) Pop() int {
	if s.list.Len() == 0 {
		return 0
	}
	elem := s.list.Head().Data
	s.list.Delete(s.list.Head().Data)
	return elem
}

func (s *Stack) Print(){
	fmt.Println("Stack" )
	s.list.Print()
	fmt.Println("\n-----------------" )
}

