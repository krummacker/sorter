package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func createRandomInts(size int) []int {
	result := make([]int, size)
	for i := 0; i < size; i++ {
		result[i] = int(rand.Int())
	}
	return result
}

// Returns the average of the specified ints or 0 if there are no elements.
func average(input []int) int {
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
	sizes := []int{10, 50, 100, 500, 1000, 5000, 7500, 10000, 50000, 100000, 500000}
	var loops int = 10

	rand.Seed(time.Now().UnixNano())

	fmt.Println()
	fmt.Println("Elements |      Bubble |       Quick |    Standard |   Goroutine |    In Place ")
	fmt.Println("---------+-------------+-------------+-------------+-------------+-------------")
	for _, size := range sizes {
		sorterToDuration := make(map[string][]int)
		for i := 0; i < loops; i++ {
			original := createRandomInts(size)
			for _, sorter := range Sorters {
				name := reflect.TypeOf(sorter).Name()

				// Bubble sort is too slow on large lists.
				if name == "BubbleSorter" && size >= 10000 {
					continue
				}

				data := make([]int, len(original))
				copy(data, original)
				before := time.Now().UnixMicro()
				sorter.Sort(data)
				duration := time.Now().UnixMicro() - before
				sorterToDuration[name] = append(sorterToDuration[name], int(duration))
			}
		}

		fmt.Printf("%8d |  %10d |  %10d |  %10d |  %10d |  %10d",
			size,
			average(sorterToDuration["BubbleSorter"]),
			average(sorterToDuration["QuickSorter"]),
			average(sorterToDuration["StandardSorter"]),
			average(sorterToDuration["GoroutineSorter"]),
			average(sorterToDuration["InPlaceSorter"]))
		fmt.Println()
	}
	fmt.Println()
}
