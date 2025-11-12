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

func NewSlicePermutator[S ~[]E, E any](x S) *Permutator {
	return NewPermutator(anySlice[E](x))
}

// NextPermutation does (makes) next permutation.
func (p *Permutator) NextPermutation() bool {

	if n, ok := calcFlipSize(p.b); ok {
		flip(p.v, n) // It is the main flip.
		return true
	}

	// It is the last flip. It helps to return the elements to the begin state.
	flip(p.v, p.v.Len())

	return false // End of permutations.
}

func (p *Permutator) WalkPermutations(wf WalkFunc) {
	walkPermutations(p, wf)
}

func calcFlipSize(b []int) (int, bool) {
	n := len(b)
	// Start from 1 because b[0] is always 0
	for i := 1; i < n; i++ {
		if b[i] < i {
			b[i]++
			return i + 1, true
		}
		b[i] = 0
	}
	return 0, false
}

// flip is a function which flips first n elements in the slice (v)
func flip(v Interface, n int) {
	i, j := 0, n-1
	for i < j {
		v.Swap(i, j)
		i, j = i+1, j-1
	}
}
