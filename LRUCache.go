package cache

import "sync"

type Node struct {
	key  interface{}
	val  interface{}
	pre  *Node
	next *Node
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		table:    make(map[interface{}]*Node),
	}
}

type LRUCache struct {
	sync.Mutex // prevent multiple goroutine write

	capacity int // the capacity of cache
	table    map[interface{}]*Node
	head     *Node
	end      *Node
}

// set node as the head of cache
func (cache *LRUCache) setHead(node *Node) {
	cache.Lock()
	defer cache.Unlock()
	node.next = cache.head
	node.pre = nil
	if cache.head != nil {
		cache.head.pre = node
	}
	cache.head = node
	if cache.end == nil {
		cache.end = node
	}
}

func (cache *LRUCache) delete(node *Node) {
	cache.Lock()
	defer cache.Unlock()
	if node.pre != nil {
		node.pre.next = node.next
	} else {
		cache.head = node.next // if node.pre == nil means node is cache.head
	}

	if node.next != nil {
		node.next.pre = node.pre
	} else {
		cache.end = node.pre // if node.next == nil means node is cache.end
	}

}

func (cache *LRUCache) Set(key, val interface{}) error {
	if v, ok := cache.table[key]; ok {
		v.val = val
		cache.delete(v)
		cache.setHead(v)
		return nil
	}
	node := &Node{
		key: key,
		val: val,
	}
	if len(cache.table) >= cache.capacity {
		delete(cache.table, cache.end.key)
		cache.delete(cache.end)
	}
	cache.setHead(node)
	cache.Lock()
	cache.table[key] = node
	cache.Unlock()
	return nil
}

func (cache *LRUCache) Get(key interface{}) (interface{}, error) {
	if v, ok := cache.table[key]; ok {
		cache.delete(v)
		cache.setHead(v)
		return v.val, nil
	}
	return nil, nil
}
