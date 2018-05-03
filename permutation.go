package permutation

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type Permutator struct {
	isFirst bool // It is first state of elements
	v       Interface
	b       []int
}

func New(v Interface) *Permutator {
	return &Permutator{
		isFirst: true,
		v:       v,
		b:       make([]int, v.Len()),
	}
}

func (p *Permutator) Next() bool {

	if p.isFirst {
		p.isFirst = false
		return true
	}

	n := flipSize(p.b)

	if k := p.v.Len(); n > k {

		// It is last flip. It helps to return in begin state.
		flip(p.v, k)
		p.isFirst = true

		return false // End of permutations
	}

	flip(p.v, n) // It is the main flip.

	return true
}

func flipSize(b []int) int {
	for i := range b {
		b[i]++
		if b[i] < i+2 {
			return i + 2
		}
		b[i] = 0
	}
	return len(b) + 1
}

// flip is a function for make flip first n elements in slice (v)
func flip(v Interface, n int) {
	i, j := 0, n-1
	for i < j {
		v.Swap(i, j)
		i, j = i+1, j-1
	}
}
