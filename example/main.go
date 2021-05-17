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
	exampleArray()
	exampleEmptySlice()
	exampleInterface()
	exampleFactorialByPermutations()
	exampleRepeat()
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

func exampleArray() {
	a := [3]int{1, 2, 3}
	p := prmt.New(prmt.MustAnySlice(&a))
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
	p := prmt.New(emptiesSlice(n))
	for p.Next() {
		count++
	}
	return count
}

// slice of empties
type emptiesSlice int

func (p emptiesSlice) Len() int    { return int(p) }
func (emptiesSlice) Swap(i, j int) {}

func exampleRepeat() {

	fmt.Println("Repeat permutations:")
	fmt.Println()

	a := []int{1, 2, 3}

	p := prmt.New(prmt.IntSlice(a))

	fmt.Println("loop 1:")
	for p.Next() {
		fmt.Println(a)
	}
	fmt.Println()

	fmt.Println("loop 2:")
	for p.Next() {
		fmt.Println(a)
	}
	fmt.Println()
}
