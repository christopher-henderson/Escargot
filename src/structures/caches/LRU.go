package caches

import "Escargot/structures/lists"

type LRU struct {
	list lists.LinkedList
	m    map[string]*lists.Node
}

func NewLRU() LRU {
	return LRU{lists.LinkedList{}, make(map[string]*lists.Node)}
}

func (l *LRU) Get(key string) interface{} {
	node, ok := l.m[key]
	if !ok {
		return nil
	}
	l.moveToFront(node)
	return node.Data
}

func (l *LRU) Put(key string, data interface{}) {
	if node, ok := l.m[key]; ok {
		node.Data = data
		l.moveToFront(node)
		return
	}
	newNode := lists.NewNode(data)
	l.m[key] = newNode
	l.list.Push(newNode)
}

func (l *LRU) moveToFront(node *lists.Node) {
	if node == l.list.Head {
		return
	}
	if node == l.list.Tail {
		l.list.Tail = node.Left
	}
	if node.Left != nil {
		node.Left.Right = node.Right
	}
	if node.Right != nil {
		node.Right.Left = node.Left
	}
	node.Right = l.list.Head
	l.list.Head.Left = node
	l.list.Head = node
}

func (l *LRU) Slice() []interface{} {
	return l.list.Slice()
}
