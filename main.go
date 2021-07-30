package main

import (
	"fmt"
)

type List struct {
	root   *ListNode //first element
	end    *ListNode // last element
	lenght int
}

func (receiver *List) PrintList() {
	fmt.Println("Start print")
	current := receiver.root //current - текущий
	for current != nil{
		fmt.Println(current)
		current = current.Next
	}
	fmt.Println("Finish")
}


func (receiver *List) len() int {
	return receiver.lenght
}

func (receiver *List) Add(node ListNode) {
	if receiver.len() == 0 {
		receiver.root = &node
		receiver.end = &node
		receiver.lenght++
		return // ранний выход
	}

	////
	LastElement := receiver.end /// храню предпоследнего человека
	receiver.end.Next = &node
	receiver.end = &node
	receiver.end.Prev = LastElement
	receiver.lenght++

}

type ListNode struct {
	Prev      *ListNode
	Next      *ListNode
	Name      string
	Purchases int
}

func main() {
	person := ListNode{
		Next:      nil,
		Prev:      nil,
		Name:      "Rustam",
		Purchases: 100,
	}
	queue := List{}
	queue.Add(person)
	fmt.Println(queue)

	person = ListNode{
		Next:      nil,
		Prev:      nil,
		Name:      "Umed",
		Purchases: 150,
	}
	queue.Add(person)
	fmt.Println(queue)

	person = ListNode{
		Next:      nil,
		Prev:      nil,
		Name:      "Surush",
		Purchases: 120,
	}
	queue.Add(person)
	fmt.Println(queue)

	queue.PrintList()

	fmt.Println(queue.root)
	fmt.Println(queue.end)
}
