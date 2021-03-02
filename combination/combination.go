package combination

// Combination
// https://en.wikipedia.org/wiki/Combination

// C(n, k)
// (n >= k)
type Combinator struct {
	n int // set len
	k int // subset len

	indexes []int // subset indexes
}

func NewComb(n, k int) *Combinator {

	if k > n {
		panic("combination.NewComb: k > n")
	}

	return &Combinator{
		n: n,
		k: k,
	}
}

func (c *Combinator) Next() bool {
	if c.indexes == nil {
		c.indexes = serialInts(c.k)
		return true
	}
	if nextComb(c.indexes, c.n) {
		return true
	}
	c.indexes = nil
	return false
}

func (c *Combinator) WalkSubset(f func(index int)) {
	for _, index := range c.indexes {
		f(index)
	}
}

// [ 0, 1, 2, ... , (n-2), (n-1) ]
func serialInts(n int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = i
	}
	return a
}

func nextComb(indexes []int, n int) (ok bool) {

	if len(indexes) == 0 {
		return false
	}

	if nextComb(indexes[1:], n) {
		return true
	}

	d := indexes[0] + 1
	if d > (n - len(indexes)) {
		return false
	}
	for i := range indexes {
		indexes[i] = d + i
	}

	return true
}
