package problem146

import "container/list"

type Item struct {
	Key   int
	Value int
}

type LRUCache struct {
	Capacity  int
	itemsMap  map[int]*list.Element
	itemsList *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity:  capacity,
		itemsMap:  make(map[int]*list.Element),
		itemsList: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	el, ok := this.itemsMap[key]
	if !ok {
		return -1
	}
	this.itemsList.MoveToFront(el)
	return el.Value.(*Item).Value
}

func (this *LRUCache) Put(key int, value int) {
	el, ok := this.itemsMap[key]
	if ok {
		el.Value.(*Item).Value = value
		this.itemsList.MoveToFront(el)
		return
	}

	el = this.itemsList.PushFront(&Item{key, value})
	this.itemsMap[key] = el

	if this.itemsList.Len() > this.Capacity {
		el = this.itemsList.Back()
		this.itemsList.Remove(el)
		delete(this.itemsMap, el.Value.(*Item).Key)
	}
}

func (this *LRUCache) Len() int {
	return this.itemsList.Len()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
