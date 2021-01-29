package LRU

/**
@Description: LRU(最近最久未使用替换算法),尾部是新节点
@Date: 1/29/2021
@Author: lichang
*/

const (
	hostbit = uint64(^uint(0)) == ^uint64(0)
	LEGTH = 100
)

type lruNode struct {
	prev *lruNode
	next *lruNode

	key int // lru key
	value int // lur value

	hnext *lruNode // 拉链hashnext  (意义未明...)
}

type LRUCache struct {
	node []lruNode

	head *lruNode // lru head node
	tail *lruNode // lru tail node

	capacity int
	used int
}

func Constructor(capacity int) LRUCache{
	return LRUCache{
		node:make([]lruNode, LEGTH),
		head: nil,
		tail: nil,
		capacity: capacity,
		used: 0,
	}
}

func (c *LRUCache) Get(key int) int {
	if c.tail == nil {  // empty
		return -1
	}

	if tmp := c.searchNode(key); tmp != nil {
		c.moveToTail(tmp)   // 使用后移动到尾部
		return tmp.value
	}

	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if tmp := c.searchNode(key); tmp != nil {  // key have exist
		tmp.value = value
		c.moveToTail(tmp)		// 使用后移动到尾部
		return
	}
	c.addNode(key, value)  // key not exist

	if c.used > c.capacity{  // 长度超出
		c.delNode()
	}
}

func (c *LRUCache) addNode(key int, value int) {
	newNode := &lruNode{
		prev:  nil,
		next:  nil,
		key:   key,
		value: value,
		hnext: nil,
	}

	tmp := &c.node[hash(key)]
	newNode.hnext = tmp.hnext
	tmp.hnext = newNode
	c.used++

	if c.tail == nil {
		c.tail, c.head = newNode, newNode
	}
	c.tail.next = newNode
	newNode.prev = c.tail
	c.tail = newNode
}

func (c *LRUCache) delNode() {
	if c.head == nil {
		return
	}
	prev:= &c.node[hash(c.head.key)]
	tmp := prev.hnext

	for tmp != nil && tmp.key != c.head.key{
		prev = tmp
		tmp = tmp.hnext
	}
	if tmp == nil {
		return
	}
	prev.hnext = tmp.hnext
	c.head = c.head.next
	c.head.prev = nil
	c.used--
}

func (c *LRUCache) searchNode(key int) *lruNode{
	if c.tail == nil {
		return nil
	}

	tmp := c.node[hash(key)].hnext
	for tmp != nil {
		if tmp.key == key{
			return tmp
		}
		tmp = tmp.hnext
	}
	return nil
}

func (c *LRUCache) moveToTail(node *lruNode){
	if c.tail == node {  // 当前节点已在尾部
		return
	}
	if c.head == node {		// 当前节点在头部
		c.head = node.next
		c.head.prev = nil
	} else {		// 不在头部，断链
		node.next.prev = node.prev
		node.prev.next = node.next
	}

	// 重新链接并设置尾指针
	node.next = nil
	c.tail.next = node
	node.prev = c.tail
	c.tail = node

}

// don't understand
func hash(key int) int {
	if hostbit {
		return (key ^ (key >> 32)) & (LEGTH - 1)
	}
	return (key ^ (key >> 16)) & (LEGTH - 1)
}