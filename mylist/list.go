// design choice: doubled linked list
package mylist

type Element struct {
	Value  any
	next   *Element
	prev   *Element
	belong *List
}

func (e *Element) Next() *Element {
	return e.next
}

type List struct {
	head *Element
	tail *Element
	size int
}

func New() *List {
	return new(List)
}

// Init initializes or clears the list.s
func (l *List) Init() *List {
	for e := l.head; e != nil; {
		next := e.next // save it before clearing the links
		e.next, e.prev, e.belong = nil, nil, nil
		e = next
	}
	l.head = nil
	l.tail = nil
	l.size = 0
	return l
}

func (l *List) Len() int {
	return l.size
}

// Front returns first element of list.
func (l *List) Front() *Element {
	return l.head
}

// PushBack inserts a new element to the back of the list.
func (l *List) PushBack(v any) *Element {
	e := &Element{Value: v, prev: l.tail, belong: l}
	if l.tail == nil {
		l.head = e
	} else {
		l.tail.next = e
	}
	l.tail = e
	l.size++
	return e
}

func (l *List) InsertAfter(v any, mark *Element) *Element {
	if mark == nil || mark.belong != l {
		return nil
	}
	e := &Element{Value: v, prev: mark, next: mark.next, belong: l}

	if mark.next == nil {
		l.tail = e // if mark is a tail, then e become tail
	} else {
		mark.next.prev = e //set the prev of elem behind mark
	}
	mark.next = e
	l.size++
	return e
}

func (l *List) Remove(e *Element) any {
	if e == nil {
		return nil
	}

	if e.belong != l {
		return e.Value
	}

	//check if e is head
	if e.prev == nil {
		l.head = e.next
	} else {
		e.prev.next = e.next
	}

	//check if e is tail
	if e.next == nil {
		l.tail = e.prev
	} else {
		e.next.prev = e.prev
	}
	e.next, e.prev, e.belong = nil, nil, nil
	l.size--
	return e.Value
}
