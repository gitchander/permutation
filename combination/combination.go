package combination

// Combination
// https://en.wikipedia.org/wiki/Combination

// C(n, k)
// (n >= k)
type Combinator struct {
	n int // length of set
	k int // length of subset

	indexes []int // subset indexes
}

func New(n, k int) *Combinator {
	if k > n {
		panic("combination.New: k > n")
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

func (c *Combinator) Indexes() []int {
	return c.indexes
}

func (c *Combinator) RangeIndexes(f func(subsetIndex, setIndex int) bool) {
	for subsetIndex, setIndex := range c.indexes {
		if !f(subsetIndex, setIndex) {
			return
		}
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
