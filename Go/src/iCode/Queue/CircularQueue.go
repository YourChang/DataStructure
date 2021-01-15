package Queue

type CircularQueue struct {
	q []interface{}
	capacity int
	head int
	tail int
}


