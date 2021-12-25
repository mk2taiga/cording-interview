package lesson0

import "fmt"

type Node struct {
	key   int
	value int
	next  *Node
}

type HashTable struct {
	Capacity int
	Table    []*Node
}

func NewHashTable(cap int) *HashTable {
	return &HashTable{cap, make([]*Node, cap)}
}

func (h *HashTable) HashValue(key int) int {
	return key % h.Capacity
}

func (h *HashTable) Search(key int) int {
	hash := h.HashValue(key)
	node := h.Table[hash]

	for node != nil {
		if node.key == key {
			return node.value
		}
		// 次のノードを検索
		node = node.next
	}

	// 検索失敗
	return -1
}

func (h *HashTable) Add(key, value int) bool {
	hash := h.HashValue(key)
	node := h.Table[hash]

	for node != nil {
		if node.key == key {
			return false
		}
		// 次のノードを検索
		node = node.next
	}

	h.Table[hash] = &Node{key: key, value: value, next: h.Table[hash]}
	return true
}

func (h *HashTable) Remove(key int) bool {
	hash := h.HashValue(key)
	node := h.Table[hash]

	var beforeNode *Node
	for node != nil {
		if node.key == key {
			if beforeNode == nil {
				h.Table[hash] = node.next
			} else {
				beforeNode.next = node.next
			}

			return true
		}
		// 次のノードを検索
		beforeNode = node
		node = node.next
	}

	return false
}

func (h *HashTable) Dump() []string {
	var msg []string
	for _, n := range h.Table {
		for n != nil {
			msg = append(msg, fmt.Sprintf("key=%v,val=%v", n.key, n.value))
			n = n.next
		}
	}
	return msg
}
