package main

import (
	"testing"
)


func TestBinary(t *testing.T){
	arr:=[]int{8,14,26,28,47,56,60,64,69,70,78,80,82,84,87,90,92,98}

	//Существующий элемент в массиве
	key:=84
	var expected int=13

	num,_:=binary_search(arr, key)
	if *num!=expected{
		t.Error("error: wrong num  expected  ",*num, expected,arr[*num] )
	}


	//Существующий элемент в массиве - первый
	key=8
	expected=0
	num,_=binary_search(arr, key)
	if *num!=expected{
		t.Error("error: wrong num  expected  ",*num, expected,arr[*num] )
	}


	//Существующий элемент в массиве - последний
	key=98
	expected=17
	num,_=binary_search(arr, key)
	if *num!=expected{
		t.Error("error: wrong num  expected  ",*num, expected,arr[*num] )
	}
	
	//Не существующий элемент
	key=25
	var exp *int
	num,_=binary_search(arr, key)
	if exp!=nil{
		t.Error("error: wrong num  expected  ",*num, expected,arr[*num] )
	}
	
}