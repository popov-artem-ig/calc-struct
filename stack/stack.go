package stack

func Push(s string, stack []string) []string {
	return append(stack, s)
}

func Pop(stack []string) (string, []string) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func Peek(stack []string) string {
	if len(stack) == 0 {
		return ""
	}
	return stack[len(stack)-1]
}

func Count(stack []string) int {
	return len(stack)
}