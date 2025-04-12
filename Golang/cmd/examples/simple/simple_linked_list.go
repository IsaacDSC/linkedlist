package main

import (
	"fmt"
	"linkedlist/Golang/algoritms"
)

func main() {
	list := algoritms.LinkedList[int]{}

	list.Append(10)
	list.Unshift(5)
	list.AppendOnIndex(7, 1)

	list.Debug() // 5 -> 7 -> 10 -> nil

	index, isFound := list.Search(7)
	fmt.Println("Índice do 7:", isFound, index) // 1

	list.Delete(7)
	list.Debug() // 5 -> 10 -> nil
}
