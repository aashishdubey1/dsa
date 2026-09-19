package lrucache

import (
	"errors"
)

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type DLL struct {
	Head *Node
	Tail *Node
}

type LRUCache struct {
	capacity int
	items    map[int]*Node
	list     *DLL
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		list:     &DLL{},
		items:    make(map[int]*Node),
	}
}

func (cache *LRUCache) Get(key int) (int, bool) {
	node, exist := cache.items[key]
	if !exist {
		return 0, false
	}

	cache.list.RemoveNode(node)
	cache.list.InsertAtFront(node)

	return node.Value, true
}

func (cache *LRUCache) Put(key int, val int) {
	node, exist := cache.items[key]
	if exist {
		node.Value = val
		cache.list.RemoveNode(node)
		cache.list.InsertAtFront(node)
		return
	}

	if len(cache.items) >= cache.capacity {
		evicted := cache.list.Tail
		cache.list.RemoveNode(evicted)
		delete(cache.items, evicted.Key)
	}

	newNode := &Node{
		Value: val,
		Key:   key,
	}

	cache.list.InsertAtFront(newNode)
	cache.items[key] = newNode
}

func (list *DLL) InsertAtFront(node *Node) {
	node.Prev = nil

	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return
	}

	curr := list.Head
	node.Next = curr
	curr.Prev = node

	list.Head = node
}

func (list *DLL) RemoveNode(node *Node) (bool, error) {
	if list.Head == nil {
		return false, errors.New("list is empty")
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		list.Head = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		list.Tail = node.Prev
	}

	node.Next = nil
	node.Prev = nil
	return true, nil
}
