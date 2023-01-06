package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"time"
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

func main() {
	sizes := []int{10, 50, 100, 500, 1000, 5000, 10000, 50000, 100000, 500000, 1000000}
	loops := 10

	rand.Seed(time.Now().UnixNano())

	fmt.Println()
	fmt.Println("Elements |    Bubble/u     Bubble/s   Standard/u   Standard/s      Quick/u      Quick/s  Goroutine/u  Goroutine/s ")
	fmt.Println("---------+--------------------------------------------------------------------------------------------------------")
	for _, size := range sizes {
		functionToDuration := make(map[string][]int)
		for i := 0; i < loops; i++ {
			original := CreateRandomInts(size)
			for _, sortFunction := range SortFunctions {
				name := runtime.FuncForPC(reflect.ValueOf(sortFunction).Pointer()).Name()

				// Bubble sort is too slow on large lists.
				if name == "main.BubbleSort" && size >= 10000 {
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
			Average(functionToDuration["main.BubbleSort.unsorted"]),
			Average(functionToDuration["main.BubbleSort.sorted"]),
			Average(functionToDuration["sort.Ints.unsorted"]),
			Average(functionToDuration["sort.Ints.sorted"]),
			Average(functionToDuration["main.QuickSort.unsorted"]),
			Average(functionToDuration["main.QuickSort.sorted"]),
			Average(functionToDuration["main.GoroutineSort.unsorted"]),
			Average(functionToDuration["main.GoroutineSort.sorted"]))
		fmt.Println()
	}
	fmt.Println()
}

// runSortFunction executes the specified sort function on the specified data
// and returns the microseconds used.
func runSortFunction(sortFunction func([]int), data []int) int {
	before := time.Now().UnixMicro()
	sortFunction(data)
	return int(time.Now().UnixMicro() - before)
}
