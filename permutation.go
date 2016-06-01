package permutation

import (
	"errors"
	"reflect"
)

type Permutation struct {
	v reflect.Value
	b []int
}

func New(v interface{}) (*Permutation, error) {

	if v == nil {
		return nil, errors.New("permutation: argument is nil")
	}

	rv := reflect.ValueOf(v)

	if t := rv.Type(); t.Kind() != reflect.Slice {
		return nil, errors.New("permutation: argument must be a slice")
	}

	return &Permutation{
		v: rv,
		b: make([]int, rv.Len()),
	}, nil
}

func (p *Permutation) Next() bool {
	b := p.b
	for i := range b {
		b[i]++
		if b[i] < i+2 {
			if i < len(b)-1 {
				flip(p.v, i+2)
				return true
			}
			flip(p.v, len(b)) // for return to begin state
			p.v = reflect.Value{}
			p.b = nil
			break
		}
		b[i] = 0
	}
	return false
}

func flip(v reflect.Value, n int) {

	i, j := 0, n-1
	for i < j {
		swap(v.Index(i), v.Index(j))
		i, j = i+1, j-1
	}
}

func swap(v1, v2 reflect.Value) {

	var (
		i1 = v1.Interface()
		i2 = v2.Interface()
	)

	v1.Set(reflect.ValueOf(i2))
	v2.Set(reflect.ValueOf(i1))
}
