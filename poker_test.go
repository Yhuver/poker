package main

import "testing"

func TestIsStraight(t *testing.T) {
	tests := []struct {
		cards    []int
		expected bool
	}{
		{[]int{2, 3, 4, 5, 6}, true},
		{[]int{14, 5, 4, 2, 3}, true},
		{[]int{7, 7, 12, 11, 3, 4, 14}, false},
		{[]int{7, 3, 2}, false},
	}

	for _, test := range tests {

		result := isStraight(test.cards)

		if result != test.expected {
			t.Errorf("isStraight(%v) = %t, esperado %t", test.cards, result, test.expected)
		}
	}
}
