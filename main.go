package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"gitlab.nonprod.dwpcloud.uk/lewis.marsdenlambert/aoc3/slope"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var m slope.SlopeMap

	for x := 0; ; x++ {
		l, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		l = strings.TrimSpace(l)

		for y, mm := range l {
			m[x][y] = slope.MapMarker(mm)
		}

		if err != nil {
			break
		}
	}

	// Different paths
	paths := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	var multipliedTrees int

	for y, x := range paths {
		w := slope.NewWalker(&m)
		t, f := walkTheWalk(w, x[0], x[1])

		if y == 0 {
			multipliedTrees = t
		} else {
			multipliedTrees *= t
		}

		fmt.Printf("Right %d, down %d met %03d trees and %03d free spaces\n", x[0], x[1], t, f)
	}

	fmt.Printf("\nAll trees multiplied: %d\n", multipliedTrees)
}

func walkTheWalk(sw *slope.SlopeWalker, right, down int) (int, int) {
	var trees int
	var free int

	for {
		if err := sw.WalkWith(right, down); err != nil {
			break
		}

		switch w := sw.Current(); w {
		case slope.MapTree:
			trees += 1
		case slope.MapFree:
			free += 1
		}
	}

	return trees, free
}
