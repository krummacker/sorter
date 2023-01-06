package main

import (
	"testing"
)

// TestCreateRandomInts tests the createRandomInts function.
func TestCreateRandomInts(t *testing.T) {

	tests := map[string]struct {
		size int
		want int
	}{
		"zero_size": {
			size: 0,
			want: 0,
		},
		"size_one": {
			size: 1,
			want: 1,
		},
		"size_thousand": {
			size: 1000,
			want: 1000,
		},
	}

	for name, test := range tests {
		got := CreateRandomInts(test.size)
		if len(got) != test.want {
			t.Errorf("%s: got %v but want %v", name, got, test.want)
		}
	}
}

// TestAverage tests the average function.
func TestAverage(t *testing.T) {

	tests := map[string]struct {
		input []int
		want  int
	}{
		"zero_size": {
			input: []int{},
			want:  0,
		},
		"size_one": {
			input: []int{42},
			want:  42,
		},
		"size_many": {
			input: []int{1, 2, 3, 4, 5, 6, 7},
			want:  4,
		},
	}

	for name, test := range tests {
		got := Average(test.input)
		if got != test.want {
			t.Errorf("%s: got %v but want %v", name, got, test.want)
		}
	}
}
