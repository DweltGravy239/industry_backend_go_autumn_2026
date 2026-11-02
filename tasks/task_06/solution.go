package main

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}
type LRUCache[K comparable, V any] struct {
	capacity int
	ll       list.List
	items    map[K]*list.Element
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{capacity: capacity, items: make(map[K]*list.Element)}
}
func (c *LRUCache[K, V]) Get(key K) (value V, ok bool) {
	e, found := c.items[key]
	if !found { return }
	result := e.Value.(entry[K,V])
	c.ll.MoveToFront(e)
	result = e.Value.(entry[K,V])
	return result.value, true
}
func (c *LRUCache[K, V]) Set(key K, value V) {
	if c.capacity <= 0 { return }
	e, found := c.items[key]
	if found { 
		e.Value = entry[K,V]{key: key, value: value} 
		return 
	}
	if len(c.items) >= c.capacity {
		last := c.ll.Back()
		c.ll.Remove(last)
		delete(c.items, last.Value.(entry[K,V]).key)
	}
	e = c.ll.PushFront(entry[K, V]{key: key, value: value})
	c.items[key] = e
}

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
}
