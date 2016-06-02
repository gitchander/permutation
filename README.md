#permutation
Simple permutation package for golang

##Install
```bash
go get github.com/gitchander/permutation
```

##Examples

####permutations of int slice:
```go
package main

import (
	"fmt"

	"github.com/gitchander/permutation"
)

func main() {
	a := []int{1, 2, 3}
	p := permutation.New(permutation.IntSlice(a))
	for {
		fmt.Println(a)
		if !p.Next() {
			break
		}
	}
}
```
result:
```bash
[1 2 3]
[2 1 3]
[3 1 2]
[1 3 2]
[2 3 1]
[3 2 1]
```

####variants for trace permutations:
```go
package main

import (
	"fmt"

	"github.com/gitchander/permutation"
)

func main() {
	a := []string{"a", "b", "c"}
	fn := func() { fmt.Println(a) }

	trace1(permutation.New(permutation.StringSlice(a)), fn)
	fmt.Println()

	trace2(permutation.New(permutation.StringSlice(a)), fn)
	fmt.Println()

	trace3(permutation.New(permutation.StringSlice(a)), fn)
	fmt.Println()
}

func trace1(p *permutation.Permutation, fn func()) {
	fn()
	for p.Next() {
		fn()
	}
}

func trace2(p *permutation.Permutation, fn func()) {
	for {
		fn()
		if !p.Next() {
			break
		}
	}
}

func trace3(p *permutation.Permutation, fn func()) {
	for ok := true; ok; ok = p.Next() {
		fn()
	}
}
```
