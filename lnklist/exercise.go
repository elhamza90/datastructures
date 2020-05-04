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