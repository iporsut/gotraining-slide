package calc

import "testing"

func TestAdd(t *testing.T) {
	n := Add(10, 20)
	if n != 30 {
		t.Errorf("Expect 30 but got %d", n)
	}
}
