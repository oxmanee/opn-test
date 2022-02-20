package main

import (
	"testing"
)

var dataSet = []int{90, 48, 63, 75, 70}

var expected = []string{"A", "F", "D", "C", "C"}

func TestFindFirstNumeric(t *testing.T) {
	for i, v := range dataSet {
		t.Run("success", func(t *testing.T) {
			result := calculateGrade(v)
			if result != expected[i] {
				t.Errorf("got %d, result %s, want %s", v, result, expected[i])
			}
		})
	}
}
