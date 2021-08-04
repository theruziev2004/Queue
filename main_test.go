package main

import "testing"

func TestList_Add(t *testing.T) {
	q := List{}
	if q.length != 0 {
		want := 0
		got := q.len()
		t.Errorf("method len() in Method( add ) is not correct for empty queue got %d want %d\n", got, want)
	}

	person := ListNode{
		Next:      nil,
		Prev:      nil,
		Name:      "Umed",
		Purchases: 150,
	}
	q.Add(person)

	if q.length != 1 {
		want := 0
		got := q.len()
		t.Errorf("method len() in Method( add ) is not correct for empty queue got %d want %d\n", got, want)
	}
}