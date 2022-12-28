package main

import "sort"

type Sorter interface {
	Sort(list []int) []int
}

// all sorter implementations
var Sorters = []Sorter{
	BubbleSorter{},
	QuickSorter{},
	StandardSorter{},
	GoroutineSorter{},
	InPlaceSorter{},
}

// ---------------------------------------------------------------------------------

type BubbleSorter struct{}

// Sorts and returns the specified list using the bubble sort algorithm.
// Leaves the specified list unchanged.
func (b BubbleSorter) Sort(list []int) []int {
	result := make([]int, len(list))
	copy(result, list)
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-1-i; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// ---------------------------------------------------------------------------------

type QuickSorter struct{}

// Sorts and returns the specified list using the quick sort algorithm.
// Leaves the specified list unchanged.
func (q QuickSorter) Sort(list []int) []int {
	if len(list) < 2 {
		return list // already sorted
	}
	pivot := list[0]
	smaller := make([]int, 0, len(list))
	bigger := make([]int, 0, len(list))

	for _, element := range list[1:] {
		if element < pivot {
			smaller = append(smaller, element)
		} else {
			bigger = append(bigger, element)
		}
	}

	result := make([]int, 0, len(list))
	result = append(result, q.Sort(smaller)...)
	result = append(result, pivot)
	result = append(result, q.Sort(bigger)...)
	return result
}

// ---------------------------------------------------------------------------------

type StandardSorter struct{}

// Sorts and returns the specified list using the standard Go sort algorithm.
// Leaves the specified list unchanged.
func (g StandardSorter) Sort(list []int) []int {
	result := make([]int, len(list))
	copy(result, list)
	sort.Ints(result)
	return result
}

// ---------------------------------------------------------------------------------

type GoroutineSorter struct{}

// Sorts and returns the specified list using the quick sort algorithm.
// Uses goroutines for large lists. Leaves the specified list unchanged.
func (g GoroutineSorter) Sort(list []int) []int {
	if len(list) < 2 {
		return list // already sorted
	}
	pivot := list[0]
	smaller := make([]int, 0, len(list))
	bigger := make([]int, 0, len(list))

	for _, element := range list[1:] {
		if element < pivot {
			smaller = append(smaller, element)
		} else {
			bigger = append(bigger, element)
		}
	}

	var first []int
	var last []int

	// Only use goroutines if we have a lot of entries.
	if len(list) > 5000 {
		firstChannel := make(chan []int)
		lastChannel := make(chan []int)
		go g.channelSort(smaller, firstChannel)
		go g.channelSort(bigger, lastChannel)
		first = <-firstChannel
		last = <-lastChannel
	} else {
		first = g.Sort(smaller)
		last = g.Sort(bigger)
	}

	result := make([]int, 0, len(list))
	result = append(result, first...)
	result = append(result, pivot)
	result = append(result, last...)
	return result
}

// internal helper to call goroutine
func (g GoroutineSorter) channelSort(list []int, c chan []int) {
	c <- g.Sort(list)
}

// ---------------------------------------------------------------------------------

type InPlaceSorter struct{}

// Sorts and returns the specified list using the quick sort algorithm.
// Creates one copy of the list and sorts there in place.
// Leaves the specified list unchanged.
func (i InPlaceSorter) Sort(list []int) []int {
	result := make([]int, len(list))
	copy(result, list)
	sortInPlace(result)
	return result
}

func sortInPlace(slice []int) {
	if len(slice) < 2 {
		return // already sorted
	}

	left, right := 1, len(slice)-1
	for left < right {
		if slice[left] > slice[0] {
			slice[left], slice[right] = slice[right], slice[left]
			right -= 1
		} else {
			left += 1
		}
	}
	var pivotIndex int
	if slice[left] <= slice[0] {
		slice[left], slice[0] = slice[0], slice[left]
		pivotIndex = left
	} else {
		slice[left-1], slice[0] = slice[0], slice[left-1]
		pivotIndex = left - 1
	}

	sortInPlace(slice[:pivotIndex])
	sortInPlace(slice[pivotIndex+1:])
}
