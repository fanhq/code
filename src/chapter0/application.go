package main

import(
	"fmt"
)


func main(){
	fmt.Println("hello world")
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)
	newSlice := slice[1:3]
	fmt.Println(newSlice)
	newSlice1 := newSlice[1:3];
	fmt.Println(newSlice1);
}