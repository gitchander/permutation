package main

import (
	"fmt"
	"log"

	prmt "github.com/gitchander/permutation"
)

func main() {
	exampleIntSlice()
	exampleStringSlice()
	exampleAnySlice()
	exampleMustAnySlice()
	exampleEmptySlice()
	exampleInterface()
	exampleBadFactorial()
	exampleRepeat()
	exampleCombinations()
}

func exampleIntSlice() {
	a := []int{1, 2, 3}
	p := prmt.New(prmt.IntSlice(a))
	for p.Next() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleStringSlice() {
	a := []string{"alpha", "beta", "gamma"}
	p := prmt.New(prmt.StringSlice(a))
	for p.Next() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleAnySlice() {

	a := []interface{}{-1, "control", 9.3}

	data, err := prmt.NewAnySlice(a)
	if err != nil {
		log.Fatal(err)
	}

	p := prmt.New(data)
	for p.Next() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleMustAnySlice() {
	a := []int{1, 2}
	p := prmt.New(prmt.MustAnySlice(a))
	for p.Next() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleEmptySlice() {
	var a []struct{}
	p := prmt.New(prmt.MustAnySlice(a))
	for p.Next() {
		fmt.Println(a)
	}
	fmt.Println()
}

type Person struct {
	Name string
	Age  int
}

type PersonSlice []Person

func (ps PersonSlice) Len() int      { return len(ps) }
func (ps PersonSlice) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }

func exampleInterface() {
	a := []Person{
		{Name: "one", Age: 1},
		{Name: "two", Age: 2},
		{Name: "three", Age: 3},
	}
	p := prmt.New(PersonSlice(a))
	for p.Next() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleBadFactorial() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
}

func factorial(n int) (val int) {
	if n < 0 {
		return -1
	}
	p := prmt.New(emptiesSlice(n))
	for p.Next() {
		val++
	}
	return
}

type emptiesSlice int

func (p emptiesSlice) Len() int    { return int(p) }
func (emptiesSlice) Swap(i, j int) {}

func exampleRepeat() {
	a := []int{1, 2, 3}

	p := prmt.New(prmt.IntSlice(a))

	for p.Next() {
		fmt.Println(a)
	}
	fmt.Println()

	for p.Next() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleCombinations() {

	var (
		vs = []string{"1", "2", "3", "4"}
		n  = len(vs)

		ds = make([]string, 3)
		k  = len(ds)
	)

	as := make([]int, k)
	for i := range as {
		as[i] = i
	}

	p := prmt.New(prmt.IntSlice(as))
	for {
		//fmt.Println(as)
		for p.Next() {
			for i, a := range as {
				ds[i] = vs[a]
			}
			fmt.Println(ds)
		}
		overflow := nextComb(as, n)
		if overflow {
			break
		}
	}
	fmt.Println()
}

func nextComb(as []int, n int) (overflow bool) {
	if (len(as) == 1) || nextComb(as[1:], n) {
		d := as[0] + 1
		if d > (n - len(as)) {
			return true
		}
		for i := range as {
			as[i] = d + i
		}
	}
	return false
}
