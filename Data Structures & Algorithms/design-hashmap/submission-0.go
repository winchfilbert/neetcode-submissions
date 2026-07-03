type Node struct{
	key int
	value int
	next *Node
}

type MyHashMap struct {
	kv []*Node // contains multiple node in a single array
}

func Constructor() MyHashMap {
	kv := make([]*Node, 1000)
	for i := range kv{
		kv[i] = &Node{key: -1, value: -1}
	}
	return MyHashMap{kv:kv}
}

func (h *MyHashMap) hash(key int) int{
	return key % len(h.kv)
}

func (this *MyHashMap) Put(key int, value int) {
		curr := this.kv[this.hash(key)]
		for curr.next != nil{
			if curr.next.key == key {
				curr.next.value = value
				return
			}
			curr = curr.next
		}
		// cara update
		curr.next = &Node{key: key, value: value}
}

func (this *MyHashMap) Get(key int) int {
	curr := this.kv[this.hash(key)]
	 for curr.next != nil {
        if curr.next.key == key {
            return curr.next.value
        }
        curr = curr.next
    }
    return -1
}

func (this *MyHashMap) Remove(key int) {
    previndex := this.kv[this.hash(key)]
	for previndex.next != nil{
			if previndex.next.key == key {
				previndex.next = previndex.next.next
				return
			}
			previndex = previndex.next
		}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */