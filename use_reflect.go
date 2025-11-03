package permutation

import (
	"errors"
	"reflect"
)

type reflectAnySlice struct {
	rv reflect.Value
}

func (p *reflectAnySlice) Len() int {
	return p.rv.Len()
}

func (p *reflectAnySlice) Swap(i, j int) {

	// reflect.Value
	var (
		v1 = p.rv.Index(i)
		v2 = p.rv.Index(j)
	)

	// interface{}
	var (
		i1 = v1.Interface()
		i2 = v2.Interface()
	)

	v1.Set(reflect.ValueOf(i2))
	v2.Set(reflect.ValueOf(i1))
}

func ReflectElements(v interface{}) (Interface, error) {
	if v == nil {
		return nil, errors.New("permutation: argument is nil")
	}
	rv := reflect.ValueOf(v)
	t := rv.Type()
	if t.Kind() == reflect.Ptr {
		switch e := t.Elem(); e.Kind() {
		case reflect.Array:
			rv = rv.Elem()
			if !rv.CanSet() {
				return nil, errors.New("permutation: element must be can set")
			}
		case reflect.Slice:
			rv = rv.Elem()
		default:
			return nil, errors.New("permutation: pointer element must be an array or a slice")
		}
	} else if t.Kind() != reflect.Slice {
		return nil, errors.New("permutation: argument must be a slice")
	}
	return &reflectAnySlice{rv}, nil
}
