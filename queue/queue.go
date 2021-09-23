package queue

func Enqueue(s string, queue []string) []string {
	return append(queue, s)
}

func Dequeue(queue []string) (string, []string) {
	item := queue[0]
	queue = queue[1:]
	return item, queue
}

func Peek(queue []string) *string {
	if len(queue) == 0 {
		return nil
	}
	return &queue[0]
}

func Count(queue []string) int {
	return len(queue)
}