package main

import "fmt"

type Queue struct {
	list *List
}

func NewQueue(cap int) *Queue {
	return &Queue{
		list: &List{},
	}
}

func (s *Queue) Push(elem int) {
	node := &Node{
		Data: elem,
	}
	s.list.Append(node)
}
func (s *Queue) Pop() int {
	if s.list.Len() == 0 {
		return 0
	}
	elem := s.list.Head().Data
	s.list.Delete(s.list.Head())
	return elem
}

func (s *Queue) Print(){
	fmt.Println("Queue " )
	node:=s.list.head
	for node!=nil {
		fmt.Print(node.Data, "  " )
		node=node.next
	}
	// for tmp := s.list.Head(); tmp != s.list.Tail(); tmp = tmp.next {
	// 	fmt.Print(tmp.Data, "  " )
	// }
	fmt.Println("\n-----------------" )
}
