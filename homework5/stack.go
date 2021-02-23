package main

import "fmt"

type Stack struct {
	sli []int
}

func NewStack(cap int) *Stack {
	return &Stack{
		sli: make([]int, 0, cap),
	}
}

func (s *Stack) Push(elem int) {
	s.sli = append(s.sli, elem)
}
func (s *Stack) Pop() int {
	if len(s.sli) == 0 {
		return 0
	}
	elem := s.sli[len(s.sli)-1]
	s.sli = s.sli[:len(s.sli)-1]
	return elem
}

func (s *Stack) Print(){
	fmt.Println("Stack" )
	for _, v := range s.sli {
		fmt.Print(v, "  ")
	}
	
	fmt.Println("\n-----------------" )
}

