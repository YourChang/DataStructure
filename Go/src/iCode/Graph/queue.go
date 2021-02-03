package Graph

type intQueue struct {
	queue []int
}

func (q *intQueue) push(u int) {
	q.queue = append(q.queue, u)
}

func (q *intQueue) pop() int {
	r := q.queue[0]
	q.queue = q.queue[1:]
	return r
}

func (q *intQueue) isEmpty() bool {
	return len(q.queue) == 0
}

func (q *intQueue) getTop() int {
	return q.queue[0]
}
