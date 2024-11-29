package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMonotonic_Push(t *testing.T) {
	cases := [...]struct {
		push       []int
		expected   []int
		increasing bool
	}{
		// Case: Pop if the value that be pushed is minor than the last value
		{
			push:       []int{1, 7, 9, 5, 8},
			expected:   []int{8, 5, 1},
			increasing: true,
		},
		// Case: Pop if the value that be pushed is greater than the last value
		{
			push:       []int{1, 7, 9, 5, 8},
			expected:   []int{8, 9},
			increasing: false,
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var stack Monotonic[int]
			stack.Increasing = c.increasing

			order := "Decreasing"
			if c.increasing {
				order = "Increasing"
			}

			t.Logf("%s order from bottom to top\n", order)

			// Pushing items
			for _, n := range c.push {
				stack.Push(n)
			}

			// Recovering items
			results := make([]int, 0, len(c.push))

			for range stack.Length {
				results = append(results, stack.Pop())
			}

			// Comparing results
			if !reflect.DeepEqual(results, c.expected) {
				t.Fatalf("expected %v, got %v", c.expected, results)
			}

			t.Logf("Success! %v\n", results)
		})
	}
}
