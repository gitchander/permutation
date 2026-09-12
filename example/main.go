package main

import (
	"fmt"

	prmt "github.com/gitchander/permutation"
)

func main() {
	exampleIntSlice()
	exampleStringSlice()
	exampleReflectElements()
	exampleArray()
	exampleEmptySlice()
	exampleInterface()
	exampleFactorialByPermutations()
	exampleRepeat()
	exampleIter()
}

func exampleIntSlice() {
	a := []int{1, 2, 3}
	for range prmt.SlicePermutations(a) {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleStringSlice() {
	a := []string{"alpha", "beta", "gamma"}
	p := prmt.NewSlicePermutator(a)
	p.WalkPermutations(func() bool {
		fmt.Println(a)
		return true
	})
	fmt.Println()
}

func exampleReflectElements() {
	a := []interface{}{-1, "control", 9.3}
	v, err := prmt.ReflectElements(a)
	if err != nil {
		panic(err)
	}
	for range prmt.IPermutations(v) {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleArray() {
	a := [3]int{1, 2, 3}
	for range prmt.SlicePermutations(a[:]) {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleEmptySlice() {
	var a []struct{}
	for range prmt.SlicePermutations(a) {
		fmt.Println(a)
	}
	fmt.Println()
}

//------------------------------------------------------------------------------

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
	for range prmt.IPermutations(PersonSlice(a)) {
		fmt.Println(a)
	}
	fmt.Println()
}

//------------------------------------------------------------------------------

func exampleFactorialByPermutations() {
	fmt.Println("Factorial by permutations:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorialByPermutations(i))
	}
	fmt.Println()
}

type onlyLen int

func (x onlyLen) Len() int    { return int(x) }
func (onlyLen) Swap(i, j int) {}

var _ prmt.Interface = onlyLen(0)

func factorialByPermutations(n int) int {
	if n < 0 {
		panic("negative factorial")
	}
	var count int
	for range prmt.IPermutations(onlyLen(n)) {
		count++
	}
	return count
}

//------------------------------------------------------------------------------

func exampleRepeat() {
	fmt.Println("Repeat permutations:")
	fmt.Println()
	a := []int{1, 2, 3}
	p := prmt.NewSlicePermutator(a)

	for i := 0; i < 3; i++ {
		fmt.Printf("loop %d:\n", i+1)
		for range p.Permutations() {
			fmt.Println(a)
		}
		fmt.Println()
	}
}

//------------------------------------------------------------------------------

func exampleIter() {
	fmt.Println("Iter permutations:")
	fmt.Println()

	a := []int{1, 2, 3}
	for range prmt.SlicePermutations(a) {
		fmt.Println(a)
	}
}
