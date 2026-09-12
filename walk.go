package permutation

// Callback function type for walking permutations
type WalkFunc func() (keepWalking bool)

func walkPermutations_V1(p *Permutator, wf WalkFunc) {
	if !wf() {
		return
	}
	for p.next() {
		if !wf() {
			return
		}
	}
}

func walkPermutations_V2(p *Permutator, wf WalkFunc) {
	for {
		if !wf() {
			return
		}
		if !p.next() {
			break
		}
	}
}

func walkPermutations_V3(p *Permutator, wf WalkFunc) {
	for ok := true; ok; ok = p.next() {
		if !wf() {
			return
		}
	}
}

func walkPermutations(p *Permutator, wf WalkFunc) {
	//walkPermutations_V1(p, wf)
	//walkPermutations_V2(p, wf)
	walkPermutations_V3(p, wf)
}
