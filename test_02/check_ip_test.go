package main

import "testing"

var dataSet = []string{"172.16.254.1", "0.1.1.256", "1.1.1.1a", "1", "64.233.16.00", "7283728"}

var expected = []bool{true, false, false, false, false, false}

func TestFindFirstNumeric(t *testing.T) {
	for i, v := range dataSet {
		t.Run("success", func(t *testing.T) {
			result := checkIp(v)
			if result != expected[i] {
				t.Errorf("got %s, result %t, want %t", v, result, expected[i])
			}
		})
	}
}
