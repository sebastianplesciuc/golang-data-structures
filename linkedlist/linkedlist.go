package linkedlist

// Value is the interface a value needs to implements to be stored in the linked list
type Value interface {
	Equals(Value) (bool, error)
}

// List interface wraps and exports the listed list object
type List interface {
	Append(data Value)

	Begin() Iterator
	End() Iterator
}

// Iterator is actually a wrapper for a linked list node
type Iterator interface {
	Get() Value
	Next() Iterator
	Previous() Iterator
}

// listNode is the struct for a linked list node
type listNode struct {
	data Value

	next     *listNode
	previous *listNode

	// I hate this. This might violate the D in SOLID
	parentList *listContainer
}

// listContainer is the struct for the linked list (parent object)
type listContainer struct {
	first *listNode
	last  *listNode

	lastIterator *listNode
}

// Get returns the Value stored in a node
func (n *listNode) Get() Value {
	return n.data
}

// Next returns the reference to the next node or the end interator
func (n *listNode) Next() Iterator {
	if n.next == nil {
		return n.parentList.lastIterator
	}

	return n.next
}

// Previous returns the reference to the previous node
func (n *listNode) Previous() Iterator {
	return n.previous
}

// Begin returns a reference to the first node in the list
func (ll *listContainer) Begin() Iterator {
	return ll.first
}

// End returns a reference to the last node in the list
func (ll *listContainer) End() Iterator {
	return ll.lastIterator
}

// Append adds an node to the end of the list
func (ll *listContainer) Append(data Value) {
	// If this is the first value, create the first node
	if ll.first == nil {
		ll.first = &listNode{
			data:       data,
			next:       nil,
			previous:   nil,
			parentList: ll,
		}

		ll.last = ll.first

		// The last iterator needs to be an element without data that represents End()
		ll.lastIterator = &listNode{
			next:       nil,
			previous:   ll.last,
			parentList: ll,
		}

		return
	}

	ll.last.next = &listNode{
		data:       data,
		previous:   ll.last,
		next:       nil,
		parentList: ll,
	}

	ll.last = ll.last.next
	ll.lastIterator = &listNode{
		next:       nil,
		previous:   ll.last,
		parentList: ll,
	}
}

// New is the factory function for a linked List
func New() List {
	return &listContainer{}
}
