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

func New(v Interface) *Permutator {
	return &Permutator{
		v: v,
		b: make([]int, v.Len()),
	}
}

func (p *Permutator) Next() bool {
	n := flipNumber(p.b)
	if n <= len(p.b) {
		flip(p.v, n)
		return true
	}
	flip(p.v, len(p.b)) // for return to begin state
	return false
}

func flipNumber(b []int) int {
	for i := range b {
		b[i]++
		if b[i] < i+2 {
			return i + 2
		}
		b[i] = 0
	}
	return len(b) + 1
}

func flip(v Interface, n int) {
	i, j := 0, n-1
	for i < j {
		v.Swap(i, j)
		i, j = i+1, j-1
	}
}
