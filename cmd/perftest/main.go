package main

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"time"

	"gitlab.com/dirk.krummacker/sorter/internal/gsorter"
)

// Average returns the average of the specified ints or 0 if there are no
// elements.
func Average(input []int) int {
	if len(input) == 0 {
		return 0
	}
	sum := 0
	for _, element := range input {
		sum += int(element)
	}
	return sum / len(input)
}

// Usage example: go run cmd/perfcheck/perfcheck.go
func main() {
	sizes := []int{10, 50, 100, 500, 1000, 5000, 10000, 50000, 100000, 500000, 1000000}
	loops := 10

	fmt.Println()
	fmt.Println("Elements |    Bubble/u     Bubble/s      Quick/u      Quick/s  Goroutine/u  Goroutine/s   Standard/u   Standard/s ")
	fmt.Println("---------+--------------------------------------------------------------------------------------------------------")
	for _, size := range sizes {
		functionToDuration := make(map[string][]int)
		for i := 0; i < loops; i++ {
			original := gsorter.CreateRandomInts(size)
			for _, sortFunction := range gsorter.SortFunctions {
				name := runtime.FuncForPC(reflect.ValueOf(sortFunction).Pointer()).Name()

				// Bubble sort is too slow on large lists.
				if name == "gitlab.com/dirk.krummacker/sorter/internal/gsorter.BubbleSort" && size >= 10000 {
					continue
				}

				data := make([]int, len(original))
				copy(data, original)

				unsortedName := name + ".unsorted"
				unsortedDuration := runSortFunction(sortFunction, data)
				functionToDuration[unsortedName] = append(functionToDuration[unsortedName],
					unsortedDuration)

				// Sort again the same data to discover if the sort algorithm
				// can cope with that.
				sortedName := name + ".sorted"
				sortedDuration := runSortFunction(sortFunction, data)
				functionToDuration[sortedName] = append(functionToDuration[sortedName],
					sortedDuration)
			}
		}

		fmt.Printf("%8d |  %10d   %10d   %10d   %10d   %10d   %10d   %10d   %10d",
			size,
			Average(functionToDuration["gitlab.com/dirk.krummacker/sorter/internal/gsorter.BubbleSort.unsorted"]),
			Average(functionToDuration["gitlab.com/dirk.krummacker/sorter/internal/gsorter.BubbleSort.sorted"]),
			Average(functionToDuration["gitlab.com/dirk.krummacker/sorter/internal/gsorter.QuickSort.unsorted"]),
			Average(functionToDuration["gitlab.com/dirk.krummacker/sorter/internal/gsorter.QuickSort.sorted"]),
			Average(functionToDuration["gitlab.com/dirk.krummacker/sorter/internal/gsorter.GoroutineSort.unsorted"]),
			Average(functionToDuration["gitlab.com/dirk.krummacker/sorter/internal/gsorter.GoroutineSort.sorted"]),
			Average(functionToDuration["sort.Sort.unsorted"]),
			Average(functionToDuration["sort.Sort.sorted"]))
		fmt.Println()
	}
	fmt.Println()
}

// runSortFunction executes the specified sort function on the specified data
// and returns the microseconds used.
func runSortFunction(sortFunction func(sort.Interface), data []int) int {
	before := time.Now().UnixMicro()
	sortFunction(gsorter.IntSortable(data))
	return int(time.Now().UnixMicro() - before)
}
