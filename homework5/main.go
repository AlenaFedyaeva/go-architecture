package main

import "fmt"

func queue() {
	//1) Fill queue with values
	queue := NewQueue(2)
	node := queue.Push(1)
	node = queue.Push(5)
	node = queue.Push(6)
	node = queue.Preppend(9)
	node = queue.Push(7)
	queue.Print()

	//2) Find existing node
	node, err := queue.FindNode(7)
	if err == nil {
		fmt.Println("Found node: ", node.Data, err)
	}

	// 3) Find wrong number
	node, err = queue.FindNode(77)
	if err == nil {
		fmt.Println("find ", node.Data, err)
	} else {
		fmt.Println("ERR: ",77,err.Error())
	}

	// 4) delete number in the middle
	queue.Delete(1)
	queue.Print()

	// 5) delete head
	queue.Delete(9)
	queue.Print()

	// 5) delete tail
	queue.Delete(7)
	queue.Print()
	queue.Delete(6)
	queue.Print()
	queue.Delete(5)
	queue.Print()
}

func stack(){
	//Добавляяем значения в стек
	stack := NewStack(2)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Print()
	// Достаем из стека
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Print()
}

func main() {
    queue()
	stack()
	
}
