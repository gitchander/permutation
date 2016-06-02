package permutation

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type Permutation struct {
	data Interface
	b    []int
}

func New(data Interface) *Permutation {
	return &Permutation{
		data: data,
		b:    make([]int, data.Len()),
	}
}

func (p *Permutation) Next() bool {
	if p.data != nil {
		if next(p.data, p.b) {
			return true
		}
		p.data = nil
		p.b = nil
	}
	return false
}

func next(data Interface, b []int) bool {
	for i := range b {
		b[i]++
		if b[i] < i+2 {
			if i < len(b)-1 {
				flip(data, i+2)
				return true
			}
			flip(data, len(b)) // for return to begin state
			return false
		}
		b[i] = 0
	}
	return false
}

func flip(data Interface, n int) {
	i, j := 0, n-1
	for i < j {
		data.Swap(i, j)
		i, j = i+1, j-1
	}
}
