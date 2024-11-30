package stack

import (
	"reflect"
	"strconv"
	"testing"
)

func TestStack_Push(t *testing.T) {
	cases := [...]struct {
		push     []int
		expected []int
	}{
		{
			push:     []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var stack Stack[int]

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

func TestStack_Pop(t *testing.T) {
	cases := [...]struct {
		push     []int
		expected []int
	}{
		{
			push:     []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var stack Stack[int]
			results := make([]int, 0, len(c.push))

			// Pushing items
			for _, n := range c.push {
				stack.Push(n)
			}

			// Popping items
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
