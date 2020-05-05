package lnklist

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

const errorVal datatype = "";


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

// Count returns the number of times an element was found in the list
func (l LinkedList) Count(val datatype) (c int) {
	if l.head == nil {
		return 0;
	}

	cur := l.head;
	c = 0;
	for {
		if cur.data == val {
			c++;
		}
		if cur.next == nil {
			break;
		} else {
			cur = cur.next;
		}
	}
	return c;
}

// GetNth returns the value stored in the nth index position
func (l LinkedList) GetNth(n int) (val datatype, found bool) {
	val = errorVal;
	found = false;
	if l.head != nil {
		cur := l.head;
		i := 0;
		for {
			if i == n {
				val = cur.data;
				found = true;
				break;
			} else {
				if cur.next != nil {
					cur = cur.next;
					i++;
				} else {
					break;
				}
			}
		}
	}
	return val, found;
}

// InsertNth inserts a node in the nth index position
func (l *LinkedList) InsertNth(val datatype, n int) (done bool) {
	done = false;
	if l.head != nil {

		node_to_be_inserted := &node {
			data: val,
		}

		cur := l.head;
		i := 0;
		for {
			if cur == nil {
				break;
			} else {
				if i == (n-1) {
					if cur.next == nil {
						cur.next = node_to_be_inserted;
						done = true;
						break;
					} else {
						node_to_be_inserted.next = cur.next;
						cur.next = node_to_be_inserted;
						done = true;
						break;
					}
				} else {
					cur = cur.next;
					i++;
				}
			}
		}

	}

	return done;
}

// Append appends a list to the end of the calling list.
func (l *LinkedList) Append(l2 *LinkedList) {
	if l.head == nil {
		l.head = l2.head;
	} else {
		cur := l.head;
		for {
			if cur.next == nil {
				cur.next = l2.head;
				break;
			} else {
				cur = cur.next;
			}
		}
	}
}

// Reverse reverses a list
func (l *LinkedList) Reverse() {
	if l.head == nil || l.head.next == nil {
		return;
	}

	var prev, tmp, cur *node = nil, nil, l.head;
	for cur != nil{
		if cur.next == nil {
			l.head = cur;
		}
		tmp = cur.next;
		cur.next = prev;
		prev = cur;
		cur = tmp;
	}
}

// RemoveNth removes the node in position n
func (l *LinkedList) RemoveNth(n int) {
	if l.head == nil {
		return;
	}

	var prev, cur *node = nil, l.head;
	i := 0;
	for {
		if i == n {
			prev.next = cur.next;
			break;
		} else {
			if cur.next == nil {
				break;
			} else {
				prev = cur;
				cur = cur.next;
				i++;
			}
		}
	}
}