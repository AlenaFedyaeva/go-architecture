package main

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

func (list *List) Find(elem int) *Node {
	if list.head != nil {
		for tmp := list.head; tmp != list.tail; tmp = tmp.next {
			if tmp.Data == elem {
				return tmp
			}
		}
	}
	return nil
}

func (list *List) Insert(prev *Node, node *Node) {
	list.len++

	node.next=prev.next
	prev.next=node
	node.prev=prev

	if prev == list.tail {
		list.tail = node
	}

	if prev == nil { //1 elem
		list.head = node

		node.next = nil
		node.prev = nil
		
		list.tail = node
	}	


	// if prev == nil {
	// 	node.next = l.head
	// 	l.head = node
	// 	return
	// }
	// if l.head == nil {
	// 	l.head = node
	// 	l.tail = l.head
	// 	return
	// }
	// node.next = prev.next
	// prev.next = node
	// if prev == l.tail {
	// 	l.tail = node
	// }
}

func (list *List) Append(node *Node) {
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
}

func (list *List) Preppend(node *Node) {
	list.Add(nil, node)
}

func (list *List) Delete(node *Node) {
	list.len--
	if list.head == list.tail {
		list.head = nil
		list.tail = nil
	}
	if list.head != nil {
		for tmp := list.head; tmp != list.tail; tmp = tmp.next {
			if tmp.next == node && node != list.tail {
				tmp.next = node.next
			}
			if tmp.next == node && node == list.tail {
				tmp.next = nil
				list.tail = tmp
			}
		}
	}
	if node == list.head {
		list.head = node.next
	}
}
