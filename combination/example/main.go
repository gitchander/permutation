package main

import (
	"fmt"

	comb "github.com/gitchander/permutation/combination"
)

func main() {
	exampleCombinations()
}

func exampleCombinations() {

	fmt.Println("Combinations:")

	var (
		set = []string{"A", "B", "C", "D", "E"}
		n   = len(set) // set len

		subset = make([]string, 3)
		k      = len(subset) // subset len
	)

	c := comb.NewComb(n, k)
	for c.Next() {

		subset = subset[:0] // reset subset.
		c.WalkSubset(func(index int) {
			subset = append(subset, set[index])
		})

		fmt.Println(subset)
	}
}

func makeSubset(set []string, c *comb.Combinator) []string {
	subset := make([]string, 0, len(set))
	c.WalkSubset(func(index int) {
		subset = append(subset, set[index])
	})
	return subset
}
