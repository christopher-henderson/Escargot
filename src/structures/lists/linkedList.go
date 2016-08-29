package lists

type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
}

func NewNode(Data interface{}) *Node {
	return &Node{Data, nil, nil}
}

type LinkedList struct {
	Head *Node
	Tail *Node
	size int
}

func (l *LinkedList) Push(data interface{}) (*Node, error) {
	defer l.increment()
	if l.size == 0 {
		return l.startCollection(data)
	}
	node := NewNode(data)
	l.Head.Left = node
	node.Right = l.Head
	l.Head = node
	return node, nil
}

func (l *LinkedList) Append(data interface{}) (*Node, error) {
	defer l.increment()
	if l.size == 0 {
		return l.startCollection(data)
	}
	node := NewNode(data)
	node.Left = l.Tail
	l.Tail.Right = node
	l.Tail = node
	return node, nil
}

func (l *LinkedList) startCollection(data interface{}) (*Node, error) {
	node := NewNode(data)
	l.Head = node
	l.Tail = node
	return node, nil
}

func (l *LinkedList) moveToFront(node *Node) {
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
	defer l.decrement()
	Data := l.Head.Data
	l.Head = l.Head.Right

	return Data
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
