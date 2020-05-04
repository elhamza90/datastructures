package lnklist


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