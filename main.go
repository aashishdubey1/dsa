package main

import (
	"fmt"

	bs "github.com/aashishdubey1/dsa/binary-search/binary-search"
)

func main(){
	fmt.Println("Hola dsa")

	arr := []int{4,5,7,8,12,15,17,19,20}

	target :=  5 

	result := bs.BinarySearch(arr,target)

	fmt.Println(result)
}