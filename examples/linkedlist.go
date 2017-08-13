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
	fmt.Println("Linked list examples")

	// List creation
	list := linkedlist.New()

	// Element types
	d1 := &myLinkedListData{data: "data1"}
	d2 := &myLinkedListData{data: "data2"}

	// Adding elements to the end of the list
	list.Append(d1)
	list.Append(d2)
	list.Append(d2)

	fmt.Println()

	// Iteration and comparison
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

	// Checking existence
	exists, err := list.Contains(d2)
	if err != nil {
		panic("Unexpected data type in list")
	}

	fmt.Printf("\nlist contains %s: %v\n", d2.data, exists)

	exists, err = list.Contains(&myLinkedListData{data: "fake"})
	if err != nil {
		panic("Unexpected data type in list")
	}

	fmt.Printf("\nlist contains fake: %v\n", exists)
}
