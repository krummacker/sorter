package sorter

import (
	"reflect"
	"sort"
	"testing"
)

// TestSort tests all sort functions. Test data is provided in a map.
func TestSort(t *testing.T) {
	for _, sortFunction := range SortFunctions {

		tests := map[string]struct {
			slice []int
			want  []int
		}{
			"empty_input": {
				slice: []int{},
				want:  []int{},
			},
			"one_element": {
				slice: []int{42},
				want:  []int{42},
			},
			"three_elements_sorted": {
				slice: []int{1, 2, 3},
				want:  []int{1, 2, 3},
			},
			"three_elements_not_sorted": {
				slice: []int{3, 1, 2},
				want:  []int{1, 2, 3},
			},
			"all_elements_positive": {
				slice: []int{4, 7, 4, 2, 8, 9, 6},
				want:  []int{2, 4, 4, 6, 7, 8, 9},
			},
			"all_elements_positive_negative": {
				slice: []int{4, 7, -4, 2, -8, 9, 6},
				want:  []int{-8, -4, 2, 4, 6, 7, 9},
			},
		}

		for name, test := range tests {
			sortFunction(test.slice)
			if !reflect.DeepEqual(test.slice, test.want) {
				t.Errorf("%s: got %v but want %v", name, test.slice, test.want)
			}
		}
	}
}

// TestLargeSlice tests all sort functions with a large unsorted slice.
func TestLargeSlice(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := CreateRandomInts(1000)
		want := make([]int, 1000)
		copy(want, slice)
		sort.Ints(want)
		sortFunction(slice)
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v", "large_slice", slice, want)
		}
	}
}
