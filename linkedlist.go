package main

import "fmt"

type Node struct {
    data int
    next *Node
}

func reverse(head *Node, k int) *Node {
    if head == nil {
        return nil
    }

    var prev, curr, Next *Node
    curr = head

    for k > 0 && curr != nil {
        Next = curr.next
        curr.next = prev
        prev = curr
        curr = Next
        k = k - 1
    }

    // Recursive call for the rest of the linked list
    if Next != nil {
        head.next = reverse(Next, 2)
    }

    return prev
}

func printLinkedList(head *Node) {
    curr := head
    for curr != nil {
        fmt.Print(curr.data, "--->")
        curr = curr.next
    }
    fmt.Println()
}

func main() {
    head := &Node{data: 1, next: &Node{data: 2, next: &Node{data: 3, next: nil}}}
    fmt.Println("Original Linked List:")
    printLinkedList(head)

    k := 2
    head = reverse(head, k)

    fmt.Println("Reversed Linked List:")
    printLinkedList(head)
}

