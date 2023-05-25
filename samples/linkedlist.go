package samples

import "fmt"

// Write a program to add an element at position k in the linked list if k is greater than the size of the linked list add at last

type Nodes struct {
	data int
	next *Nodes
}

type LinkedLists struct {
	head *Nodes
}

func (ll *LinkedLists) insertAtPosition(data, position int) {
	newNode := &Nodes{data: data}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	if position <= 0 {
		newNode.next = ll.head
		ll.head = newNode
		return
	}

	current := ll.head
	previous := current

	for position > 0 && current != nil {
		previous = current
		current = current.next
		position--
	}

	previous.next = newNode
	newNode.next = current

	if current == nil {
		fmt.Println("Position is greater than the size of the linked list. Adding at the last position.")
	}
}

func (ll *LinkedLists) display() {
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

func LinkedListSampleStart() {
	ll := LinkedLists{}

	ll.insertAtPosition(1, 0)  // Add at position 0 (head position)
	ll.insertAtPosition(2, 1)  // Add at position 1
	ll.insertAtPosition(3, 10) // Position is greater than size, add at last
	ll.insertAtPosition(4, -1) // Add at position 0 (head position)
	ll.insertAtPosition(5, 2)  // Add at position 2
	ll.insertAtPosition(6, 3)  // Add at position 3

	ll.display() // Output: 4 1 5 6 2 3
}
