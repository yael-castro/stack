package stack

// Item represents an item of the stack
type Item[T any] struct {
	Value T
	Next  *Item[T]
}

// Stack is a linear data structure that follows the Last In, First Out (LIFO) principle,
// where items are added (Push) to the top and removed (Pop) from the top.
//
//	|-------|
//	|   3   |  <-- TOP
//	|-------|
//	|   2   |
//	|-------|
//	|   1   |
//	|-------|
//
// WARNING: Is not concurrent safe
type Stack[T any] struct {
	*Item[T]
	Length int
}

// Push pushes an item to top
//
// Complexity: O(1)
func (t *Stack[T]) Push(value T) {
	t.Item = &Item[T]{Value: value, Next: t.Item}
	t.Length++
}

// Pop removes an item from the top and returns its value
//
// Complexity: O(1)
func (t *Stack[T]) Pop() (value T) {
	if t.Item == nil {
		return
	}

	value = t.Item.Value
	t.Item = t.Item.Next
	t.Length--
	return
}

// Peek returns the top value
//
// Complexity: O(1)
func (t *Stack[T]) Peek() (value T) {
	if t.Item == nil {
		return
	}

	return t.Item.Value
}
