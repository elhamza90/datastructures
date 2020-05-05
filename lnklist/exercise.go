package lnklist

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