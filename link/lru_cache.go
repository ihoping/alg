package link

import "fmt"

// https://leetcode-cn.com/problems/lru-cache/

//单向链表
type LinkedNode struct {
	key   int
	value int
	next  *LinkedNode
}

//双向链表
type DoublyLinkedNode struct {
	key   int
	value int
	next  *DoublyLinkedNode
	prev  *DoublyLinkedNode
}

type LRUCache struct {
	cap  int //最大容量
	cCap int //当前容量
	m    map[int]*DoublyLinkedNode
	h    *DoublyLinkedNode //双向链表头节点
	e    *DoublyLinkedNode //双向链表尾指针
}

func Constructor(capacity int) LRUCache {
	end := &DoublyLinkedNode{}
	head := &DoublyLinkedNode{
		next: end,
	}

	return LRUCache{
		cap:  capacity,
		cCap: 0,
		m:    make(map[int]*DoublyLinkedNode),
		h:    head,
		e:    end,
	}
}

func (lc *LRUCache) Get(key int) int {
	if node, ok := lc.m[key]; ok {
		//将该node放到队列前面
		node.prev.next = node.next
		node.next.prev = node.prev

		tNode := lc.h.next
		lc.h.next = node
		node.prev = lc.h

		node.next = tNode
		tNode.prev = node

		return node.value
	}
	return -1
}

func (lc *LRUCache) Put(key int, value int) {
	//是否已存在
	if node, ok := lc.m[key]; ok {
		//更新该node
		node.value = value

		//将该node放到队列前面
		node.prev.next = node.next
		node.next.prev = node.prev

		tNode := lc.h.next
		lc.h.next = node
		node.prev = lc.h

		node.next = tNode
		tNode.prev = node
		return
	}

	newNode := &DoublyLinkedNode{
		key:   key,
		value: value,
	}

	//放在队列最前面
	tNode := lc.h.next

	lc.h.next = newNode
	newNode.prev = lc.h
	newNode.next = tNode
	tNode.prev = newNode

	//放到map
	lc.m[key] = newNode

	//cap++
	lc.cCap++

	//判断是否超出大小
	if lc.cCap > lc.cap {
		//删除m缓存的
		delete(lc.m, lc.e.prev.key)

		//删掉最后一个节点
		lc.e.prev.prev.next = lc.e
		lc.e.prev = lc.e.prev.prev

		lc.cCap--
	}
}

func (lc *LRUCache) Print() {
	node := lc.h.next
	for node.value != 0 {
		fmt.Println(node.value)
		node = node.next
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
