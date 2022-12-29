package main

import (
	"reflect"
	"runtime"
	"sort"
	"testing"
)

// Tests all sort functions with an empty slice.
func TestBubbleSortEmpty(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := []int{}
		want := []int{}
		sortFunction(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				runtime.FuncForPC(reflect.ValueOf(sortFunction).Pointer()).Name(),
				slice, want)
		}
	}
}

// Tests all sort functions with only one element.
func TestSortOneElement(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := []int{42}
		want := []int{42}
		sortFunction(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				runtime.FuncForPC(reflect.ValueOf(sortFunction).Pointer()).Name(),
				slice, want)
		}
	}
}

// Tests all sort functions with two elements that are already sorted.
func TestTwoElementsSorted(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := []int{1, 2}
		want := []int{1, 2}
		sortFunction(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				runtime.FuncForPC(reflect.ValueOf(sortFunction).Pointer()).Name(),
				slice, want)
		}
	}
}

// Tests all sort functions with two elements that are not sorted.
func TestTwoElementsNotSorted(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := []int{2, 1}
		want := []int{1, 2}
		sortFunction(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				runtime.FuncForPC(reflect.ValueOf(sortFunction).Pointer()).Name(),
				slice, want)
		}
	}
}

// Tests all sort functions with a list of positive numbers.
func TestSortAllPositive(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := []int{4, 7, 4, 2, 8, 9, 6}
		want := []int{2, 4, 4, 6, 7, 8, 9}
		sortFunction(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				runtime.FuncForPC(reflect.ValueOf(sortFunction).Pointer()).Name(),
				slice, want)
		}
	}
}

// Tests all sort functions with positive and negative numbers.
func TestSortPositiveNegative(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := []int{4, 7, -4, 2, -8, 9, 6}
		want := []int{-8, -4, 2, 4, 6, 7, 9}
		sortFunction(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				runtime.FuncForPC(reflect.ValueOf(sortFunction).Pointer()).Name(),
				slice, want)
		}
	}
}

// Tests all sort functions with a large unsorted slice.
func TestLargeSlice(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := createRandomInts(1000)
		want := make([]int, 1000)
		copy(want, slice)
		sort.Ints(want)
		sortFunction(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				runtime.FuncForPC(reflect.ValueOf(sortFunction).Pointer()).Name(),
				slice, want)
		}
	}
}
