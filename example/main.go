package main

import (
	"fmt"
	"log"

	"github.com/gitchander/permutation"
)

func main() {
	useVariants()
	exampleTrace()
	exampleFactorial()
}

func useVariants() {

	vs := []interface{}{true, -5, "str"}

	fmt.Println("use variant 1:")
	useExample1(vs)
	fmt.Println()

	fmt.Println("use variant 2:")
	useExample2(vs)
	fmt.Println()

	fmt.Println("use variant 3:")
	useExample3(vs)
	fmt.Println()
}

func useExample1(v interface{}) {

	data, err := permutation.NewAnySlice(v)
	if err != nil {
		log.Fatal(err)
	}
	p := permutation.New(data)

	fmt.Println(v)
	for p.Next() {
		fmt.Println(v)
	}
}

func useExample2(v interface{}) {

	p := permutation.New(permutation.MustAnySlice(v))

	for {
		fmt.Println(v)
		if !p.Next() {
			break
		}
	}
}

func useExample3(v interface{}) {

	p := permutation.New(permutation.MustAnySlice(v))

	for ok := true; ok; ok = p.Next() {
		fmt.Println(v)
	}
}

func exampleTrace() {

	vs := []interface{}{
		[]int{},
		[]bool{true, false},
		[]int{1, 2, 3},
		[]string{"one", "two", "three", "four"},
	}

	for _, v := range vs {
		useExample3(v)
		fmt.Println()
	}
}

func exampleFactorial() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
}

func factorial(n int) int {
	i := 0
	p := permutation.New(permutation.EmptySlice(n))
	for {
		i++
		if !p.Next() {
			break
		}
	}
	return i
}
