package gsorter

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

// TestIntSort tests all sort functions with slices of ints.
// Test data is provided in a map.
func TestIntSort(t *testing.T) {
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
			"five_elements_sorted": {
				slice: []int{1, 2, 3, 4, 5},
				want:  []int{1, 2, 3, 4, 5},
			},
			"six_elements_sorted": {
				slice: []int{1, 2, 3, 4, 5, 6},
				want:  []int{1, 2, 3, 4, 5, 6},
			},
			"seven_elements_sorted": {
				slice: []int{1, 2, 3, 4, 5, 6, 7},
				want:  []int{1, 2, 3, 4, 5, 6, 7},
			},
			"three_elements_not_sorted": {
				slice: []int{3, 1, 2},
				want:  []int{1, 2, 3},
			},
			"all_elements_positive": {
				slice: []int{4, 7, 4, 2, 8, 9, 6},
				want:  []int{2, 4, 4, 6, 7, 8, 9},
			},
			"all_elements_positive_even": {
				slice: []int{4, 7, 4, 2, 1, 8, 9, 6},
				want:  []int{1, 2, 4, 4, 6, 7, 8, 9},
			},
			"all_elements_positive_negative": {
				slice: []int{4, 7, -4, 2, -8, 9, 6},
				want:  []int{-8, -4, 2, 4, 6, 7, 9},
			},
			"repeating_elements": {
				slice: []int{1, 1, 1, 8, 8, 8, 5, 5, 5},
				want:  []int{1, 1, 1, 5, 5, 5, 8, 8, 8},
			},
		}
		for name, test := range tests {
			sortFunction(IntSortable(test.slice))
			if !reflect.DeepEqual(test.slice, test.want) {
				t.Errorf("%s: got %v but want %v", name, test.slice, test.want)
			}
		}
	}
}

// TestStringSort tests all sort functions with slices of strings.
// Test data is provided in a map.
func TestStringSort(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		tests := map[string]struct {
			slice []string
			want  []string
		}{
			"empty_input": {
				slice: []string{},
				want:  []string{},
			},
			"one_element": {
				slice: []string{"foo"},
				want:  []string{"foo"},
			},
			"three_elements_sorted": {
				slice: []string{"a", "b", "c"},
				want:  []string{"a", "b", "c"},
			},
			"three_elements_not_sorted": {
				slice: []string{"c", "a", "b"},
				want:  []string{"a", "b", "c"},
			},
			"five_elements_sorted": {
				slice: []string{"Alice", "Bob", "Charlie", "Dick", "Ernie"},
				want:  []string{"Alice", "Bob", "Charlie", "Dick", "Ernie"},
			},
			"five_elements_not_sorted": {
				slice: []string{"Bob", "Alice", "Ernie", "Dick", "Charlie"},
				want:  []string{"Alice", "Bob", "Charlie", "Dick", "Ernie"},
			},
		}
		for name, test := range tests {
			sortFunction(StringSortable(test.slice))
			if !reflect.DeepEqual(test.slice, test.want) {
				t.Errorf("%s: got %v but want %v", name, test.slice, test.want)
			}
		}
	}
}

// TestTimeSort tests all sort functions with slices of time.Time.
// Test data is provided in a map.
func TestTimeSort(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		date1 := time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC)
		date2 := time.Date(1974, time.November, 29, 4, 3, 3, 0, time.UTC)
		date3 := time.Date(1982, time.July, 21, 4, 3, 3, 0, time.UTC)
		date4 := time.Date(2001, time.September, 11, 14, 6, 3, 0, time.UTC)
		date5 := time.Date(2023, time.March, 2, 10, 10, 0, 0, time.UTC)
		tests := map[string]struct {
			slice []time.Time
			want  []time.Time
		}{
			"empty_input": {
				slice: []time.Time{},
				want:  []time.Time{},
			},
			"one_element": {
				slice: []time.Time{date1},
				want:  []time.Time{date1},
			},
			"three_elements_sorted": {
				slice: []time.Time{date1, date2, date3},
				want:  []time.Time{date1, date2, date3},
			},
			"three_elements_not_sorted": {
				slice: []time.Time{date3, date1, date2},
				want:  []time.Time{date1, date2, date3},
			},
			"five_elements_sorted": {
				slice: []time.Time{date1, date2, date3, date4, date5},
				want:  []time.Time{date1, date2, date3, date4, date5},
			},
			"five_elements_not_sorted": {
				slice: []time.Time{date3, date5, date1, date4, date2},
				want:  []time.Time{date1, date2, date3, date4, date5},
			},
		}
		for name, test := range tests {
			sortFunction(TimeSortable(test.slice))
			if !reflect.DeepEqual(test.slice, test.want) {
				t.Errorf("%s: got %v but want %v", name, test.slice, test.want)
			}
		}
	}
}

// TestLargeIntSlice tests all sort functions with a large unsorted slice of ints.
func TestLargeIntSlice(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := CreateRandomInts(1000)
		want := make([]int, 1000)
		copy(want, slice)
		sort.Ints(want)
		sortFunction(IntSortable(slice))
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v", "large_slice", slice, want)
		}
	}
}

