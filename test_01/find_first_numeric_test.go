package main

import "testing"

var dataSet = []string{"Var-----___1_int", "Q2q-q", "eef3243**s"}

var expected = []string{"1", "2", "3"}

func TestFindFirstNumeric(t *testing.T) {
	for i, v := range dataSet {
		t.Run("success", func(t *testing.T) {
			result := findFirstNumeric(v)
			if result == nil {
				t.Errorf("got %s, result %s, want %s", v, *result, expected[i])
			} else {
				if *result != expected[i] {
					t.Errorf("got %s, result %s, want %s", v, *result, expected[i])
				}
			}
		})
	}
}
