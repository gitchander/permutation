package permutation

// Iterate all permutation for v.
// If f returns false, Walk stops the iteration.
func Walk(v interface{}, f func() bool) {

	i, ok := v.(Interface)
	if !ok {
		i = MustAnySlice(v)
	}

	p := New(i)

	iterate_V1(p, f)
	//	iterate_V2(p, f)
	//	iterate_V3(p, f)
}

func iterate_V1(p *Permutator, f func() bool) {
	if !f() {
		return
	}
	for p.Next() {
		if !f() {
			return
		}
	}
}

func iterate_V2(p *Permutator, f func() bool) {
	for {
		if !f() {
			return
		}
		if !p.Next() {
			break
		}
	}
}

func iterate_V3(p *Permutator, f func() bool) {
	for ok := true; ok; ok = p.Next() {
		if !f() {
			return
		}
	}
}
