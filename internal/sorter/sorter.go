package sorter

import (
	"math/rand"
	"sort"
	"sync"
)

// CreateRandomInts returns a slice of the specified length that consists of
// random positive int values.
func CreateRandomInts(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = int(rand.Int())
	}
	return result
}

// SortFunctions is a slice of all sort functions implemented in this package.
var SortFunctions = []func([]int){
	BubbleSort,
	sort.Ints,
	QuickSort,
	GoroutineSort,
}

// BubbleSort sorts the specified list using the bubblesort algorithm.
func BubbleSort(slice []int) {
	for i := 0; i < len(slice)-1; i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

// QuickSort sorts the specified list using the quicksort algorithm.
func QuickSort(slice []int) {
	if len(slice) < 2 {
		return // already sorted
	}

	selectBestPivot(slice)
	pivotIndex := splitUsingPivot(slice)

	QuickSort(slice[:pivotIndex])
	QuickSort(slice[pivotIndex+1:])
}

// GoroutineSort sorts the specified list using the quicksort algorithm.
// This function uses goroutines for large lists.
func GoroutineSort(slice []int) {
	if len(slice) < 2 {
		return // already sorted
	}

	selectBestPivot(slice)
	pivotIndex := splitUsingPivot(slice)

	// Only use goroutines if we have a lot of entries.
	if len(slice) > 5000 {
		var wg sync.WaitGroup
		wg.Add(2)
		go parallelSort(slice[:pivotIndex], &wg)
		go parallelSort(slice[pivotIndex+1:], &wg)
		wg.Wait()
	} else {
		QuickSort(slice[:pivotIndex])
		QuickSort(slice[pivotIndex+1:])
	}
}

// selectBestPivot inspects the specified slice and makes sure that the first
// element is a suitable pivot. This implementation uses the median of the
// first, middle and last elements.
func selectBestPivot(slice []int) {
	firstIndex := 0
	middleIndex := (len(slice) - 1) / 2
	lastIndex := len(slice) - 1

	if slice[firstIndex] <= slice[middleIndex] && slice[middleIndex] <= slice[lastIndex] {
		slice[firstIndex], slice[middleIndex] = slice[middleIndex], slice[firstIndex]
	} else if slice[middleIndex] <= slice[lastIndex] && slice[lastIndex] <= slice[firstIndex] {
		slice[firstIndex], slice[lastIndex] = slice[lastIndex], slice[firstIndex]
	}
}

// splitUsingPivot takes the first element of the specified slice as a pivot
// element. Then it sorts all other elements into two groups, those that are
// bigger and those that are smaller/equal. It then arranges in the slice first
// the smaller/equal elements, then the pivot element and finally the bigger
// elements. The function returns the index of the pivot.
func splitUsingPivot(slice []int) int {
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
	return pivotIndex
}

// parallelSort is an internal helper for calling the goroutine function.
func parallelSort(slice []int, wg *sync.WaitGroup) {
	defer wg.Done()
	GoroutineSort(slice)
}
