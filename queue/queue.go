package queue

func Enqueue(s string, queue string) string {
	res := queue + s
	return res
}

func Dequeue(queue string) (uint8, string) {
	item := queue[0]
	queue = queue[1:]
	return item, queue
}

func Peek(queue string) uint8 {
	if len(queue) == 0 {
		return 0
	}
	return queue[0]
}

func Count(queue string) int {
	return len(queue)
}