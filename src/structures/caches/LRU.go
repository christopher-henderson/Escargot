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
	l.list.MoveToFront(node)
	return node.Data
}

func (l *LRU) Put(key string, data interface{}) {
	if node, ok := l.m[key]; ok {
		node.Data = data
		l.list.MoveToFront(node)
		return
	}
	newNode, err := l.list.Push(data)
	if err != nil {
		return
	}
	l.m[key] = newNode
}

func (l *LRU) Slice() []interface{} {
	return l.list.Slice()
}
