package permutation

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {

	var ass = [][]int{
		nil,
		[]int{},
		[]int{1},
		[]int{1, 2},
		[]int{1, 2, 3},
		[]int{1, 2, 3, 4},
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5, 6},
	}

	for _, as := range ass {
		testIntSlice(t, as)
	}
}

func cloneInts(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func factorial(x int) int {
	if x < 2 {
		if x < 0 {
			panic("negative factorial")
		}
		return 1 // 0! = 1! = 1
	}
	return x * factorial(x-1)
}

func testIntSlice(t *testing.T, as []int) {

	i := 0

	var vs [][]int

	p := New(IntSlice(as))
	for p.Next() {
		for j, v := range vs {
			if equalIntSlices(as, v) {
				t.Fatalf("v(%d) == v(%d)", j, i)
			}
		}
		i++
		vs = append(vs, cloneInts(as))
	}

	if n := factorial(len(as)); i != n {
		t.Logf("factorial invalid value: %d != %d", i, n)
	}
}

func equalIntSlicesV1(a, b []int) bool {
	return reflect.DeepEqual(a, b)
}

func equalIntSlicesV2(a, b []int) bool {
	n := len(a)
	if len(b) != n {
		return false
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

var (
	//equalIntSlices = equalIntSlicesV1
	equalIntSlices = equalIntSlicesV2
)
