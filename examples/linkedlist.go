package examples

import (
	"fmt"

	"github.com/sebastianplesciuc/golang-data-structures/linkedlist"
)

type myLinkedListData struct {
	data string
}

func (m *myLinkedListData) Equals(otherValue linkedlist.Value) (bool, error) {
	other, ok := otherValue.(*myLinkedListData)
	if !ok {
		return false, fmt.Errorf("Invalid type")
	}
	return other.data == m.data, nil
}

// LinkedList runs the example for a linked list
func LinkedList() {
	list := linkedlist.New()

	d1 := &myLinkedListData{data: "data1"}
	d2 := &myLinkedListData{data: "data2"}

	list.Append(d1)
	list.Append(d2)
	list.Append(d2)

	iterator := list.Begin()
	for iterator != list.End() {
		dataVal, ok := iterator.Get().(*myLinkedListData)
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
