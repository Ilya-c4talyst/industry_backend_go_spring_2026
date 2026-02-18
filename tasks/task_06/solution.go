package main

type node[K comparable, V any] struct {
	key   K
	value V
	prev  *node[K, V]
	next  *node[K, V]
}

type LRUCache[K comparable, V any] struct {
	capacity int
	cache    map[K]*node[K, V]
	head     *node[K, V]
	tail     *node[K, V]
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*node[K, V]),
	}
}

func (l *LRUCache[K, V]) moveToTail(n *node[K, V]) {
	if n == l.tail {
		return
	}

	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	if n == l.head {
		l.head = n.next
	}
	n.prev = l.tail
	n.next = nil
	if l.tail != nil {
		l.tail.next = n
	}
	l.tail = n

	if l.head == nil {
		l.head = n
	}
}

func (l *LRUCache[K, V]) removeLRU() {
	if l.head == nil {
		return
	}
	delete(l.cache, l.head.key)
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
}

func (l *LRUCache[K, V]) Get(key K) (V, bool) {
	n, ok := l.cache[key]
	if !ok {
		var zeroValue V
		return zeroValue, false
	}

	l.moveToTail(n)
	return n.value, true
}

func (l *LRUCache[K, V]) Set(key K, value V) {
	if l.capacity <= 0 {
		return
	}
	if n, ok := l.cache[key]; ok {
		n.value = value
		l.moveToTail(n)
		return
	}
	n := &node[K, V]{
		key:   key,
		value: value,
	}
	if len(l.cache) >= l.capacity {
		l.removeLRU()
	}

	l.cache[key] = n
	l.moveToTail(n)
}
