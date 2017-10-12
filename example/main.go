package main

import (
	"fmt"
	"log"

	"github.com/gitchander/permutation"
)

func main() {
	exampleIntSlice()
	exampleStringSlice()
	exampleAnySlice()
	exampleMustAnySlice()
	exampleEmptySlice()
	exampleInterface()
	exampleBadFactorial()
}

func exampleIntSlice() {
	a := []int{1, 2, 3}
	p := permutation.New(permutation.IntSlice(a))
	for p.Scan() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleStringSlice() {
	a := []string{"alpha", "beta", "gamma"}
	p := permutation.New(permutation.StringSlice(a))
	for p.Scan() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleAnySlice() {

	a := []interface{}{-1, "control", 9.3}

	data, err := permutation.NewAnySlice(a)
	if err != nil {
		log.Fatal(err)
	}

	p := permutation.New(data)
	for p.Scan() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleMustAnySlice() {
	a := []int{1, 2}
	p := permutation.New(permutation.MustAnySlice(a))
	for p.Scan() {
		fmt.Println(a)
	}
	fmt.Println()
}

func exampleEmptySlice() {
	a := []interface{}{}
	p := permutation.New(permutation.MustAnySlice(a))
	for p.Scan() {
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
	p := permutation.New(PersonSlice(a))
	for p.Scan() {
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
	p := permutation.New(emptiesSlice(n))
	for p.Scan() {
		val++
	}
	return
}

type emptiesSlice int

func (p emptiesSlice) Len() int    { return int(p) }
func (emptiesSlice) Swap(i, j int) {}
