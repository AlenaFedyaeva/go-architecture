package main

import "fmt"



func main() {
	queue := NewQueue(2)
	queue.Push(5)
	queue.Push(6)
	queue.Push(7)
	queue.Print()
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())


	// stack := NewStack(2)
	// stack.Push(5)
	// stack.Push(6)
	// stack.Push(7)
	// stack.Print()
	// fmt.Println(stack.Pop()) 
	// stack.Print()
	// fmt.Println(stack.Pop())
	// stack.Print()
	// fmt.Println(stack.Pop())
}
