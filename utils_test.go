package main

import "testing"

func TestRandomArray(t *testing.T) {
	rangeEnds := 3.0
	rows := 3
	cols := 4
	size := rows * cols

	results := RandomArray(size, rangeEnds)

	if len(results) != size {
		t.Errorf("The array length was incorrect, got: %d, want: %d.", len(results), size)
	}

	for _, r := range results {
		if (r > 1.0) || (r < -1.0) {
			t.Errorf("The array contains a value out of range (-1, 1)")
		}
	}
}
