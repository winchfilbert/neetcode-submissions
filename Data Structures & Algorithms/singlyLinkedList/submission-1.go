type Node struct {
	value int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil, tail: nil}
}

func (ll *LinkedList) Get(index int) int {
	iter := 0
	curr := ll.head
	for curr != nil{
		if(iter == index){
			return curr.value
		}
		iter++
		curr = curr.next
	}
	return -1
}

func (ll *LinkedList) InsertHead(val int) {
	node := &Node{value: val}
	if(ll.head == nil){
		ll.head = node
		ll.tail = node
	}else {
		node.next = ll.head
		ll.head = node
	}
}

func (ll *LinkedList) InsertTail(val int) {
	node := &Node{value: val}
	if(ll.tail == nil){
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}
}

func (ll *LinkedList) Remove(index int) bool {
    if ll.head == nil || index < 0 {
        return false
    }

    // if head
    if index == 0 {
        ll.head = ll.head.next
        if ll.head == nil { // If the list is now empty, fix the tail too
            ll.tail = nil
        }
        return true
    }

    // Remove elements after the head
    iter := 0
    curr := ll.head
    
    // Move 'curr' until it points to the node right BEFORE the one we want to delete
    for curr.next != nil && iter < index-1 {
        curr = curr.next
        iter++
    }

	// out of bound error
    if curr.next == nil {
        return false
    }

    // if tail
    if curr.next == ll.tail {
        ll.tail = curr
    }

    curr.next = curr.next.next
    return true
}

func (ll *LinkedList) GetValues() []int {
	var values []int
	curr := ll.head
	for curr != nil{
		values = append(values, curr.value)
		curr = curr.next
	}
	return values
}
