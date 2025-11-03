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
}

func exampleIntSlice() {
	a := []int{1, 2, 3}
	p := prmt.NewPermutatorForSlice(a)
	for ok := true; ok; ok = p.NextPermutation() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleStringSlice() {
	a := []string{"alpha", "beta", "gamma"}
	p := prmt.NewPermutatorForSlice(a)
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
	p := prmt.NewPermutator(v)
	for {
		fmt.Println(a)
		if !(p.NextPermutation()) {
			break
		}
	}
	fmt.Println()
}

func exampleArray() {
	a := [3]int{1, 2, 3}
	p := prmt.NewPermutatorForSlice(a[:])
	for ok := true; ok; ok = p.NextPermutation() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleEmptySlice() {
	var a []struct{}
	p := prmt.NewPermutatorForSlice(a)
	for ok := true; ok; ok = p.NextPermutation() {
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
	p := prmt.NewPermutator(PersonSlice(a))
	for ok := true; ok; ok = p.NextPermutation() {
		fmt.Println(a)
	}
	fmt.Println()
}

//------------------------------------------------------------------------------

func exampleFactorialByPermutations() {
	fmt.Println("Factorial by permutations:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
	fmt.Println()
}

func factorial(n int) int {
	if n < 0 {
		panic("negative factorial")
	}
	var count int
	p := prmt.NewPermutator(prmt.OnlyLen(n))
	for ok := true; ok; ok = p.NextPermutation() {
		count++
	}
	return count
}

//------------------------------------------------------------------------------

func exampleRepeat() {
	fmt.Println("Repeat permutations:")
	fmt.Println()
	a := []int{1, 2, 3}
	p := prmt.NewPermutatorForSlice(a)
	for i := 0; i < 3; i++ {
		fmt.Printf("loop %d:\n", i+1)
		for ok := true; ok; ok = p.NextPermutation() {
			fmt.Println(a)
		}
		fmt.Println()
	}
}
