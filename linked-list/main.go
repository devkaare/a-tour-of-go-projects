package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

type LinkedList struct {
    head *List[int]
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func (ll *LinkedList) Append(val int) {
    newList := &List[int]{val: val, next: nil}

    if ll.head == nil {
        ll.head = newList
        return
    }

    current := ll.head 
    for current.next != nil {
        current = current.next
    }

    current.next = newList
}

func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.val)
        current = current.next
    }
    fmt.Println("nil")
}

func main() {
    ll := NewLinkedList()

    ll.Append(10)
    ll.Append(20)
    ll.Append(30)
    ll.Append(40)

    fmt.Println("Linked list:")
    ll.Display()
}
