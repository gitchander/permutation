package permutation

type anySlice[E any] []E

func (x anySlice[E]) Len() int      { return len(x) }
func (x anySlice[E]) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func NewSlicePermutator[S ~[]E, E any](x S) *Permutator {
	return NewPermutator(anySlice[E](x))
}

// SliceI
func sliceToInterface[S ~[]E, E any](x S) Interface {
	return anySlice[E](x)
}

func _() {
	type Ints []int
	sliceToInterface(Ints([]int{}))
}
