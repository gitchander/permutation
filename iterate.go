package permutation

func Walk(v interface{}, f func()) {

	i, ok := v.(Interface)
	if !ok {
		i = MustAnySlice(v)
	}

	p := New(i)

	//------------------------------------------
	// Variant: 1
	f()
	for p.Next() {
		f()
	}
	//------------------------------------------
	// Variant: 2
	//	for {
	//		f()
	//		if !p.Next() {
	//			break
	//		}
	//	}
	//------------------------------------------
	// Variant: 3
	//	for ok := true; ok; ok = p.Next() {
	//		f()
	//	}
	//------------------------------------------
}
