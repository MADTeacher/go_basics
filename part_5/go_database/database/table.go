package database

import (
	"fmt"
	"strconv"
	"strings"
)

type Node[T IAttribute] struct {
	Data T
	Next *Node[T]
}

type Table[T IAttribute] struct {
	head  *Node[T]
	tail  *Node[T]
	Title string
}

func NewTable[T IAttribute](title string) *Table[T] {
	return &Table[T]{Title: title}
}

func (t *Table[T]) First() T {
	if t.head != nil {
		return t.head.Data
	}
	var zero T
	return zero
}

func (t *Table[T]) Last() T {
	if t.tail != nil {
		return t.tail.Data
	}
	var zero T
	return zero
}

func (t *Table[T]) Contains(data T) bool {
	user, ok := any(data).(*User)
	if !ok {
		return false
	}

	temp := t.head
	for temp != nil {
		if temp.Data.Check(FieldID, strconv.Itoa(user.ID())) {
			return true
		}
		temp = temp.Next
	}
	return false
}

func (t *Table[T]) Insert(data T) bool {
	if t.Contains(data) {
		return false
	}
	node := &Node[T]{Data: data}
	if t.head == nil {
		t.head = node
		t.tail = node
	} else {
		t.tail.Next = node
		t.tail = node
	}
	return true
}

func (t *Table[T]) Remove(id string) {
	var prev *Node[T]
	curr := t.head

	for curr != nil {
		if curr.Data.Check("id", id) {
			if prev == nil {
				t.head = curr.Next
				if t.head == nil {
					t.tail = nil
				}
			} else {
				prev.Next = curr.Next
				if prev.Next == nil {
					t.tail = prev
				}
			}
			return
		}
		prev = curr
		curr = curr.Next
	}
}

func (t *Table[T]) Intersect(attribute, value string,
	other *Table[T]) *Table[T] {

	newTable := NewTable[T](fmt.Sprintf("%s-%s", t.Title, other.Title))

	for temp := t.head; temp != nil; temp = temp.Next {
		if temp.Data.Check(attribute, value) {
			newTable.Insert(temp.Data)
		}
	}
	for temp := other.head; temp != nil; temp = temp.Next {
		if temp.Data.Check(attribute, value) &&
			!newTable.Contains(temp.Data) {
			newTable.Insert(temp.Data)
		}
	}
	return newTable
}

func (t *Table[T]) Union(other *Table[T]) *Table[T] {
	newTable := NewTable[T](fmt.Sprintf("%s-%s", t.Title, other.Title))

	for temp := t.head; temp != nil; temp = temp.Next {
		newTable.Insert(temp.Data)
	}
	for temp := other.head; temp != nil; temp = temp.Next {
		if !newTable.Contains(temp.Data) {
			newTable.Insert(temp.Data)
		}
	}
	return newTable
}

func (t *Table[T]) Selection(attribute, value string) *Table[T] {
	newTable := NewTable[T](t.Title + "-new")

	for temp := t.head; temp != nil; temp = temp.Next {
		if temp.Data.Check(attribute, value) {
			newTable.Insert(temp.Data)
		}
	}
	return newTable
}

func (t *Table[T]) ForEach(action func(T)) {
	for temp := t.head; temp != nil; temp = temp.Next {
		action(temp.Data)
	}
}

func (t *Table[T]) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s%s%s\n", strings.Repeat("*", 10),
		t.Title, strings.Repeat("*", 10)))

	for temp := t.head; temp != nil; temp = temp.Next {
		sb.WriteString(fmt.Sprintf("%s\n", temp.Data.String()))
	}

	return sb.String()
}
