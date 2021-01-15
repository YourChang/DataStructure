package Queue

type Queue interface {
	EnQueue(v interface{}) bool
	DeQueue() interface{}
	String() string
	IsEmpty() bool
	IsFull() bool
}
