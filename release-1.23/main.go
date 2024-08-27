package main

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
)

func main() {
	// slice all example
	a := []int{1, 2, 3, 4, 5}
	for i, v := range slices.All(a) {
		fmt.Printf("slice all result: %d: %d\n", i, v)
	}

	// slice values note: this is only return the values, not the index
	for v := range slices.Values(a) {
		fmt.Printf("slices values result: %d\n", v)
	}

	// slice backwards note: this is return the values in reverse order
	for _, v := range slices.Backward(a) {
		fmt.Printf("slices backward result: %d\n", v)
	}

	// collect values from iterator into new slices
	sq := func(yield func(int) bool) {
		for i := 0; i < 10; i += 2 {
			if !yield(i) {
				return
			}
		}
	}
	sqs := slices.Collect(sq)
	fmt.Println("slice collect result:", sqs)

	// append values from iterator into new slice then sort the slice
	seq := func(yield func(int) bool) {
		for i := 0; i < 10; i++ {
			if !yield(i) {
				return
			}
		}
	}
	seqs := slices.AppendSeq([]int{1, 2}, seq)
	fmt.Println("appendSeq: ", seqs)

	// Sorted collects values from seq into a new slice, sorts the slice, and returns it.
	sorted := slices.Sorted(seq)
	fmt.Println("sorted: ", sorted)
	fmt.Print("result slice sorted", slices.IsSorted(sorted), "\n")

	// SortedFunc collects values from seq into a new slice, sorts the slice using the comparison function, and returns it.
	seq2 := func(yield func(int) bool) {
		flag := -1
		for i := 0; i < 10; i += 2 {
			flag = -flag
			if !yield(i * flag) {
				return
			}
		}
	}

	sortFunc := func(a, b int) int {
		return cmp.Compare(b, a) // the comparison is being done in reverse
	}

	s := slices.SortedFunc(seq2, sortFunc)
	fmt.Println("result custom sortfunc in slices sorted func:", s)

	// SortedStableFunc collects values from seq into a new slice. It then sorts the slice while keeping the original order of equal elements, using the comparison function to compare elements. It returns the new slice.

	type MathScore struct {
		Name  string
		Score int
	}
	scores := []MathScore{
		{
			"John",
			100,
		},
		{
			"Jane",
			90,
		},
		{
			"James",
			80,
		},
		{
			"Eila",
			70,
		},
	}

	sortScoreFunc := func(a, b MathScore) int {
		return cmp.Compare(a.Score, b.Score)
	}

	sortedScores := slices.SortedStableFunc(slices.Values(scores), sortScoreFunc)
	fmt.Println("sorted scores:", sortedScores)

	// Chunk returns an iterator over consecutive sub-slices of up to n elements of s. All but the last sub-slice will have size n. All sub-slices are clipped to have no capacity beyond the length. If s is empty, the sequence is empty: there is no empty slice in the sequence. Chunk panics if n is less than 1.
	for chunkScore := range slices.Chunk(scores, 2) {
		fmt.Println("chunk score:", chunkScore)
	}

	// maps section, there is a several functions that added to maps package in go 1.23

	// All returns an iterator over key-value pairs from m. The iteration order is not specified and is not guaranteed to be the same from one call to the next.
	m1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	m2 := map[string]int{
		"one":   10,
		"two":   21,
		"three": 31,
	}
	// Insert adds the key-value pairs from seq to m. If a key in seq already exists in m, its value will be overwritten.

	maps.Insert(m2, maps.All(m1))
	fmt.Println("m2:", m2)

	// Keys returns an iterator over the keys of m. The iteration order is not specified and is not guaranteed to be the same from one call to the next

	m3 := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	keys := slices.Sorted(maps.Keys(m3))
	fmt.Println("keys m3:", keys)

	// Values returns an iterator over the values of m. The iteration order is not specified and is not guaranteed to be the same from one call to the next
	fmt.Println("values m3:", slices.Sorted(maps.Values(m3)))

	// Collect collects key-value pairs from seq into a new map and returns it.
	s1 := []string{"zero", "one", "two", "three"}
	fmt.Println("map s1:", maps.Collect(slices.All(s1)))
}
