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

func NewCombinator(n, k int) *Combinator {
	if k > n {
		panic("combination.New: k > n")
	}
	return &Combinator{
		n: n,
		k: k,

		indexes: makeIntsSerial(k),
	}
}

func (c *Combinator) NextCombination() bool {
	if nextComb(c.indexes, c.n) {
		return true
	}
	setIntsSerial(c.indexes) // reset initial state
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

func setIntsSerial(a []int) {
	for i := range a {
		a[i] = i
	}
}

// [ 0, 1, 2, ... , (n-2), (n-1) ]
func makeIntsSerial(n int) []int {
	a := make([]int, n)
	setIntsSerial(a)
	return a
}

func nextComb(indexes []int, n int) bool {

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
