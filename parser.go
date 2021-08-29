package parser

import "strings"

const (
	insideMap = iota
	insideArr
)

type stack []int

// IsEmpty: check if stack is empty
func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *stack) Push(val int) {
	*s = append(*s, val) // Simply append the new value to the end of the stack
}

// Peek the top value on the stack
func (s *stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func ToJSON(input string) ([]byte, error) {
	var stack stack
	var out strings.Builder
	keyExpected, valExpected := false, false
	i := 0
	for i < len(input) {
		char := input[i]

		if keyExpected {
			if char == '}' {
				continue
			}
			out.WriteString(`"`)
			for i < len(input) && char != ':' {
				out.WriteByte(char)
				i++
				char = input[i]
			}
			out.WriteString(`":`)
			i++
			char = input[i]
			keyExpected = false
			valExpected = true
		}

		if valExpected {
			if input[i:i+4] == "map[" {
				i += 4
				out.WriteString("{")
				stack.Push(insideMap)
				keyExpected = true
				continue
			}
			if input[i] == '[' {
				i++
				out.WriteString("[")
				stack.Push(insideArr)
				continue
			}
			out.WriteString(`"`)
			for i < len(input)-1 && char != ' ' && char != ']' {
				out.WriteByte(char)
				i++
				char = input[i]
			}
			out.WriteString(`"`)
			valExpected = false
		}

		switch char {
		case '{', '[':
			out.WriteString(`{`)
			i++
			if input[i] != '}' && input[i] != ']' {
				if currLoc, got := stack.Peek(); got && currLoc == insideArr {
					valExpected = true
					keyExpected = false
				} else {
					keyExpected = true
				}
			}
			continue
		case ' ':
			i++
			out.WriteString(`, `)
			if input[i] != '}' && input[i] != ']' {
				if currLoc, got := stack.Peek(); got && currLoc == insideArr {
					valExpected = true
					keyExpected = false
				} else {
					keyExpected = true
				}
			}
			continue
		case '}', ']':
			if currLoc, got := stack.Pop(); got && currLoc == insideArr {
				out.WriteString(`]`)
			} else {
				out.WriteString(`}`)
			}
			i++
			continue
		}
		i++
	}

	return []byte(out.String()), nil
}
