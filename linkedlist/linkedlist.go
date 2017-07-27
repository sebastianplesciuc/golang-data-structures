package linkedlist

// LinkedList interface
type LinkedList interface {
	Append(data string)

	Begin() Iterator
	End() Iterator
}

// Iterator bla bla
type Iterator interface {
	Get() string
	Next() Iterator
	Previous() Iterator
}

type node struct {
	data string

	next     *node
	previous *node

	// I hate this. This might violate the D in SOLID
	parentList *listContainer
}

type listContainer struct {
	first *node
	last  *node

	lastIterator *node
}

func (n *node) Get() string {
	return n.data
}

func (n *node) Next() Iterator {
	if n.next == nil {
		return n.parentList.lastIterator
	}

	return n.next
}

func (n *node) Previous() Iterator {
	return n.previous
}

func (ll *listContainer) Begin() Iterator {
	return ll.first
}

func (ll *listContainer) End() Iterator {
	return ll.lastIterator
}

func (ll *listContainer) Append(data string) {
	if ll.first == nil {
		ll.first = &node{
			data:       data,
			next:       nil,
			previous:   nil,
			parentList: ll,
		}

		ll.last = ll.first
		ll.lastIterator = &node{
			next:       nil,
			previous:   ll.last,
			parentList: ll,
		}

		return
	}

	finder := ll.first
	for finder.next != nil {
		finder = finder.next
	}

	finder.next = &node{
		data:       data,
		previous:   finder,
		next:       nil,
		parentList: ll,
	}

	ll.last = finder.next
	ll.lastIterator = &node{
		next:       nil,
		previous:   ll.last,
		parentList: ll,
	}
}

// New returns a LinkedList
func New() LinkedList {
	return &listContainer{}
}
