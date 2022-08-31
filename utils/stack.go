package utils

type Stack []float64

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Length() int {
	return len(*s)
}

func (s *Stack) ValueAt(i int) float64 {
	return (*s)[i]
}

// Push a new value onto the stack
func (s *Stack) Push(f float64) {
	*s = append(*s, f)
}

func (s *Stack) Pop() (float64, bool) {
	if s.IsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
