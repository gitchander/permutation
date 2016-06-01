package permutation

import "errors"

var ErrTraceBreak = errors.New("permutation: trace break")

func Trace(v interface{}, fn func(v interface{}) bool) error {

	p, err := New(v)
	if err != nil {
		return err
	}

	if !fn(v) {
		return ErrTraceBreak
	}
	for p.Next() {
		if !fn(v) {
			return ErrTraceBreak
		}
	}

	return nil
}
