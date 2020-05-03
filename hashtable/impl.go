package hashtable


type node struct {
	key string
	val int
	next *node
}

type htable []*node

func New(maxCap int) *htable {
	t := make(htable, maxCap, maxCap);
	return &t;
}

func (t htable) hash(k string) int {
	s := 0;
	for _, c := range k {
		s += int(c);
	}
	return s % cap(t);
}

func (t *htable) Set(k string, v int) {
	i := t.hash(k);
	(*t)[i] = &node {
		key: k,
		val: v,
	};
}

func (t htable) Get(k string) (val int , found bool) {
	i := t.hash(k);
	cur := (t)[i];
	for  {
		if cur.key == k {
			val = cur.val;
			found = true;
			break;
		} else if cur.next != nil {
			val = 0;
			found = false;
			break;
		} else {
			cur = cur.next;
		}
	}

	return val, found
}