package permutation

type Seq0 func(yield func() bool)

func (p *Permutator) Permutations() Seq0 {
	return func(yield func() bool) {
		for ok := true; ok; ok = p.next() {
			if !yield() {
				return
			}
		}
	}
}

func IPermutations(v Interface) Seq0 {
	p := NewPermutator(v)
	return p.Permutations()
}

func SlicePermutations[S ~[]E, E any](x S) Seq0 {
	v := anySlice[E](x)
	return IPermutations(v)
}
