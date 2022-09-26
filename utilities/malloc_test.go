package utilities

import "testing"

func Test_make2D(t *testing.T) {
	m, n := 5, 5
	expect := n

	if result := len(make2D[int](n, m)); result != expect {
		t.Errorf("result = %v, expect = %v", result, expect)
	}
}
