package samples

import "fmt"

type NodeMiddle struct {
	Data int
	Next *NodeMiddle
}

type LinkedListMiddle struct {
	Head *NodeMiddle
}

func (ll *LinkedListMiddle) Insert(data int) {
	newNode := &NodeMiddle{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

func (ll *LinkedListMiddle) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

func (ll *LinkedListMiddle) RemoveMiddle() {
	if ll.Head == nil || ll.Head.Next == nil {
		return
	}

	// Find the middle node
	slow := ll.Head
	fast := ll.Head
	var prev *NodeMiddle

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		prev = slow
		slow = slow.Next
	}

	// Remove the middle node
	prev.Next = slow.Next
}

func main() {
	ll := LinkedListMiddle{}

	// Insert elements
	ll.Insert(1)
	ll.Insert(2)
	ll.Insert(3)
	ll.Insert(4)
	ll.Insert(5)

	fmt.Println("Original List:")
	ll.Display()

	// Remove middle element
	ll.RemoveMiddle()

	fmt.Println("List after removing the middle element:")
	ll.Display()
}
