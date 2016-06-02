#permutation
Simple permutation package for golang

##Install
```bash
go get github.com/gitchander/permutation
```

##Usage

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
```
[1 2 3]
[2 1 3]
[3 1 2]
[1 3 2]
[2 3 1]
[3 2 1]
```

####permutation of slice with any elements:
```go
package main

import (
	"fmt"
	"log"

	"github.com/gitchander/permutation"
)

func main() {

	a := []interface{}{true, -5, "one"}

	data, err := permutation.NewAnySlice(a)
	if err != nil {
		log.Fatal(err)
	}

	p := permutation.New(data)

	for ok := true; ok; ok = p.Next() {
		fmt.Println(a)
	}
}
```
or shorter:
```go
package main

import (
	"fmt"

	"github.com/gitchander/permutation"
)

func main() {
	var (
		a = []interface{}{true, -5, "one"}
		p = permutation.New(permutation.MustAnySlice(a))
	)
	for ok := true; ok; ok = p.Next() {
		fmt.Println(a)
	}
}
```
result:
```
[true -5 one]
[-5 true one]
[one true -5]
[true one -5]
[-5 one true]
[one -5 true]
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
