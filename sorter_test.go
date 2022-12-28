package main

import (
	"reflect"
	"testing"
)

// Calls all sorters with an empty slice,
// checking for a valid return value.
func TestBubbleSortEmpty(t *testing.T) {
	for _, sorter := range Sorters {
		slice := []int{}
		want := []int{}
		sorter.Sort(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), slice, want)
		}
	}
}

// Calls all sorters with only one element,
// checking for a valid return value.
func TestSortOneElement(t *testing.T) {
	for _, sorter := range Sorters {
		slice := []int{42}
		want := []int{42}
		sorter.Sort(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), slice, want)
		}
	}
}

// Calls all sorters with two elements that are already sorted,
// checking for a valid return value.
func TestTwoElementsSorted(t *testing.T) {
	for _, sorter := range Sorters {
		slice := []int{1, 2}
		want := []int{1, 2}
		sorter.Sort(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), slice, want)
		}
	}
}

// Calls all sorters with two elements that are not sorted,
// checking for a valid return value.
func TestTwoElementsNotSorted(t *testing.T) {
	for _, sorter := range Sorters {
		slice := []int{2, 1}
		want := []int{1, 2}
		sorter.Sort(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), slice, want)
		}
	}
}

// Calls all sorters with a list of positive numbers,
// checking for a valid return value.
func TestSortAllPositive(t *testing.T) {
	for _, sorter := range Sorters {
		slice := []int{4, 7, 4, 2, 8, 9, 6}
		want := []int{2, 4, 4, 6, 7, 8, 9}
		sorter.Sort(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), slice, want)
		}
	}
}

// Calls all sorters with positive and negative numbers,
// checking for a valid return value.
func TestSortPositiveNegative(t *testing.T) {
	for _, sorter := range Sorters {
		slice := []int{4, 7, -4, 2, -8, 9, 6}
		want := []int{-8, -4, 2, 4, 6, 7, 9}
		sorter.Sort(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), slice, want)
		}
	}
}
