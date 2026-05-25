package main

import (
	"fmt"

	searchinsertposition "github.com/aashishdubey1/dsa/binary-search/search-insert-position"
)

func main(){
	fmt.Println("Hola dsa")

	arr := []int{4,5,7,8,12,15,17,19,20}

	target :=  6

	result := searchinsertposition.SearchInsert(arr,target)

	fmt.Println(result)
}