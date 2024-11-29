package main

// Number defines the allowed values for the Monotonic Stack
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Monotonic Stack
//
// WARNING: It is not concurrent safe
type Monotonic[T Number] struct {
	Stack[T]
	// Increasing indicates if the Monotonic Stack is "increasing" or "decreasing"
	Increasing bool
}

// Push pushes a value with a variable behavior depending on Increasing
//
// Monotonic decreasing Stack:
// Pushes elements onto the stack while ensuring that the stack maintains a decreasing order from bottom to top.
//
// Monotonic increasing Stack:
// Pushes elements onto the stack while ensuring that the stack maintains an increasing order from bottom to top.
//
// Complexity: O(n)
func (m *Monotonic[T]) Push(value T) {
	for m.Item != nil {
		if m.Value == value || (m.Increasing && value > m.Value) || (!m.Increasing && value < m.Value) {
			break
		}

		_ = m.Pop()
	}

	m.Stack.Push(value)
}
