package samples

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) insertAtBeginning(data int) {
	newNode := &Node{data: data}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	newNode.next = ll.head
	ll.head = newNode
}

func (ll *LinkedList) append(data int) {
	newNode := &Node{data: data}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func (ll *LinkedList) delete(data int) {
	if ll.head == nil {
		return
	}

	if ll.head.data == data {
		ll.head = ll.head.next
		return
	}

	current := ll.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

func (ll *LinkedList) display() {
	if ll.head == nil {
		fmt.Println("Linked list is empty.")
		return
	}

	current := ll.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func manipulating() {
	ll := LinkedList{}

	ll.insertAtBeginning(3) // Insert at the beginning
	ll.insertAtBeginning(2)
	ll.insertAtBeginning(1)

	ll.append(4) // Append at the end
	ll.append(5)

	ll.display() // Output: 1 2 3 4 5

	ll.delete(3) // Delete node with value 3
	ll.display() // Output: 1 2 4 5
}
