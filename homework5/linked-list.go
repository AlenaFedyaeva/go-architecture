package main

import (
	"errors"
	"fmt"
)

type Node struct {
	next *Node
	prev *Node
	Data int
}

type List struct {
	len  int
	head *Node
	tail *Node
}

func (list *List) Head() *Node {
	return list.head
}

func (list *List) Tail() *Node {
	return list.tail
}

func (list *List) Len() int {
	return list.len
}

func (list *List) Find(elem int) (*Node, error) {

	if list.head != nil {
		for tmp := list.head; tmp != nil; tmp = tmp.next {
			if tmp.Data == elem {
				return tmp, nil
			}
		}
	}
	return nil, errors.New("Node not found")
}


func (list *List) Append(elem int)*Node{
	node := &Node{
		Data: elem,
	}
	list.len++
	if list.head != nil {
		tail := list.tail
		tail.next = node
		node.prev = tail
		node.next = nil
		list.tail = node
	}
	if list.head == nil { //1 elem
		node.next = nil
		node.prev = nil
		list.head = node
		list.tail = node
	}
	return node
}

func (list *List) Preppend(elem int) *Node{
	node := &Node{
		Data: elem,
	}
	list.len++
	if list.head != nil {
		list.head.prev=node
		node.next = list.head
		node.prev = nil
		list.head = node
	}
	if list.head == nil { //1 elem
		node.next = nil
		node.prev = nil
		list.head = node
		list.tail = node
	}
	
	return node

}

func (list *List) Delete(data int) {

	node, err := list.Find(data)
	if err != nil {
		return
	}

	list.len--

	nodePre := node.prev
	nodeNext := node.next

	if nodeNext!=nil{
		nodeNext.prev = nodePre
	}
	if nodePre!=nil{
			nodePre.next = nodeNext
	}


	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	}

	if node == list.head {
		list.head = nodeNext
		nodeNext.prev = nil
	}
	if node == list.tail {
		list.tail = nodePre
		nodePre.next = nil
	}

}

func (list *List) Print() {
	node:=list.head
	for node!=nil {
		fmt.Print(node.Data, "  " )
		node=node.next
	}
}

