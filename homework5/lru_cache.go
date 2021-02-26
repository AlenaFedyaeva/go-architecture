package main

import "fmt"

type LRUCache struct{
	list *List
	maxSize int
} 

func NewLRUCache(maxSize int) *LRUCache {
	return &LRUCache{
		list: &List{},
		maxSize: maxSize,
	}
}

func (c *LRUCache) Push(elem int) {

	node,err:=c.list.Find(elem)
	// not present in cache
	if  err!=nil {
		//chache is full
		if c.list.len==c.maxSize{
			//delete least recently used element
			c.list.Delete(c.list.tail.Data)
		}
	}

	//present in cache
	if node!=nil{
		c.list.Delete(elem)
	}

	c.list.Preppend(elem)
}

func (c *LRUCache) Print() {
	fmt.Println("LRUCache" )
	c.list.Print()
	fmt.Println("\n-----------------" )

}

// func (c *LRUCache) Refer(elem int) {
// 	c.list.Append(elem)
// }