// TestLargeStringSlice tests all sort functions with a large unsorted slice of strings.
func TestLargeStringSlice(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := CreateRandomStrings(1000, 5)
		want := make([]string, 1000)
		copy(want, slice)
		sort.Strings(want)
		sortFunction(StringSortable(slice))
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v", "large_slice", slice, want)
		}
	}
}

// TestLargeTimeSlice tests all sort functions with a large unsorted slice of date/times.
func TestLargeTimeSlice(t *testing.T) {
	for _, sortFunction := range SortFunctions {
		slice := CreateRandomTimes(1000)
		want := make([]time.Time, 1000)
		copy(want, slice)
		sort.Slice(want, func(i, j int) bool { return want[i].UnixNano() < want[j].UnixNano() })
		sortFunction(TimeSortable(slice))
		if !reflect.DeepEqual(slice, want) {
			t.Errorf("%s: got %v but want %v", "large_slice", slice, want)
		}
	}
}

// TestSelectBestPivot tests the selectBestPivot function.
func TestSelectBestPivot(t *testing.T) {
	tests := map[string]struct {
		slice []int
		from  int
		to    int
		want  []int
	}{
		"three_elements_sorted": {
			slice: []int{1, 2, 3},
			from:  0,
			to:    2,
			want:  []int{2, 1, 3},
		},
		"three_elements_not_sorted": {
			slice: []int{3, 1, 2},
			from:  0,
			to:    2,
			want:  []int{2, 1, 3},
		},
		"all_elements_positive": {
			slice: []int{1, 2, 3, 4, 5, 6, 7},
			from:  0,
			to:    6,
			want:  []int{4, 2, 3, 1, 5, 6, 7},
		},
		"all_elements_positive_negative": {
			slice: []int{-3, -2, -1, 0, 1, 2, 3},
			from:  0,
			to:    6,
			want:  []int{0, -2, -1, -3, 1, 2, 3},
		},
		"even_number_elements": {
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8},
			from:  0,
			to:    7,
			want:  []int{4, 2, 3, 1, 5, 6, 7, 8},
		},
		"mixed_and_even_number_elements": {
			slice: []int{2, 6, 3, 8, 7, 1, 4, 9},
			from:  0,
			to:    7,
			want:  []int{8, 6, 3, 2, 7, 1, 4, 9},
		},
		"mixed_and_range_prefix": {
			slice: []int{7, 2, 2, 6, 3, 8, 7, 1, 4, 9, 5},
			from:  2,
			to:    9,
			want:  []int{7, 2, 8, 6, 3, 2, 7, 1, 4, 9, 5},
		},
		"mixed_and_range_suffix": {
			slice: []int{2, 6, 3, 8, 7, 1, 4, 9, 5, 1, 2, 3, 4, 5},
			from:  0,
			to:    7,
			want:  []int{8, 6, 3, 2, 7, 1, 4, 9, 5, 1, 2, 3, 4, 5},
		},
	}
	for name, test := range tests {
		selectBestPivot(IntSortable(test.slice), test.from, test.to)
		if !reflect.DeepEqual(test.slice, test.want) {
			t.Errorf("%s: got %v but want %v", name, test.slice, test.want)
		}
	}
}

// TestSplitUsingPivot tests the splitUsingPivot function.
func TestSplitUsingPivot(t *testing.T) {
	tests := map[string]struct {
		slice []int
		from  int
		to    int
		want  []int
	}{
		"three_elements_not_sorted": {
			slice: []int{2, 1, 3},
			from:  0,
			to:    2,
			want:  []int{1, 2, 3},
		},
		"four_elements_sorted": {
			slice: []int{1, 2, 3, 4},
			from:  0,
			to:    3,
			want:  []int{1, 3, 4, 2},
		},
		"all_elements_positive": {
			slice: []int{4, 2, 3, 1, 5, 6, 7},
			from:  0,
			to:    6,
			want:  []int{1, 2, 3, 4, 6, 7, 5},
		},
		"all_elements_positive_negative": {
			slice: []int{0, -2, -1, -3, 1, 2, 3},
			from:  0,
			to:    6,
			want:  []int{-3, -2, -1, 0, 2, 3, 1},
		},
		"even_number_elements": {
			slice: []int{4, 2, 3, 1, 5, 6, 7, 8},
			from:  0,
			to:    7,
			want:  []int{1, 2, 3, 4, 6, 7, 8, 5},
		},
		"mixed_and_even_number_elements": {
			slice: []int{8, 6, 3, 2, 7, 1, 4, 9},
			from:  0,
			to:    7,
			want:  []int{4, 6, 3, 2, 7, 1, 8, 9},
		},
		"mixed_and_range": {
			slice: []int{3, 3, 8, 6, 3, 2, 7, 1, 4, 9},
			from:  2,
			to:    9,
			want:  []int{3, 3, 4, 6, 3, 2, 7, 1, 8, 9},
		},
	}
	for name, test := range tests {
		splitUsingPivot(IntSortable(test.slice), test.from, test.to)
		if !reflect.DeepEqual(test.slice, test.want) {
			t.Errorf("%s: got %v but want %v", name, test.slice, test.want)
		}
	}
}
