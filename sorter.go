package main

import (
	"sort"
)

// A slice of all implemented sort functions.
var SortFunctions = []func([]int){
	BubbleSort,
	QuickSort,
	sort.Ints,
	GoroutineSort,
	InPlaceSort,
}

// ---------------------------------------------------------------------------------

// Sorts the specified list using the bubblesort algorithm.
func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// ---------------------------------------------------------------------------------

// Sorts the specified list using the quicksort algorithm.
func QuickSort(slice []int) {
	tmp := quicksort(slice)
	copy(slice, tmp)
}

// internal helper
func quicksort(slice []int) []int {
	if len(slice) < 2 {
		return slice // already sorted
	}
	pivot := slice[0]
	smaller := make([]int, 0, len(slice))
	bigger := make([]int, 0, len(slice))

	for _, element := range slice[1:] {
		if element < pivot {
			smaller = append(smaller, element)
		} else {
			bigger = append(bigger, element)
		}
	}

	result := make([]int, 0, len(slice))
	result = append(result, quicksort(smaller)...)
	result = append(result, pivot)
	result = append(result, quicksort(bigger)...)
	return result
}

// ---------------------------------------------------------------------------------

// Sorts and returns the specified list using the quicksort algorithm.
// Uses goroutines for large lists.
func GoroutineSort(slice []int) {
	tmp := quicksortGoroutine(slice)
	copy(slice, tmp)
}

// internal helper
func quicksortGoroutine(slice []int) []int {
	if len(slice) < 2 {
		return slice // already sorted
	}
	pivot := slice[0]
	smaller := make([]int, 0, len(slice))
	bigger := make([]int, 0, len(slice))

	for _, element := range slice[1:] {
		if element < pivot {
			smaller = append(smaller, element)
		} else {
			bigger = append(bigger, element)
		}
	}

	var first []int
	var last []int

	// Only use goroutines if we have a lot of entries.
	if len(slice) > 5000 {
		firstChannel := make(chan []int)
		lastChannel := make(chan []int)
		go channelSort(smaller, firstChannel)
		go channelSort(bigger, lastChannel)
		first = <-firstChannel
		last = <-lastChannel
	} else {
		first = quicksortGoroutine(smaller)
		last = quicksortGoroutine(bigger)
	}

	result := make([]int, 0, len(slice))
	result = append(result, first...)
	result = append(result, pivot)
	result = append(result, last...)
	return result
}

// internal helper to call goroutine
func channelSort(list []int, c chan []int) {
	c <- quicksortGoroutine(list)
}

// ---------------------------------------------------------------------------------

// Sorts and returns the specified list using the quick sort algorithm.
// Creates one copy of the list and sorts there in place.
func InPlaceSort(slice []int) {
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

	InPlaceSort(slice[:pivotIndex])
	InPlaceSort(slice[pivotIndex+1:])
}
