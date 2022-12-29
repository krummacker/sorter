package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
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
	sizes := []int{10, 50, 100, 500, 1000, 5000, 10000, 50000, 100000, 500000, 1000000}
	loops := 10

	rand.Seed(time.Now().UnixNano())

	fmt.Println()
	fmt.Println("Elements |      Bubble |    Standard |       Quick |   Goroutine ")
	fmt.Println("---------+-------------+-------------+-------------+-------------")
	for _, size := range sizes {
		functionToDuration := make(map[string][]int)
		for i := 0; i < loops; i++ {
			original := createRandomInts(size)
			for _, sortFunction := range SortFunctions {
				name := runtime.FuncForPC(reflect.ValueOf(sortFunction).Pointer()).Name()

				// Bubble sort is too slow on large lists.
				if name == "main.BubbleSort" && size >= 10000 {
					continue
				}

				data := make([]int, len(original))
				copy(data, original)
				before := time.Now().UnixMicro()
				sortFunction(data)
				duration := time.Now().UnixMicro() - before
				functionToDuration[name] = append(functionToDuration[name], int(duration))
			}
		}

		fmt.Printf("%8d |  %10d |  %10d |  %10d |  %10d",
			size,
			average(functionToDuration["main.BubbleSort"]),
			average(functionToDuration["sort.Ints"]),
			average(functionToDuration["main.QuickSort"]),
			average(functionToDuration["main.GoroutineSort"]))
		fmt.Println()
	}
	fmt.Println()
}
