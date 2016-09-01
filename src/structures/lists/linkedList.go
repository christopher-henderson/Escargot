package lists

type Node struct {
	Key   string
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(key string, data interface{}) *Node {
	return &Node{key, data, nil, nil}
}

type LinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func (l *LinkedList) Push(key string, data interface{}) (*Node, error) {
	defer l.increment()
	node := NewNode(key, data)
	if l.size == 0 {
		l.Head = node
		l.Tail = node
		return node, nil
	}
	l.Head.Left = node
	node.Right = l.Head
	l.Head = node
	return node, nil
}

func (l *LinkedList) Append(key string, data interface{}) (*Node, error) {
	defer l.increment()
	node := NewNode(key, data)
	if l.size == 0 {
		l.Head = node
		l.Tail = node
		return node, nil
	}
	node.Left = l.Tail
	l.Tail.Right = node
	l.Tail = node
	return node, nil
}

func (l *LinkedList) MoveToFront(node *Node) {
	if node == l.Head {
		return
	}
	if node == l.Tail {
		l.Tail = node.Left
	}
	if node.Left != nil {
		node.Left.Right = node.Right
	}
	if node.Right != nil {
		node.Right.Left = node.Left
	}
	node.Right = l.Head
	l.Head.Left = node
	l.Head = node
}

func (l *LinkedList) GetHead() interface{} {
	if l.Head != nil {
		return l.Head.Data
	}
	return nil
}

func (l *LinkedList) Pop() interface{} {
	if l.Head == nil {
		return nil
	}
	head := l.PopNode()
	if head == nil {
		return head
	}
	return head.Data
}

func (l *LinkedList) PopNode() *Node {
	if l.Head == nil {
		return nil
	}
	defer l.decrement()
	head := l.Head
	l.Head = l.Head.Right
	if l.Head != nil {
		l.Head.Left = nil
	}
	return head
}

func (l *LinkedList) GetTail() interface{} {
	if l.Tail != nil {
		return l.Tail.Data
	}
	return nil
}

func (l *LinkedList) Slice() []interface{} {
	node := l.Head
	var slice []interface{}
	for node != nil {
		slice = append(slice, node.Data)
		node = node.Right
	}
	return slice
}

func (l *LinkedList) increment() {
	l.size++
}

func (l *LinkedList) decrement() {
	l.size--
}
