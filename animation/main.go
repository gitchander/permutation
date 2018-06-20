package main

import (
	"bytes"
	"fmt"
	"image"
)

func main() {
	testGIF()
}

const elementSize = 9
const unitWidth = 4
const elementsDistatce = unitWidth * (elementSize + 5)

var mdig = func() map[int][][]bool {

	m := make(map[int][][]bool)
	for _, d := range digits {
		ssb, size := parseBitmap(d.bitmap, '0')

		if size.X != elementSize {
			panic("invalid element size x")
		}
		if size.Y != elementSize {
			panic("invalid element size y")
		}

		m[d.val] = ssb

		//fmt.Println(d.name, size)
		//printSSB(ssb)
	}
	return m
}()

type digit struct {
	name   string
	val    int
	bitmap []string
}

var digits = []digit{
	{
		name: "zero",
		val:  0,
		bitmap: []string{
			"---------",
			"---------",
			"---000---",
			"---0-0---",
			"---0-0---",
			"---0-0---",
			"---000---",
			"---------",
			"---------",
		},
	},
	{
		name: "one",
		val:  1,
		bitmap: []string{
			"---------",
			"---------",
			"----0----",
			"---00----",
			"----0----",
			"----0----",
			"---000---",
			"---------",
			"---------",
		},
	},
	{
		name: "two",
		val:  2,
		bitmap: []string{
			"---------",
			"---------",
			"---000---",
			"-----0---",
			"---000---",
			"---0-----",
			"---000---",
			"---------",
			"---------",
		},
	},
	{
		name: "three",
		val:  3,
		bitmap: []string{
			"---------",
			"---------",
			"---000---",
			"-----0---",
			"---000---",
			"-----0---",
			"---000---",
			"---------",
			"---------",
		},
	},
	{
		name: "four",
		val:  4,
		bitmap: []string{
			"---------",
			"---------",
			"---0-0---",
			"---0-0---",
			"---000---",
			"-----0---",
			"-----0---",
			"---------",
			"---------",
		},
	},
	{
		name: "five",
		val:  5,
		bitmap: []string{
			"---------",
			"---------",
			"---000---",
			"---0-----",
			"---000---",
			"-----0---",
			"---000---",
			"---------",
			"---------",
		},
	},
	{
		name: "six",
		val:  6,
		bitmap: []string{
			"---------",
			"---------",
			"---000---",
			"---0-----",
			"---000---",
			"---0-0---",
			"---000---",
			"---------",
			"---------",
		},
	},
	{
		name: "seven",
		val:  7,
		bitmap: []string{
			"---------",
			"---------",
			"---000---",
			"-----0---",
			"-----0---",
			"-----0---",
			"-----0---",
			"---------",
			"---------",
		},
	},
	{
		name: "eight",
		val:  8,
		bitmap: []string{
			"---------",
			"---------",
			"---000---",
			"---0-0---",
			"---000---",
			"---0-0---",
			"---000---",
			"---------",
			"---------",
		},
	},
	{
		name: "nine",
		val:  9,
		bitmap: []string{
			"---------",
			"---------",
			"---000---",
			"---0-0---",
			"---000---",
			"-----0---",
			"---000---",
			"---------",
			"---------",
		},
	},
}

func parseBitmap(bitmap []string, target rune) (ssb [][]bool, size image.Point) {

	var ps []image.Point

	size.Y = len(bitmap)
	for y, s := range bitmap {
		x := 0
		for _, r := range s {
			if r == target {
				p := image.Point{X: x, Y: y}
				ps = append(ps, p)
			}
			x++
			size.X = maxInt(size.X, x)
		}
	}

	ssb = make([][]bool, size.Y)
	for y := range ssb {
		ssb[y] = make([]bool, size.X)
	}

	for _, p := range ps {
		ssb[p.Y][p.X] = true
	}

	return ssb, size
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func printSSB(ssb [][]bool) {
	var buf bytes.Buffer
	for _, sb := range ssb {
		for _, b := range sb {
			if b {
				buf.WriteByte('0')
			} else {
				buf.WriteByte('-')
			}
		}
		buf.WriteByte('\n')
	}
	fmt.Println(buf.String())
}
