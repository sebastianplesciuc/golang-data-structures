package main

import (
	"fmt"

	"github.com/sebastianplesciuc/golang-data-structures/linkedlist"
)

func main() {
	list := linkedlist.New()

	list.Append("data1")
	list.Append("data2")

	iterator := list.Begin()
	for iterator != list.End() {
		fmt.Println(iterator.Get())
		iterator = iterator.Next()
	}

}
