package syntaxBase

import "testing"

func TestCompareGetMax(t *testing.T) {
	var a, b, c int
	a = 1
	b = 2
	c = CompareGetMax(a, b)
	if c != 2 {
		t.Errorf("CompareGetMax(%d, %d) = %d; want 2", a, b, c)
	}
}
