package createlinkedlist

//Node ... This defines the structure of linked lists
type Node struct {
	data int
	next *Node
}

//InsertNodeAtBeginning ... Inserts node at beginning of linkedlist
func InsertNodeAtBeginning(head *Node, data int) *Node {
	// Receives head node as pointer.
	newnode := Node{data, nil}
	if head == nil {
		head = &newnode
	} else {
		newnode.next = head
		head = &newnode
	}
	return head
}

//InsertNodeAtEnd ... Inserts node at end of linkedlist
func InsertNodeAtEnd(head *Node, data int) *Node {
	// Receives head node as pointer.
	newnode := Node{data, nil}
	if head == nil {
		head = &newnode
	} else {
		temp := head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = &newnode
	}
	return head
}

//LengthOfLinkedList ... Returns length of linked list
func LengthOfLinkedList(head *Node) int {
	temp := head
	count := 0
	for temp != nil {
		count++
	}
	return count
}

//SearchElement ... Search an element iteratively and returns true/false
func SearchElement(head *Node, element int) bool {
	temp := head
	for temp != nil && temp.data != element {
		temp = temp.next
	}
	if temp == nil {
		return false
	}
	return true
}
