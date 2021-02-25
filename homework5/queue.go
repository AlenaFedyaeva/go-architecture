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

func (s *Queue) Push(elem int) *Node{
	return s.list.Append(elem)
}
func (s *Queue) Preppend(elem int) *Node{
	return s.list.Preppend(elem)
}
func (s *Queue) FindNode(data int) (*Node,error) {
	node,err:=s.list.Find(data)
	return node,err
}

func (s *Queue) Pop() int {
	if s.list.Len() == 0 {
		return 0
	}
	elem := s.list.Head().Data
	s.list.Delete(s.list.Head().Data)
	return elem
}

func (s *Queue) Print(){
	fmt.Println("Queue " )
	s.list.Print()
	fmt.Println("\n-----------------" )
}

//Delete if exist, do nothing if not exist
func (s *Queue) Delete(data int)  {
	s.list.Delete(data)
}