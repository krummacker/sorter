package gsorter

import (
	"math/rand"
	"sort"
	"sync"
	"time"
)

// IntSortable is a convenience wrapper for int slices that are to be sorted.
type IntSortable []int

func (a IntSortable) Len() int           { return len(a) }
func (a IntSortable) Less(i, j int) bool { return a[i] < a[j] }
func (a IntSortable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// StringSortable is a convenience wrapper for string slices that are to be sorted.
type StringSortable []string

func (a StringSortable) Len() int           { return len(a) }
func (a StringSortable) Less(i, j int) bool { return a[i] < a[j] }
func (a StringSortable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// TimeSortable is a convenience wrapper for time.Time slices that are to be sorted.
type TimeSortable []time.Time

func (a TimeSortable) Len() int           { return len(a) }
func (a TimeSortable) Less(i, j int) bool { return a[i].UnixNano() < a[j].UnixNano() }
func (a TimeSortable) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortFunctions is a slice of all sort functions implemented in this package.
var SortFunctions = []func(sort.Interface){
	sort.Sort,
	BubbleSort,
	QuickSort,
	GoroutineSort,
}

// BubbleSort sorts the specified data using the bubblesort algorithm.
func BubbleSort(data sort.Interface) {
	length := data.Len()
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j+1, j)
			}
		}
	}
}

// QuickSort sorts the specified data using the quicksort algorithm.
func QuickSort(data sort.Interface) {
	quickSortRange(data, 0, data.Len()-1)
}

// quickSortRange is an internal function for recursive calls. It sorts only
// those parts of the data specified by the 'from' and 'to' indexes.
func quickSortRange(data sort.Interface, from int, to int) {
	if to-from < 1 {
		return // already sorted
	}
	selectBestPivot(data, from, to)
	pivotIndex := splitUsingPivot(data, from, to)
	quickSortRange(data, from, pivotIndex-1)
	quickSortRange(data, pivotIndex+1, to)
}

// selectBestPivot inspects the specified data and makes sure that the first
// element is a suitable pivot. This implementation uses the median of the
// first, middle and last elements.
func selectBestPivot(data sort.Interface, from int, to int) {
	firstIndex := from
	middleIndex := from + (to-from)/2
	lastIndex := to
	if data.Less(firstIndex, middleIndex) && data.Less(middleIndex, lastIndex) {
		data.Swap(firstIndex, middleIndex)
	} else if data.Less(middleIndex, lastIndex) && data.Less(lastIndex, firstIndex) {
		data.Swap(firstIndex, lastIndex)
	}
}

// splitUsingPivot takes the first element of the specified data as a pivot
// element. Then it sorts all other elements into two groups, those that are
// bigger and those that are smaller/equal. It then arranges in the data first
// the smaller/equal elements, then the pivot element and finally the bigger
// elements. The function returns the index of the pivot.
func splitUsingPivot(data sort.Interface, from int, to int) int {
	left, right := from+1, to
	for left < right {
		if data.Less(from, left) {
			data.Swap(left, right)
			right -= 1
		} else {
			left += 1
		}
	}
	var pivotIndex int
	if data.Less(left, from) {
		data.Swap(from, left)
		pivotIndex = left
	} else {
		data.Swap(from, left-1)
		pivotIndex = left - 1
	}
	return pivotIndex
}

// GoroutineSort sorts the specified data using the quicksort algorithm.
// This function uses goroutines for large lists.
func GoroutineSort(data sort.Interface) {
	goroutineSortRange(data, 0, data.Len()-1)
}

// goroutineSortRange is an internal function for recursive calls. It sorts
// only those parts of the data specified by the 'from' and 'to' indexes.
func goroutineSortRange(data sort.Interface, from int, to int) {
	if to-from < 1 {
		return // already sorted
	}
	selectBestPivot(data, from, to)
	pivotIndex := splitUsingPivot(data, from, to)

	// Only use goroutines if we have a lot of entries.
	if to-from > 5000 {
		var wg sync.WaitGroup
		wg.Add(2)
		go parallelSort(data, from, pivotIndex-1, &wg)
		go parallelSort(data, pivotIndex+1, to, &wg)
		wg.Wait()
	} else {
		quickSortRange(data, from, pivotIndex-1)
		quickSortRange(data, pivotIndex+1, to)
	}
}

// parallelSort is an internal helper for calling the goroutine function.
func parallelSort(data sort.Interface, from int, to int, wg *sync.WaitGroup) {
	defer wg.Done()
	goroutineSortRange(data, from, to)
}

// CreateRandomInts returns a slice of the specified size that consists of
// random positive int values.
func CreateRandomInts(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = int(rand.Int())
	}
	return result
}

// CreateRandomStrings returns a slice of the specified size that consists of
// random strings of the specified length.
func CreateRandomStrings(size int, length int) []string {
	result := make([]string, size)
	for i := 0; i < size; i++ {
		c := byte('A' + rand.Intn(26))
		result[i] = string(c)
		for j := 1; j < length; j++ {
			c := byte('a' + rand.Intn(26))
			result[i] += string(c)
		}
	}
	return result
}

// CreateRandomTimes returns a slice of the specified size that consists of
// random date/times of the past 100 years.
func CreateRandomTimes(size int) []time.Time {
	result := make([]time.Time, size)
	for i := 0; i < size; i++ {
		randYears := rand.Intn(99)
		randMonths := rand.Intn(12)
		randDays := rand.Intn(31)
		result[i] = time.Now().AddDate(-randYears, -randMonths, -randDays)
	}
	return result
}
