package Iterator

import "testing"

func TestNumArrayIterator(t *testing.T) {
	numArray := NewNumArray([]int{1, 2, 3, 4, 5})
	for numArray.HasNext() {
		item := numArray.Next()
		t.Logf("num: %d %T", item.val.(int), item.val)
	}
}
