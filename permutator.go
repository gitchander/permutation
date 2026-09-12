package permutation

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type Permutator struct {
	v Interface
	b []int
}

func NewPermutator(v Interface) *Permutator {
	return &Permutator{
		v: v,
		b: make([]int, v.Len()),
	}
}

// Next generates the next permutation and returns true if more permutations remain.
func (p *Permutator) next() bool {
	n, ok := calcFlipSize(p.b)
	flip(p.v, n)
	return ok
}

// func (p *Permutator) NextPermutation() bool {
// 	return p.next()
// }

// WalkPermutations iterates over all permutations, calling the provided walk function.
func (p *Permutator) WalkPermutations(wf WalkFunc) {
	walkPermutations(p, wf)
}

// calcFlipSize calculates the size of the next flip using the auxiliary array b.
func calcFlipSize(b []int) (int, bool) {
	n := len(b)
	// Start from index 1 because b[0] is always 0.
	for i := 1; i < n; i++ {
		b[i]++
		if b[i] < (i + 1) {
			return i + 1, true // Main flip operation.
		}
		b[i] = 0
	}
	// Final flip (n elements) to return the elements to their initial state.
	return n, false // End of permutations.
}

// flip reverses the order of the first n elements in the collection (v).
func flip(v Interface, n int) {
	i, j := 0, n-1
	for i < j {
		v.Swap(i, j)
		i, j = i+1, j-1
	}
}
