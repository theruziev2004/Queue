package main

import "fmt"

type List struct {
	root   *ListNode
	end    *ListNode
	length int
}

func (receiver *List) len() int {
	return receiver.length
}

func (receiver *List) PrintList() {
	fmt.Println("Start print...")
	current := receiver.root
	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
	fmt.Println("Finish")
}

type ListNode struct {
	Prev      *ListNode
	Next      *ListNode
	Name      string
	Purchases int
}

func (receiver *List) Add(node ListNode) {
	if receiver.len() == 0 {
		receiver.root = &node
		receiver.end = &node
		receiver.length++
		return
	}
	lastElement := receiver.end
	receiver.end.Next = &node
	receiver.end = &node
	receiver.end.Prev = lastElement
	receiver.length++
}

func (receiver *List) PopLast() {
	if receiver.len() == 0 {
		return
	}

	previousToDelete := receiver.root
	for previousToDelete.Next.Next != nil {
		previousToDelete = previousToDelete.Next
	}
	previousToDelete.Next = nil
	previousToDelete = receiver.end
	receiver.length--
}

func (receiver *List) PopFirst() {
	if receiver.len() == 0 {
		return
	}
	nextToDelete := receiver.root.Next
	receiver.root = nextToDelete
	nextToDelete.Prev = nil
	receiver.length--
}

func main() {
	person := ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Rustam",
		Purchases: 100,
	}
	queue := List{}
	queue.Add(person)
	fmt.Println(queue)

	person = ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Surush",
		Purchases: 200,
	}

	queue.Add(person)
	fmt.Println(queue)

	person = ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Aziz",
		Purchases: 300,
	}
	queue.Add(person)
	fmt.Println(queue)

	person = ListNode{
		Prev:      nil,
		Next:      nil,
		Name:      "Amin",
		Purchases: 400,
	}
	queue.Add(person)

	queue.PrintList()

	queue.PopLast()
	queue.PrintList()

}