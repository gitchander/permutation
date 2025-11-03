package permutation

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// flip is a function which flips first n elements in the slice (v)
func flip(v Interface, n int) {
	i, j := 0, n-1
	for i < j {
		v.Swap(i, j)
		i, j = i+1, j-1
	}
}

//------------------------------------------------------------------------------

type onlyLen int

func (x onlyLen) Len() int    { return int(x) }
func (onlyLen) Swap(i, j int) {}

var _ Interface = onlyLen(0)

func OnlyLen(n int) Interface {
	return onlyLen(n)
}

//------------------------------------------------------------------------------
