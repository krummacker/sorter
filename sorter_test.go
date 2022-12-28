package main

import (
	"reflect"
	"testing"
)

// Calls all sorters with an empty slice,
// checking for a valid return value.
func TestBubbleSortEmpty(t *testing.T) {
	for _, sorter := range Sorters {
		data := []int{}
		want := []int{}
		got := sorter.Sort(data)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), got, want)
		}
	}
}

// Calls all sorters with only one element,
// checking for a valid return value.
func TestSortOneElement(t *testing.T) {
	for _, sorter := range Sorters {
		data := []int{42}
		want := []int{42}
		got := sorter.Sort(data)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), got, want)
		}
	}
}

// Calls all sorters with two elements that are already sorted,
// checking for a valid return value.
func TestTwoElementsSorted(t *testing.T) {
	for _, sorter := range Sorters {
		original := []int{1, 2}
		data := make([]int, len(original))
		copy(data, original)
		want := []int{1, 2}
		got := sorter.Sort(data)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), got, want)
		}
		// make sure that parameter content unchanged
		if !reflect.DeepEqual(data, original) {
			t.Errorf("%s: parameter content has changed: %v",
				reflect.TypeOf(sorter).Name(), data)
		}
	}
}

// Calls all sorters with two elements that are not sorted,
// checking for a valid return value.
func TestTwoElementsNotSorted(t *testing.T) {
	for _, sorter := range Sorters {
		original := []int{2, 1}
		data := make([]int, len(original))
		copy(data, original)
		want := []int{1, 2}
		got := sorter.Sort(data)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), got, want)
		}
		// make sure that parameter content unchanged
		if !reflect.DeepEqual(data, original) {
			t.Errorf("%s: parameter content has changed: %v",
				reflect.TypeOf(sorter).Name(), data)
		}
	}
}

// Calls all sorters with a list of positive numbers,
// checking for a valid return value.
func TestSortAllPositive(t *testing.T) {
	for _, sorter := range Sorters {
		original := []int{4, 7, 4, 2, 8, 9, 6}
		data := make([]int, len(original))
		copy(data, original)
		want := []int{2, 4, 4, 6, 7, 8, 9}
		got := sorter.Sort(data)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), got, want)
		}
		// make sure that parameter content unchanged
		if !reflect.DeepEqual(data, original) {
			t.Errorf("%s: parameter content has changed: %v",
				reflect.TypeOf(sorter).Name(), data)
		}
	}
}

// Calls all sorters with positive and negative numbers,
// checking for a valid return value.
func TestSortPositiveNegative(t *testing.T) {
	for _, sorter := range Sorters {
		original := []int{4, 7, -4, 2, -8, 9, 6}
		data := make([]int, len(original))
		copy(data, original)
		want := []int{-8, -4, 2, 4, 6, 7, 9}
		got := sorter.Sort(data)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%s: got %v but want %v",
				reflect.TypeOf(sorter).Name(), got, want)
		}
		// make sure that parameter content unchanged
		if !reflect.DeepEqual(data, original) {
			t.Errorf("%s: parameter content has changed: %v",
				reflect.TypeOf(sorter).Name(), data)
		}
	}
}
