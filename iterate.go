package permutation

func Walk(v interface{}, f func()) {

	i, ok := v.(Interface)
	if !ok {
		i = MustAnySlice(v)
	}

	p := New(i)

	iterate_V1(p, f)
	//	iterate_V2(p, f)
	//	iterate_V3(p, f)
}

func iterate_V1(p *Permutator, f func()) {
	f()
	for p.Next() {
		f()
	}
}

func iterate_V2(p *Permutator, f func()) {
	for {
		f()
		if !p.Next() {
			break
		}
	}
}

func iterate_V3(p *Permutator, f func()) {
	for ok := true; ok; ok = p.Next() {
		f()
	}
}
