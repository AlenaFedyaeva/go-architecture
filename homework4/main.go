package main

import (
	"errors"
	"fmt"
)

//Сложность log2n
//Ищем индекс элемента, если он есть
func binary_search( arr []int, key int) (*int,error) {
	l:=0
	r:=len(arr)
	for !(l>=r) {
		mid:=(l+r)/2 //середина области, округл в меньшую сторону
		if key==arr[mid]{
			r=mid
			break
		}
		if arr[mid]<key {
			l=mid+1
		} else{
			r=mid
		}
		
	}
	a:=arr[r]
	if key!=a{
		return nil,errors.New("Key not found")
	}
	return &r,nil
}
func main() {
	key:=27
	arr:=[]int{8,14,26,28,47,56,60,64,69,70,78,80,82,84,87,90,92,98}
	num,err:=binary_search(arr, key)
	fmt.Println("main", arr,*num,err)
}