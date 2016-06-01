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

	p, err := permutation.New(v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
	for p.Next() {
		fmt.Println(v)
	}
}

func useExample2(v interface{}) {

	p, err := permutation.New(v)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println(v)
		if !p.Next() {
			break
		}
	}
}

func useExample3(v interface{}) {

	p, err := permutation.New(v)
	if err != nil {
		log.Fatal(err)
	}

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
		traceValue(v)
		fmt.Println()
	}
}

func traceValue(v interface{}) {
	err := permutation.Trace(v,
		func(w interface{}) bool {
			fmt.Println(w)
			return true
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func exampleFactorial() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d! = %d\n", i, factorial(i))
	}
}

func factorial(n int) int {
	a := make([]struct{}, n)
	i := 0
	err := permutation.Trace(a,
		func(_ interface{}) bool {
			i++
			return true
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	return i
}
