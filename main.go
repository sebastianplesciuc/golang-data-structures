package main

import (
	"fmt"

	"github.com/sebastianplesciuc/golang-data-structures/linkedlist"
)

type myData struct {
	data string
}

func (m *myData) Equals(otherValue linkedlist.Value) (bool, error) {
	other, ok := otherValue.(*myData)
	if !ok {
		return false, fmt.Errorf("Invalid type")
	}
	return other.data == m.data, nil
}

func main() {
	list := linkedlist.New()

	d1 := &myData{data: "data1"}
	d2 := &myData{data: "data2"}

	list.Append(d1)
	list.Append(d2)
	list.Append(d2)

	iterator := list.Begin()
	for iterator != list.End() {
		dataVal, ok := iterator.Get().(*myData)
		if !ok {
			panic("Unexpected data type in list")
		}

		equals, err := iterator.Get().Equals(d2)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%s == %s : %v \n", dataVal.data, d2.data, equals)

		iterator = iterator.Next()
	}

}
