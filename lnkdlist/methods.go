package lnkdlist

import "fmt"

// Type Definitions:

type datatype string;

type LinkedList struct {
	head *Node
	length int
}

type Node struct {
	data datatype
	next *Node
}

// Print prints the whole linked list X -> Y -> Z
func (l *LinkedList) Print() {
	current := l.head
	for {
		fmt.Printf("%s",current.data);
		if current.next == nil {
			fmt.Printf("\n")
			break;
		} else {
			current = current.next;
			fmt.Printf(" -> ")
		}
	}
}

// Append adds a new Node in the end of the linked list.
func (l *LinkedList) Append(newData datatype) {

	node_to_be_added := Node {
			data: newData,
	};

	if l.head == nil { // Linked List is empty.
		l.head = &node_to_be_added
	} else {
		current := l.head;
		for current.next != nil {
			current = current.next
		}
		current.next = &node_to_be_added
	}

	l.length = l.length+1
}

// Contains searches for a value in each node of the linked list
// and returns true if it finds it.
func (l *LinkedList) Contains(value datatype) bool{
	// If list is empty no point in searching.
	if l.head == nil {
		return false;
	}

	current := l.head;
	found := false
	for {
		if current.data == value {
			found = true;
			break;
		} else {
			if current.next != nil {
				current = current.next;

			} else {
				break;
			}
		}
	}

	return found;
}

// Pop returns the last Node and deletes it from the list
func (l *LinkedList) Pop() *Node {
	// If list is empty no point in searching.
	if l.head == nil {
		return nil;
	}

	var current *Node = l.head;
	var prev  *Node = nil;
	for current.next != nil {
		prev = current;
		current = current.next;
	}
	prev.next = nil;
	return current;
}