package main

import (
	"fmt"

	prmt "github.com/gitchander/permutation"
	comb "github.com/gitchander/permutation/combination"
)

func main() {
	exampleCombinations()
	examplePermutationOfCombinations()
}

func exampleCombinations() {

	fmt.Println("Combinations:")

	var (
		set    = []string{"A", "B", "C", "D", "E"}
		subset = make([]string, 3)
	)

	var (
		n = len(set) // length of set
		k = len(subset)
	)

	c := comb.New(n, k)

	for c.Next() {

		// fill subset by indexes
		for subsetIndex, setIndex := range c.Indexes() {
			subset[subsetIndex] = set[setIndex]
		}

		fmt.Println(subset)
	}
	fmt.Println()
}

func examplePermutationOfCombinations() {

	fmt.Println("Permutation of combinations:")

	var (
		set    = []string{"A", "B", "C", "D", "E"}
		subset = make([]string, 3)
	)

	var (
		n = len(set)    // length of set
		k = len(subset) // length of subset
	)

	c := comb.New(n, k)
	p := prmt.New(prmt.StringSlice(subset))

	for c.Next() {

		// fill subset by indexes
		for subsetIndex, setIndex := range c.Indexes() {
			subset[subsetIndex] = set[setIndex]
		}

		for p.Next() {
			fmt.Println(subset)
		}
	}
}

func fillSubsetV1(set, subset []string, c *comb.Combinator) {
	c.RangeIndexes(
		func(subsetIndex, setIndex int) bool {
			subset[subsetIndex] = set[setIndex]
			return true
		})
}

func fillSubsetV2(set, subset []string, c *comb.Combinator) {
	for subsetIndex, setIndex := range c.Indexes() {
		subset[subsetIndex] = set[setIndex]
	}
}
