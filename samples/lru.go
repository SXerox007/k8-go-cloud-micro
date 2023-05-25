package samples

import (
	"container/list"
)

//Create struct called Node to store the value of the key and pointer to the list element.

type Node1 struct {
	Data   int
	KeyPtr *list.Element
}

//Create another struct called LRUCache
type LRUCache struct {
	Queue    *list.List
	Items    map[int]*Node1
	Capacity int
}

// init or intitilizer
func Constructor(capacity int) LRUCache {
	return LRUCache{Queue: list.New(), Items: make(map[int]*Node1), Capacity: capacity}
}

//If key is already present in the map, update the data and move the corresponding element to the front of the list.
// If not, if map size is equals to capacity of the cache,
// remove the last element of the queue and remove the key from map.
func (l *LRUCache) Put(key int, value int) {
	if item, ok := l.Items[key]; !ok {
		if l.Capacity == len(l.Items) {
			back := l.Queue.Back()
			l.Queue.Remove(back)
			delete(l.Items, back.Value.(int))
		}
		l.Items[key] = &Node1{Data: value, KeyPtr: l.Queue.PushFront(key)}
	} else {
		item.Data = value
		l.Items[key] = item
		l.Queue.MoveToFront(item.KeyPtr)
	}
}

//This will simply check if the map if key exists, if it does, then move the key to the front of the queue.
//If not, return -1
func (l *LRUCache) Get(key int) int {
	if item, ok := l.Items[key]; ok {
		l.Queue.MoveToFront(item.KeyPtr)
		return item.Data
	}
	return -1
}
