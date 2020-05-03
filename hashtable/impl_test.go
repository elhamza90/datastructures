package hashtable


import "testing"
//import "fmt"

func TestNew(t *testing.T) {
	tb := New(5);
	if len(*tb) != 5 || cap(*tb) != 5 {
		t.Error("Error creating new Hash Table with capacity 5");
	}
}

func TestSet(t *testing.T) {
	cap := 15;
	tb := New(cap);
	// Insert (Key, Val)
	key := "Random";
	val := 22;
	tb.Set(string(key), val);
	
	// Check if it was inserted according to hashing algorithm
	i := tb.hash(key);
	if (*(*tb)[i]).val != val {
		t.Error("Key Value Insertion Error");
	}
}

func TestGet(t *testing.T) {
	tb := New(5);
	tb.Set("A", 11);
	v, f := (*tb).Get("A");
	if v != 11 || f != true {
		t.Error("Hash Table Get function return value unexpected")
	}
}