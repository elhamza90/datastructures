package lnklist

// This file contains the Implementation of a Linked List
// and its basic utility functions:
// String(), Length(), Push(), Contains() and Pop()

import "strings"


// Type Definitions:

type datatype string;

type LinkedList struct {
	head *node
}

type node struct {
	data datatype
	next *node
}

// _________________________________

// New creates a linked list with arbitrary number of initial values
// Order of the values is taken into account.
func New(vals ...datatype) (l *LinkedList) {
	l = &LinkedList{};
	var cur, n *node = nil, nil;
	for _, v := range vals {
		n = &node {
			data: v,
		};
		if cur == nil {
			cur = n;
			l.head = n;
		} else {
			cur.next = n;
			cur = cur.next;
		}
	}
	return l;
}

// String prints the whole linked list X -> Y -> Z
func (l *LinkedList) String() (s string) {
	current := l.head;
	s = "";
	for {
		s += string(current.data);
		if current.next == nil {
			s += "\n";
			break;
		} else {
			current = current.next;
			s += " -> ";
		}
	}
	return strings.TrimSpace(s);
}

func (l *LinkedList) Length() int {
	if l.head == nil {
		return 0;
	}
	current := l.head;
	length := 1;
	for current.next != nil {
		current = current.next;
		length++;
	}
	return length;
}

// Append adds a new node in the end of the linked list.
func (l *LinkedList) Push(newData datatype) {

	node_to_be_added := node {
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

// Pop returns the last node and deletes it from the list
func (l *LinkedList) Pop() *node {
	// If list is empty no point in searching.
	if l.head == nil {
		return nil;
	}

	var current *node = l.head;
	var prev  *node = nil;
	for current.next != nil {
		prev = current;
		current = current.next;
	}
	prev.next = nil;
	return current;
}